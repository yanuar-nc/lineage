package database

import (
	"reflect"

	"github.com/google/uuid"
	"github.com/mindstand/gogm/v2"
)

type tdString string
type tdInt int

type Family struct {
	gogm.BaseNode

	Name     string `gogm:"name=name" json:"name"`
	Username string `gogm:"name=username" json:"username"`
	Uuid     string `gogm:"name=uuid"`

	Group    *Group    `gogm:"direction=outgoing;relationship=OWNER"`
	Children []*Person `gogm:"direction=incoming;relationship=CHILDREN"`
	Wife     []*Person `gogm:"direction=incoming;relationship=WIFE"`
	Husband  []*Person `gogm:"direction=incoming;relationship=HUSBAND"`
}

type Group struct {
	gogm.BaseNode

	ID   string `gogm:"name=id"`
	Name string `gogm:"name=name"`
	Uuid string `gogm:"name=uuid"`

	Families []*Family `gogm:"direction=incoming;relationship=OWNER"`
}

type Person struct {
	gogm.BaseNode

	ID   string `gogm:"name=id"`
	Name string `gogm:"name=name"`
	Uuid string `gogm:"name=uuid"`

	Husband  *Family `gogm:"direction=outgoing;relationship=HUSBAND"`
	Wife     *Family `gogm:"direction=outgoing;relationship=WIFE"`
	Children *Family `gogm:"direction=outgoing;relationship=CHILDREN"`
}

type ContingentUponEdge struct {
	gogm.BaseNode

	Start       *Group  `json:"group"`
	End         *Person `json:"person"`
	Criticality string  `gogm:"name=criticality"`
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

	pk := &gogm.PrimaryKeyStrategy{
		StrategyName: "UUID",
		DBName:       "uuid",
		FieldName:    "UUID",
		Type:         reflect.TypeOf(""),
		GenIDFunc: func() (id interface{}) {
			return uuid.New().String()
		},
	}

	// register all vertices and edges
	// this is so that GoGM doesn't have to do reflect processing of each edge in real time
	// use nil or gogm.DefaultPrimaryKeyStrategy if you only want graph ids
	// we are using the default key strategy since our vertices are using BaseNode
	_gogm, err := gogm.New(&config, pk, &Family{}, &Group{}, &Person{}, &ContingentUponEdge{})
	if err != nil {
		panic(err)
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
