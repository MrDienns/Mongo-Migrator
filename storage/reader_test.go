package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadingOfValidStorageFolder(t *testing.T) {
	act, err := ReadStorage("reader_test/valid")
	assert.NoError(t, err)
	exp := &Storage{
		Documents: make(map[string]map[string][]byte, 1),
	}
	exp.Documents["coll"] = make(map[string][]byte, 2)
	exp.Documents["coll"]["one.json"] = []byte(`{"hello": "world"}`)
	exp.Documents["coll"]["two.json"] = []byte(`{"foo": "bar"}`)
	assert.Equal(t, exp, act)
}
