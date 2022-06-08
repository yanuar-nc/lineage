package main

import (
	"os"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/yanuar-nc/golang/helper"
	"github.com/yanuar-nc/lineage/config"
	"github.com/yanuar-nc/lineage/config/database"
)

func main() {

	// call config.Load() before start up
	cfg, err := config.Load()
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "load_config")
		os.Exit(1)
	}

	neo4jDB, err := database.GetNeo4jConn(cfg.Neo4jDB.Host, cfg.Neo4jDB.Port, cfg.Neo4jDB.Username, cfg.Neo4jDB.Password)
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "neo4jDB")
		os.Exit(1)
	}

	echoServer, err := NewEchoServer(cfg, neo4jDB, neo4jDB)
	if err != nil {
		helper.Log(log.FatalLevel, err.Error(), "Main", "echo_server")
		os.Exit(1)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		echoServer.Run()
	}()

	// Wait All services to end
	wg.Wait()
}
