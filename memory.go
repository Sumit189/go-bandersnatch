package main

type Message struct {
	Role    string // Role can be "user" or "assistant"
	Content string // The content of the message
}

type Memory struct {
	conversationHistory []Message
}

func (m *Memory) AddToMemory(role, content string) {
	m.conversationHistory = append(m.conversationHistory, Message{Role: role, Content: content})
}

func (m *Memory) GetMemory() []Message {
	return m.conversationHistory
}
