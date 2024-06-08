package main

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
	"k8s.io/klog/v2"
)

type WebSocketServer struct {
	id        string
	clients   map[*websocket.Conn]bool
	broadcast chan *Message
}

func NewWebSocket() *WebSocketServer {
	return &WebSocketServer{
		id:        uuid.New().String(),
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan *Message),
	}
}

func (s *WebSocketServer) HandleWebSocket(ctx *websocket.Conn) {
	// Register a new Client
	s.clients[ctx] = true
	defer func() {
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
		message.ClientName = s.id

		s.broadcast <- &message
	}
}

func (s *WebSocketServer) HandleMessages() {
	for {
		msg := <-s.broadcast

		// Send the message to all Clients
		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, getMessageTemplate(msg))
			if err != nil {
				klog.Errorf("Write  Error: %v ", err)
				client.Close()
				delete(s.clients, client)
			}

		}
	}
}

func getMessageTemplate(msg *Message) []byte {
	tmpl, err := template.ParseFiles("views/message.html")
	if err != nil {
		klog.Exitf("template parsing: %s", err)
	}

	// Render the template with the message as data.
	var renderedMessage bytes.Buffer
	err = tmpl.Execute(&renderedMessage, msg)
	if err != nil {
		klog.Exitf("template execution: %s", err)
	}

	return renderedMessage.Bytes()
}
