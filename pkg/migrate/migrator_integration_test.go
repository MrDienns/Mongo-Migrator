package migrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidConnectionToDatabase(t *testing.T) {
	migr := NewMigrator("localhost:27017", "", "", "")
	err := migr.Connect()
	assert.NoError(t, err)
	migr.Disconnect()
}

func TestInvalidConnectionUrl(t *testing.T) {
	migr := NewMigrator("i-am:invalid", "", "", "")
	err := migr.Connect()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing uri")
}

func TestTimeoutConnectionToDatabase(t *testing.T) {
	migr := NewMigrator("localhost:9999", "", "", "")
	err := migr.Connect()
	assert.Error(t, err)
}

func TestCollectionNameDiscovery(t *testing.T) {
	migr := migrator(t)
	exp := []string{"a", "b", "c", "d", "e"}
	act, err := migr.DocumentIds("_integration_test_collection")
	assert.NoError(t, err)
	assert.Equal(t, exp, act)
	reset(migr, "_integration_test_collection")
}

func migrator(t *testing.T) *Migrator {
	migr := NewMigrator("localhost:27017", "", "", "_integration_test_database")
	err := migr.Connect()
	assert.NoError(t, err)
	coll := migr.client.Database(migr.database).Collection("_integration_test_collection")
	coll.InsertMany(migr.ctx, mockDocuments())
	return migr
}

func mockDocuments() []interface{} {
	res := make([]interface{}, 5)
	docIds := []string{"a", "b", "c", "d", "e"}
	for i, id := range docIds {
		m := make(map[string]interface{})
		m["_id"] = id
		res[i] = m
	}
	return res
}

func reset(m *Migrator, coll string) {
	m.client.Database(m.database).Collection(coll).Drop(m.ctx)
}
