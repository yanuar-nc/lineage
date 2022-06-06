package main

import (
	"fmt"
	"net/http"

	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/yanuar-nc/lineage/config"

	familyDeliveryJson "github.com/yanuar-nc/lineage/src/family/delivery/json"
	familyRepository "github.com/yanuar-nc/lineage/src/family/repository"
	familyUsecasePackage "github.com/yanuar-nc/lineage/src/family/usecase"

	"github.com/yanuar-nc/lineage/src/shared/usecase"

	"github.com/labstack/echo"
)

// EchoServer structure
type EchoServer struct {
	familyEchoHandler *familyDeliveryJson.EchoHandler
}

// Run main function for serving echo http server
func (s *EchoServer) Run() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Up and running !!")
	})

	// family v1 route
	familyGroupV1 := e.Group("/family")
	s.familyEchoHandler.Mount(familyGroupV1)

	listenerPort := fmt.Sprintf(":%d", config.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}

// NewEchoServer function
func NewEchoServer(writeDb, readDb *neo4j.Session) (*EchoServer, error) {
	familyRepositoryImpl := familyRepository.NewFamilyRepositoryGorm(writeDb)

	familyUsecase := familyUsecasePackage.NewFamilyUsecaseImpl(familyRepositoryImpl)

	familyEchoHandler := familyDeliveryJson.NewEchoHandler(familyUsecase)

	u := usecase.NewUsecase().PutFamilyRepository(familyRepositoryImpl)
	_ = u
	return &EchoServer{
		familyEchoHandler: familyEchoHandler,
	}, nil
}
