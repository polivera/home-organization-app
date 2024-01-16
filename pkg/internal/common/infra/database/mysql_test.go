package database

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMysqlDB_fillConnectionData(t *testing.T) {
	t.Run("empty port will default in port 3306", func(t *testing.T) {
		mdb := &mysqlDB{}
		err := mdb.fillConnectionData()
		assert.NoError(t, err)
		assert.Equal(t, uint16(3306), mdb.port)
	})

	t.Run("environment port will be set on the connection data", func(t *testing.T) {
		var err error
		err = os.Setenv(databasePort, "5934")
		assert.NoError(t, err)

		mdb := &mysqlDB{}
		err = mdb.fillConnectionData()
		assert.NoError(t, err)
		assert.Equal(t, uint16(5934), mdb.port)
	})

	t.Run("error setting port value", func(t *testing.T) {
		var err error
		err = os.Setenv(databasePort, "-321")
		assert.NoError(t, err)

		mdb := &mysqlDB{}
		err = mdb.fillConnectionData()
		assert.Error(t, err)
	})

	t.Run("environment port will be set on the connection data", func(t *testing.T) {
		var err error
		err = os.Setenv(databaseHost, "mock-host")
		assert.NoError(t, err)
		err = os.Setenv(databasePort, "1234")
		assert.NoError(t, err)
		err = os.Setenv(databaseUser, "mock-user")
		assert.NoError(t, err)
		err = os.Setenv(databasePass, "mock-pass")
		assert.NoError(t, err)
		err = os.Setenv(databaseName, "mock-db-name")
		assert.NoError(t, err)

		mdb := &mysqlDB{}
		err = mdb.fillConnectionData()
		assert.NoError(t, err)
		assert.Equal(t, "mock-host", mdb.host)
		assert.Equal(t, uint16(1234), mdb.port)
		assert.Equal(t, "mock-user", mdb.user)
		assert.Equal(t, "mock-pass", mdb.pass)
		assert.Equal(t, "mock-db-name", mdb.dbname)
	})
}

func TestMysqlDB_buildConnectionString(t *testing.T) {
	t.Run("set connection string with environment variables", func(t *testing.T) {
		var err error
		err = os.Setenv(databaseHost, "mock-host")
		assert.NoError(t, err)
		err = os.Setenv(databasePort, "1234")
		assert.NoError(t, err)
		err = os.Setenv(databaseUser, "mock-user")
		assert.NoError(t, err)
		err = os.Setenv(databasePass, "mock-pass")
		assert.NoError(t, err)
		err = os.Setenv(databaseName, "mock-db-name")
		assert.NoError(t, err)

		mdb := &mysqlDB{}
		err = mdb.fillConnectionData()
		assert.NoError(t, err)
		actual := mdb.buildConnectionString()
		assert.Equal(t, "mock-user:mock-pass@tcp(mock-host:1234)/mock-db-name?tls=skip-verify", actual)
	})
}
