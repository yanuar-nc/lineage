package repository

import (
	"context"

	"github.com/yanuar-nc/go-boiler-plate/src/shared/domain"
)

type MessageBroker interface {
	Publish(ctx context.Context, topic string, event domain.EventMessage) error
}
