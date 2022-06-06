package repository

import (
	"context"

	"github.com/yanuar-nc/go-boiler-plate/src/family/domain"
)

// Family interface
type Family interface {
	Save(ctx context.Context, data *domain.Family) error
}
