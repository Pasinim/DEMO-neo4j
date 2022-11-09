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
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"testing"
	"time"
)

type TestContainer struct {
	Instance testcontainers.Container
}

func NewTestContainer() *TestContainer {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := testcontainers.ContainerRequest{
		Image:        "neo4j:4.4.12-community",
		ExposedPorts: []string{"7687/tcp", "7474/tcp"},
		AutoRemove:   true,
		Env: map[string]string{
			"NEO4J_AUTH": "neo4j/demo", //psw: demo
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

	//time.Sleep(time.Minute * 5)
	return &TestContainer{
		Instance: neo4j,
	}
}

func (container *TestContainer) Port() int {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	p, err := container.Instance.MappedPort(ctx, "7687")
	//container.Instance.Host(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PORT:", p)
	return p.Int()
}

func (container *TestContainer) Host() string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	h, err := container.Instance.Host(ctx)
	//container.Instance.Host(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("HOST:", h)
	return h
}

func (container *TestContainer) ConnectionString() string {
	return fmt.Sprintf("neo4j://neo4j:demo@%s:%d/", container.Host(), container.Port())
}

func (container *TestContainer) Close(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	require.NoError(t, container.Instance.Terminate(ctx))
}
