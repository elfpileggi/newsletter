package domain

import (
	"context"
	"time"
)

type Message struct {
	ID        int64     `json:"id"`
	Subject   string    `json:"subject" validate:"required"`
	Detail    string    `json:"detail" validate:"required"`
	Timestamp time.Time `json:"timestap"`
}

type MessageUsecase interface {
	Store(context.Context, *Message) error
}
