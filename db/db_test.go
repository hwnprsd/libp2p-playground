package db

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewDB(t *testing.T) {
	tempFile, _ := ioutil.TempDir("", "testdb")
	defer os.Remove(tempFile)

	db, err := NewDB(tempFile)
	if err != nil {
		t.Fatal(err)
	}
	testKey := []byte("TEST_KEY")
	testValue := []byte("TEST_VALUE")

	db.Set(testKey, testValue)
	returnedValue := db.Get(testKey)
	require.Equal(t, testValue, returnedValue)
}
