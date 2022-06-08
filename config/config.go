package config

import (
	"errors"
	"os"
	"strconv"

	dotenv "github.com/joho/godotenv"
	"github.com/yanuar-nc/lineage/helper"
)

type Config struct {
	// Development env checking, this env for debug purpose
	Development string

	// HTTPPort config
	HTTPPort uint16

	// GRPCPort config
	GRPCPort string

	Neo4jDB ConfigNeo4jDB
}

// Load function will load all config from environment variable
func Load() (Config, error) {

	var (
		c Config
	)

	errCust := helper.ErrorCustom{}

	// load .env
	errCust.Append(dotenv.Load(".env")).
		Append(c.setDevelopmentMode()).
		Append(c.setHTTPPort()).
		Append(c.setNeo4j())

	return c, errCust.Message()
}

func (c *Config) setDevelopmentMode() error {
	// load .env
	err := dotenv.Load(".env")
	if err != nil {
		return errors.New(".env is not loaded properly")
	}

	development, ok := os.LookupEnv("DEVELOPMENT")
	if !ok {
		return errors.New("DEVELOPMENT env is not loaded")
	}

	c.Development = development
	return nil
}

func (c *Config) setHTTPPort() error {
	httpPortStr, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		return errors.New("HTTP_PORT env is not loaded")
	}

	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		return errors.New("HTTP_PORT env is not valid")
	}

	// set http port
	c.HTTPPort = uint16(httpPort)
	return nil
}

func (c *Config) setNeo4j() error {
	errCust := helper.ErrorCustom{}
	neo4jHost, ok := os.LookupEnv("NEO4J_DB_HOST")
	if !ok {
		errCust.Append(errors.New("NEO4J_DB_HOST env is not loaded"))
	}
	c.Neo4jDB.Host = neo4jHost

	neo4jPort, ok := os.LookupEnv("NEO4J_DB_PORT")
	if !ok {
		errCust.Append(errors.New("NEO4J_DB_PORT env is not loaded"))
	}
	c.Neo4jDB.Port = neo4jPort

	neo4jUsername, ok := os.LookupEnv("NEO4J_DB_NAME")
	if !ok {
		errCust.Append(errors.New("NEO4J_DB_NAME env is not loaded"))
	}
	c.Neo4jDB.Username = neo4jUsername

	neo4jPassword, ok := os.LookupEnv("NEO4J_DB_PASSWORD")
	if !ok {
		errCust.Append(errors.New("NEO4J_DB_PASSWORD env is not loaded"))
	}
	c.Neo4jDB.Password = neo4jPassword

	return errCust.Message()
}
