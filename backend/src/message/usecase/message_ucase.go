package usecase

import (
	"context"
	"time"
)

type messageUsecase struct {
	contextTimeout time.Duration
}

func (m *messageUsecase) Store(c context.Context) (err error) {
	return
}
