package migrate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidConnectionToDatabase(t *testing.T) {
	migr := NewMigrator("localhost:27017", "", "")
	err := migr.Connect()
	assert.NoError(t, err)
	migr.Disconnect()
}

func TestInvalidConnectionUrl(t *testing.T) {
	migr := NewMigrator("i-am:invalid", "", "")
	err := migr.Connect()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing uri")
}

func TestTimeoutConnectionToDatabase(t *testing.T) {
	migr := NewMigrator("localhost:9999", "", "")
	err := migr.Connect()
	assert.Error(t, err)
}
