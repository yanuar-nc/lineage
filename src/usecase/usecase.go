package usecase

import "github.com/yanuar-nc/lineage/src/repository"

// UsecaseImplementation struct
type UsecaseImplementation struct {
	repository repository.Repository
}

// NewUsecaseImplementation function
func NewUsecaseImplementation() *UsecaseImplementation {
	return &UsecaseImplementation{}
}

func (u *UsecaseImplementation) PutRepository(repo repository.Repository) *UsecaseImplementation {
	u.repository = repo
	return u
}
