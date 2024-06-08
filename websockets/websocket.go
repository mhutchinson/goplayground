package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"sync"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"k8s.io/klog/v2"
)

type WebSocketServer struct {
	formatting templates
	broadcast  chan *Message

	clientsMux sync.RWMutex
	clients    map[*websocket.Conn]bool
}

func NewWebSocket() *WebSocketServer {
	tmpl, err := template.ParseFiles("views/message.html")
	if err != nil {
		klog.Exitf("template parsing: %s", err)
	}

	return &WebSocketServer{
		formatting: templates{message: tmpl},
		broadcast:  make(chan *Message),
		clients:    make(map[*websocket.Conn]bool),
	}
}

// HandleWebSocket accepts a websocker connection and services it
// for the lifetime of the connection. This method will not return
// until the connection is closed.
func (s *WebSocketServer) HandleWebSocket(ctx *websocket.Conn) {
	id := uuid.New().String()
	klog.Infof("Registering client %s", id)
	// Register a new Client
	s.clientsMux.Lock()
	s.clients[ctx] = true
	s.clientsMux.Unlock()
	defer func() {
		klog.Infof("Removing client %s", id)
		s.clientsMux.Lock()
		defer s.clientsMux.Unlock()
		delete(s.clients, ctx)
		ctx.Close()
	}()

	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			klog.Error("Read Error:", err)
			break
		}

		// send the message to the broadcast channel
		klog.Info(string(msg))
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			klog.Exit("Error Unmarshalling", err)
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
		rendered := s.formatting.formatMessage(msg)
		// Send the message to all Clients
		s.clientsMux.RLock()
		defer s.clientsMux.RUnlock()
		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, rendered)
			if err != nil {
				klog.Errorf("Write  Error: %v ", err)
			}
		}

	}
	for {
		handleMessage(<-s.broadcast)
	}
}

type templates struct {
	message *template.Template
}

func (t templates) formatMessage(msg *Message) []byte {
	// Render the template with the message as data.
	var renderedMessage bytes.Buffer
	err := t.message.Execute(&renderedMessage, msg)
	if err != nil {
		klog.Exitf("template execution: %s", err)
	}

	return renderedMessage.Bytes()
}
