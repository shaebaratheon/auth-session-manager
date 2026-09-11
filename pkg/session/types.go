package session

import (
	"time"
)

type SessionStatus string

const (
	StatusActive  SessionStatus = "ACTIVE"
	StatusRevoked SessionStatus = "REVOKED"
	StatusExpired SessionStatus = "EXPIRED"
)

type UserSession struct {
	SessionID string            `json:"session_id"`
	UserID    string            `json:"user_id"`
	IPAddress string            `json:"ip_address"`
	UserAgent string            `json:"user_agent"`
	CreatedAt time.Time         `json:"created_at"`
	ExpiresAt time.Time         `json:"expires_at"`
	Status    SessionStatus     `json:"status"`
	Claims    map[string]string `json:"claims"`
}

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions

// Model expansion definitions
