package database

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlDB struct {
	db     *sql.DB
	host   string
	port   uint16
	user   string
	pass   string
	dbname string
}

func NewMySQLConnection() Connection {
	return &mysqlDB{}
}

func (mdb *mysqlDB) Open() error {
	var err error
	if err = mdb.fillConnectionData(); err == nil {
		mdb.db, err = sql.Open("mysql", mdb.buildConnectionString())
	}
	return err
}

func (mdb *mysqlDB) Connect(ctx context.Context) (*sql.Conn, error) {
	return mdb.db.Conn(ctx)
}

func (mdb *mysqlDB) GetDB() *sql.DB {
	return mdb.db
}

func (mdb *mysqlDB) fillConnectionData() error {
	mdb.port = 3306
	strPort, isPortSet := os.LookupEnv(databasePort)
	if isPortSet {
		uIntPort, err := strconv.ParseUint(strPort, 10, 16)
		if err != nil {
			return err
		}
		mdb.port = uint16(uIntPort)
	}

	mdb.host = os.Getenv(databaseHost)
	mdb.user = os.Getenv(databaseUser)
	mdb.pass = os.Getenv(databasePass)
	mdb.dbname = os.Getenv(databaseName)

	return nil
}

func (mdb *mysqlDB) buildConnectionString() string {
	options := url.Values{}
	options.Set("tls", "skip-verify")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		mdb.user,
		mdb.pass,
		mdb.host,
		mdb.port,
		mdb.dbname,
		options.Encode(),
	)
}
