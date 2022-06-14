package database

import (
	"time"

	"github.com/mindstand/gogm/v2"
)

type tdString string
type tdInt int

//structs for the example (can also be found in decoder_test.go)
type VertexA struct {
	// provides required node fields
	gogm.BaseNode

	TestField         string            `gogm:"name=test_field"`
	TestTypeDefString tdString          `gogm:"name=test_type_def_string"`
	TestTypeDefInt    tdInt             `gogm:"name=test_type_def_int"`
	MapProperty       map[string]string `gogm:"name=map_property;properties"`
	SliceProperty     []string          `gogm:"name=slice_property;properties"`
	SingleA           *VertexB          `gogm:"direction=incoming;relationship=test_rel"`
	ManyA             []*VertexB        `gogm:"direction=incoming;relationship=testm2o"`
	MultiA            []*VertexB        `gogm:"direction=incoming;relationship=multib"`
	SingleSpecA       *EdgeC            `gogm:"direction=outgoing;relationship=special_single"`
	MultiSpecA        []*EdgeC          `gogm:"direction=outgoing;relationship=special_multi"`
}

type VertexB struct {
	// provides required node fields
	gogm.BaseNode

	TestField  string     `gogm:"name=test_field"`
	TestTime   time.Time  `gogm:"name=test_time"`
	Single     *VertexA   `gogm:"direction=outgoing;relationship=test_rel"`
	ManyB      *VertexA   `gogm:"direction=outgoing;relationship=testm2o"`
	Multi      []*VertexA `gogm:"direction=outgoing;relationship=multib"`
	SingleSpec *EdgeC     `gogm:"direction=incoming;relationship=special_single"`
	MultiSpec  []*EdgeC   `gogm:"direction=incoming;relationship=special_multi"`
}

type EdgeC struct {
	// provides required node fields
	gogm.BaseNode

	Start *VertexA
	End   *VertexB
	Test  string `gogm:"name=test"`
}

func GetNeo4jConn(host string, port int, username, password string) (gogm.SessionV2, error) {
	// define your configuration
	config := gogm.Config{
		Host: host,
		Port: port,
		// deprecated in favor of protocol
		// IsCluster:                 false,
		Protocol:      "neo4j", //also supports neo4j+s, neo4j+ssc, bolt, bolt+s and bolt+ssc
		Username:      username,
		Password:      password,
		PoolSize:      50,
		IndexStrategy: gogm.VALIDATE_INDEX, //other options are ASSERT_INDEX and IGNORE_INDEX
		TargetDbs:     nil,
		// default logger wraps the go "log" package, implement the Logger interface from gogm to use your own logger
		Logger: gogm.GetDefaultLogger(),
		// define the log level
		LogLevel: "DEBUG",
		// enable neo4j go driver to log
		EnableDriverLogs: false,
		// enable gogm to log params in cypher queries. WARNING THIS IS A SECURITY RISK! Only use this when debugging
		EnableLogParams: false,
		// enable open tracing. Ensure contexts have spans already. GoGM does not make root spans, only child spans
		OpentracingEnabled: false,
		// specify the method gogm will use to generate Load queries
		// LoadStrategy: gogm.PATH_LOAD_STRATEGY // set to SCHEMA_LOAD_STRATEGY for schema-aware queries which may reduce load on the database
	}

	_gogm, err := gogm.New(&config, gogm.DefaultPrimaryKeyStrategy, &VertexA{}, &VertexB{}, &EdgeC{})
	if err != nil {
		return nil, err
	}

	//param is readonly, we're going to make stuff so we're going to do read write
	sess, err := _gogm.NewSessionV2(gogm.SessionConfig{AccessMode: gogm.AccessModeWrite})
	if err != nil {
		return nil, err
	}

	//close the session
	defer sess.Close()
	return sess, nil
}
