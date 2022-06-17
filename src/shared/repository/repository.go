package repository

import (
	"context"

	"github.com/yanuar-nc/lineage/src/shared/domain"
)

type MessageBroker interface {
	Publish(ctx context.Context, topic string, event domain.EventMessage) error
}
