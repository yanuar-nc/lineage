package usecase

import "github.com/yanuar-nc/lineage/src/family/repository"

type Usecase struct {
	FamilyRepository repository.Family
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (u *Usecase) PutFamilyRepository(repo repository.Family) *Usecase {
	u.FamilyRepository = repo
	return u
}
