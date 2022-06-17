package repository

import (
	"context"

	"github.com/yanuar-nc/lineage/src/domain"
)

// Repository interface
type Repository interface {
	SavePerson(ctx context.Context, req *domain.Person) error
}
