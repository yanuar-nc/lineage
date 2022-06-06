package database

import "github.com/neo4j/neo4j-go-driver/neo4j"

func GetNeo4jConn(uri, username, password string) (neo4j.Session, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	defer driver.Close()

	session, err := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	if err != nil {
		return nil, err
	}
	defer session.Close()
	return session, nil
}
