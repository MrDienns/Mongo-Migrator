package migrate

import (
	"context"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Migrator struct is used to build an object containing the relevant information
// required to perform the migration.
type Migrator struct {
	url      string
	username string
	password string
	database string
	ctx      context.Context
	client   *mongo.Client
}

// NewMigrator accepts a number of parameters and constructs a new Migrator object
// and returns the pointer.
func NewMigrator(url string, username string, password string, database string) *Migrator {
	if !strings.HasPrefix(url, "mongodb://") {
		url = "mongodb://" + url
	}
	return &Migrator{
		url:      url,
		username: username,
		password: password,
		database: database,
	}
}

// Connect will build a mongo.Client using the structs parameters. If constructing the
// client resulted in an error, it will be returned.  If it succeeded, it will try to
// connect to the MongoDB instance. If it failed, an error is returned. It will then
// try to ping the server. If it failed, an error is returned. If all goes well, nil is
// returned to indicate that there was no error.
func (m *Migrator) Connect() error {
	ops := options.Client().ApplyURI(m.url)
	if m.username != "" && m.password != "" {
		ops.SetAuth(options.Credential{Username: m.username, Password: m.password})
	}
	client, err := mongo.NewClient(ops)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	m.ctx = ctx
	m.client = client
	return nil
}

// Disconnect will disconnect the client from the MongoDB instance.
func (m *Migrator) Disconnect() error {
	if m.client == nil {
		return nil
	}
	err := m.client.Disconnect(m.ctx)
	if err != nil {
		return err
	}
	return nil
}
