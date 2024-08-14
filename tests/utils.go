package tests

import (
	"context"
	"log"
	"testing"
	"time"
	"toasty/core"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainerMeta struct {
	Config    core.ConnectionConfig
	Container *postgres.PostgresContainer
	Context   context.Context
}

func SetupPostgresContainer(t *testing.T) PostgresContainerMeta {
	dbName := "test"
	dbUser := "admin"
	dbPassword := "secret"

	ctx := context.Background()

	postgresContainer, err := postgres.Run(ctx,
		"docker.io/postgres:16-alpine",
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		log.Fatalf("failed to start container: %s", err)
		panic(err)
	}

	mappedPort, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("failed to get mapped port: %s", err)
		panic(err)
	}

	config := core.ConnectionConfig{
		Host:     "localhost",
		Port:     mappedPort.Int(),
		Database: dbName,
		Username: dbUser,
		Password: dbPassword,
	}

	return PostgresContainerMeta{
		Container: postgresContainer,
		Context:   ctx,
		Config:    config,
	}
}

func Setup(t *testing.T, numberOfDBs int) []PostgresContainerMeta {
	var containerMetaList []PostgresContainerMeta

	log.Println("setting up test containers")
	for _ = range numberOfDBs {
		containerMeta := SetupPostgresContainer(t)
		containerMetaList = append(containerMetaList, containerMeta)
	}

	return containerMetaList
}

func TeardownPostgresContainer(containerMeta PostgresContainerMeta) {
	if err := containerMeta.Container.Terminate(containerMeta.Context); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
}

func Teardown(containerMetas []PostgresContainerMeta) {
	log.Println("tearing down test containers")
	for _, containerMeta := range containerMetas {
		TeardownPostgresContainer(containerMeta)
	}
}
