package main

import (
	"fmt"
	"net/http"

	"github.com/mindstand/gogm/v2"
	"github.com/yanuar-nc/lineage/config"

	jsonDelivery "github.com/yanuar-nc/lineage/src/delivery/json"
	neo4jRepository "github.com/yanuar-nc/lineage/src/repository/neo4j"
	"github.com/yanuar-nc/lineage/src/usecase"

	"github.com/labstack/echo/v4"
)

// EchoServer structure
type EchoServer struct {
	echoHandler *jsonDelivery.EchoHandler
	cfg         config.Config
}

// Run main function for serving echo http server
func (s *EchoServer) Run() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Up and running !!")
	})

	// family v1 route
	familyGroupV1 := e.Group("/")
	s.echoHandler.Mount(familyGroupV1)

	listenerPort := fmt.Sprintf(":%d", s.cfg.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}

// NewEchoServer function
func NewEchoServer(cfg config.Config, writeDb gogm.SessionV2, readDb gogm.SessionV2) (*EchoServer, error) {
	repositoryImpl := neo4jRepository.NewRepository(writeDb, readDb)

	u := usecase.NewUsecaseImplementation().PutRepository(repositoryImpl)
	echoHandler := jsonDelivery.NewEchoHandler(u)

	return &EchoServer{
		echoHandler: echoHandler,
		cfg:         cfg,
	}, nil
}
