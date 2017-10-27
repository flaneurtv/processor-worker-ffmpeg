package message

import "time"

const (
	LevelCritical = "critical"
	LevelError    = "error"
	LevelWarning  = "warning"
	LevelInfo     = "info"
	LevelDebug    = "debug"
)

type Error struct {
	Level      string    `json:"level,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	LogMessage string    `json:"log_message,omitempty"`
}

func NewError(level string, msg string) *Error {
	var e Error
	e.Level = level
	e.CreatedAt = time.Now()
	e.LogMessage = msg
	return &e
}
