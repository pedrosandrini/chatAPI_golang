package entity

import (
	"errors"

	"github.com/google/uuid"
)

type ChatConfig struct {
	Model            *Model
	Temperature      float32  // 0.0 to 1.0
	TopP             float32  // 0.0 to 1.0
	N                int      // Number of messages to generate
	Stop             []string // lsit of tokens to stop on
	MaxTokens        int      // number of tokens to generate
	PresencePenalty  float32  // -2.0 to 2.0
	FrequencyPenalty float32  // -2.0 to 2.0
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}

func NewChat(userID string, initialSystemMessage *Message, chatConfig *ChatConfig) (*Chat, error) {
	chat := &Chat{
		ID:                   uuid.New().String(),
		UserID:               userID,
		InitialSystemMessage: initialSystemMessage,
		Status:               "active",
		Config:               chatConfig,
		TokenUsage:           0,
	}

	chat.AddMessage(initialSystemMessage)

	if err := chat.Validate(); err != nil {
		return nil, err
	}

	return chat, nil
}

func (c *Chat) Validate() error {
	if c.UserID == "" {
		return errors.New("user id is empty")
	}

	if c.Status != "active" && c.Status != "ended" {
		return errors.New("invalid status")
	}

	if c.Config.Temperature < 0 || c.Config.Temperature > 1 {
		return errors.New("invalid temerature")
	}
	// More validations for config

	return nil
}

func (c *Chat) AddMessage(m *Message) error {
	if c.Status == "ended" {
		return errors.New("chat is ended. no more messages allowed")
	}

	for {
		if c.Config.Model.GetMaxtokens() >= m.GetQtdTokens()+c.TokenUsage {
			c.Messages = append(c.Messages, m)
			c.RefreshTokenUsage()
			break
		}
		c.ErasedMessages = append(c.ErasedMessages, c.Messages[0])
		c.Messages = c.Messages[1:]
		c.RefreshTokenUsage()
	}

	return nil
}

func (c *Chat) GetMessages() []*Message {
	return c.Messages
}

func (c *Chat) CountMessages() int {
	return len(c.Messages)
}

func (c *Chat) End() {
	c.Status = "ended"
}

func (c *Chat) RefreshTokenUsage() {
	c.TokenUsage = 0
	for m := range c.Messages {
		c.TokenUsage += c.Messages[m].GetQtdTokens()
	}
}
