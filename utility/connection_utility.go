package utility

//in seguito le chiamate da fare nei test per connettersi

// usage:
// testDB := testhelpers.NewTestDatabase(t)
// defer testDB.Close(t)
// println(testDB.ConnectionString(t))

import (
	"context"
	"fmt"
	_ "github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"testing"
	"time"
)

type TestContainer struct {
	Instance testcontainers.Container
}

/*
*
https://hub.docker.com/_/neo4j/
Mi devo connettere a questo docker
*/

func NewTestContainer() *TestContainer {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := testcontainers.ContainerRequest{
		Image:        "neo4j:latest",
		ExposedPorts: []string{"7474/tcp", "7687/tcp"},
		AutoRemove:   true,
		Env: map[string]string{
			"NEO4J_AUTH": "neo4j/demo",
		},
		WaitingFor: wait.ForListeningPort("7687"),
	}
	neo4j, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sleeppooo")
	time.Sleep(time.Minute * 5)
	return &TestContainer{
		Instance: neo4j,
	}
}

func (db *TestContainer) Port() int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := db.Instance.MappedPort(ctx, "7474")
	if err != nil {
		log.Fatal(err)
	}
	//require.NoError(t, err) serve se fa asser nil?
	return p.Int()
}

//
//func (db *TestContainer) ConnectionString() string {
//	return fmt.Sprintf("postgres://demo:demo@127.0.0.1:%d/demo?sslmode=disable", db.Port())
//}

func (db *TestContainer) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, db.Instance.Terminate(ctx))
}
