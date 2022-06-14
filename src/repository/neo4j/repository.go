package repository

import "github.com/mindstand/gogm/v2"

// RepositoryNeo4j struct
type RepositoryNeo4j struct {
	writeDb gogm.SessionV2
	readDb  gogm.SessionV2
}

// NewRepositoryNeo4j function
func NewRepositoryNeo4j(writeDb, readDb gogm.SessionV2) *RepositoryNeo4j {
	return &RepositoryNeo4j{writeDb: writeDb, readDb: readDb}
}
