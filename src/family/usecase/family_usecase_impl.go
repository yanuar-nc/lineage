package usecase

import (
	"context"

	"github.com/yanuar-nc/lineage/src/family/domain"
	"github.com/yanuar-nc/lineage/src/family/repository"
)

// FamilyUsecaseImpl struct
type FamilyUsecaseImpl struct {
	repository repository.Family
}

// NewFamilyUsecaseImpl function
func NewFamilyUsecaseImpl(familyRepository repository.Family) *FamilyUsecaseImpl {
	return &FamilyUsecaseImpl{repository: familyRepository}
}

func (u *FamilyUsecaseImpl) Save(ctx context.Context, req domain.Family) error {

	err := u.repository.Save(ctx, &req)
	if err != nil {
		return err
	}

	return nil
}
