package repository

import "github.com/neo4j/neo4j-go-driver/neo4j"

// RepositoryNeo4j struct
type RepositoryNeo4j struct {
	writeDb neo4j.Session
	readDb  neo4j.Session
}

// NewRepositoryNeo4j function
func NewRepositoryNeo4j(writeDb, readDb neo4j.Session) *RepositoryNeo4j {
	return &RepositoryNeo4j{writeDb: writeDb, readDb: readDb}
}
