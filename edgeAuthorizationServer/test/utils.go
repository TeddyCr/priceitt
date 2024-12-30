package test

import (
	"context"
	"log"
	"time"

	"github.com/TeddyCr/priceitt/models"
	"github.com/TeddyCr/priceitt/utils/database"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func SetUp(handlers []ITestHandler) {
	// Set up tests
	for _, handler := range handlers {
		log.Printf("Setting up test: %v", handler.GetHandlerName())
		handler.SetUp()
	}
}

func TearDown(handlers []ITestHandler) {
	// Tear down tests
	for _, handler := range handlers {
		log.Printf("Tearing down test: %v", handler.GetHandlerName())
		handler.TearDown()
	}
}


type ITestHandler interface {
	SetUp()
	TearDown()
	GetHandlerName() string
}

// PostgresTestHandler is a test handler for postgres
func DefaultPostgresTestHandler() *PostgresTestHandler {
	return &PostgresTestHandler{
		userName: "user",
		password: "password",
		databaseName: "edge_authorization_server",
		psqlContainer: nil,
	}
}

type PostgresTestHandler struct {
	userName string
	password string
	databaseName string
	psqlContainer testcontainers.Container
}

func (p *PostgresTestHandler) SetUp() {
	ctx := context.Background()
	psqlContainer, err := postgres.Run(
		ctx,
		"postgres:17.2",
		testcontainers.WithHostConfigModifier(
			func(hc *container.HostConfig) {
				hc.PortBindings = map[nat.Port][]nat.PortBinding{
					"5432/tcp": {{HostIP: "0.0.0.0", HostPort: "54321"}},
				}
			},
		),
 		postgres.WithDatabase(p.databaseName),
		postgres.WithUsername(p.userName),
		postgres.WithPassword(p.password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		log.Fatalf("Could not start postgres container: %v", err)
		if err := testcontainers.TerminateContainer(psqlContainer); err != nil {
			log.Fatalf("Could not terminate postgres container: %v", err)
		}
	}
	p.psqlContainer = psqlContainer
}

func (p *PostgresTestHandler) TearDown() {
	if err := testcontainers.TerminateContainer(p.psqlContainer); err != nil {
		log.Fatalf("Could not terminate postgres container: %v", err)
	}
	p.psqlContainer = nil
}

func (p PostgresTestHandler) GetHandlerName() string {
	return "PostgresTestHandler"
}

func GetDatabaseConnection() *sqlx.DB {
	config := models.DatabaseConfig{
		DriverClass: "postgres",
		ConnectionString: "postgresql://user:password@localhost:5432/edge_authorization_server?sslmode=disable",
	}
	db := database.Connect(config)
	return db
}
