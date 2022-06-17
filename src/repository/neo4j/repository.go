package repository

import (
	"github.com/mindstand/gogm/v2"
)

// RepositoryNeo4j struct
type Repository struct {
	writeDb gogm.SessionV2
	readDb  gogm.SessionV2
}

// NewRepositoryNeo4j function
func NewRepository(writeDb, readDb gogm.SessionV2) *Repository {
	return &Repository{writeDb: writeDb, readDb: readDb}
}
