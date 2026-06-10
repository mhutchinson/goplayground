package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log/slog"
	"os"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

type WebSocketServer struct {
	formatting templates
	broadcast  chan *Message

	clientsMux sync.RWMutex
	clients    map[string]clientConn
}

func NewWebSocket() *WebSocketServer {
	sentTmpl, err := template.ParseFiles("views/message_sent.html")
	if err != nil {
		slog.Error("template parsing failed", "error", err)
		os.Exit(1)
	}
	recvTmpl, err := template.ParseFiles("views/message_recv.html")
	if err != nil {
		slog.Error("template parsing failed", "error", err)
		os.Exit(1)
	}

	return &WebSocketServer{
		formatting: templates{
			recvMessage: recvTmpl,
			sentMessage: sentTmpl,
		},
		broadcast: make(chan *Message),
		clients:   make(map[string]clientConn),
	}
}

// HandleWebSocket accepts a websocker connection and services it
// for the lifetime of the connection. This method will not return
// until the connection is closed.
func (s *WebSocketServer) HandleWebSocket(conn *websocket.Conn) {
	id := uuid.New().String()
	c := clientConn{
		id:   id,
		conn: conn,
	}
	slog.Info("Registering client", "id", id)
	// Register a new Client
	s.clientsMux.Lock()
	s.clients[id] = c
	s.clientsMux.Unlock()
	defer func() {
		slog.Info("Removing client", "id", id)
		s.clientsMux.Lock()
		defer s.clientsMux.Unlock()
		delete(s.clients, id)
		conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			slog.Error("Read Error", "error", err)
			break
		}

		// send the message to the broadcast channel
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			slog.Error("Error Unmarshalling", "error", err)
			break
		}
		message.ClientName = id

		s.broadcast <- &message
	}
}

// HandleMessages runs for the lifetime of the server and should only
// be called once. It continually drains the broadcast queue and sends
// messages to clients with open websocket connections.
func (s *WebSocketServer) HandleMessages() {
	handleMessage := func(msg *Message) {
		// Send the message to all Clients
		s.clientsMux.RLock()
		defer s.clientsMux.RUnlock()
		for id, client := range s.clients {
			var rendered []byte
			var err error
			if id == msg.ClientName {
				rendered, err = s.formatting.formatSentMessage(msg)
			} else {
				rendered, err = s.formatting.formatRecvMessage(msg)
			}
			if err != nil {
				slog.Error("Template execution failed", "error", err)
				continue
			}
			err = client.conn.WriteMessage(websocket.TextMessage, rendered)
			if err != nil {
				slog.Error("Write Error", "error", err)
			}
		}
	}
	for {
		handleMessage(<-s.broadcast)
	}
}

type clientConn struct {
	id   string
	conn *websocket.Conn
}

type templates struct {
	sentMessage *template.Template
	recvMessage *template.Template
}

func (t templates) formatSentMessage(msg *Message) ([]byte, error) {
	var renderedMessage bytes.Buffer
	err := t.sentMessage.Execute(&renderedMessage, msg)
	if err != nil {
		return nil, err
	}
	return renderedMessage.Bytes(), nil
}

func (t templates) formatRecvMessage(msg *Message) ([]byte, error) {
	var renderedMessage bytes.Buffer
	err := t.recvMessage.Execute(&renderedMessage, msg)
	if err != nil {
		return nil, err
	}
	return renderedMessage.Bytes(), nil
}
