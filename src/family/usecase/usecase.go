package usecase

import (
	"context"

	"github.com/yanuar-nc/lineage/src/family/domain"
)

// FamilyUsecase interface
type FamilyUsecase interface {
	Save(ctx context.Context, param domain.Family) error
}
