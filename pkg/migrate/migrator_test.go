package migrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultProtocolWhenMissing(t *testing.T) {
	migr := NewMigrator("localhost:27017", "", "")
	assert.Equal(t, "mongodb://localhost:27017", migr.url)
}

func TestDefaultProtocolWhenPresent(t *testing.T) {
	migr := NewMigrator("mongodb://localhost:27017", "", "")
	assert.Equal(t, "mongodb://localhost:27017", migr.url)
}

func TestDisconnectWithoutClient(t *testing.T) {
	migr := NewMigrator("mongodb://localhost:27017", "", "")
	err := migr.Disconnect()
	assert.NoError(t, err)
}
