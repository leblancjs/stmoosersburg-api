package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"
	"testing"
)

var mConn = &mockConnection{}
var mConnector = &mockConnector{mConn}
var mDriver = &mockDriver{connector: mConnector}

func init() {
	sql.Register("postgres", mDriver)
}

func TestOpeningPostgresDatabase(t *testing.T) {
	t.Run("fails when driver fails to open connection", func(t *testing.T) {
		mDriver.failOnOpen = true

		db := Postgres{}

		if err := db.Open(); err == nil {
			t.Fail()
		}
		defer db.Close()

		mDriver.failOnOpen = false
	})

	t.Run("fails when driver fails to ping", func(t *testing.T) {
		mDriver.connector.conn.failOnPing = true

		db := Postgres{}

		if err := db.Open(); err == nil {
			t.Fail()
		}
		defer db.Close()

		mDriver.connector.conn.failOnPing = false
	})

	t.Run("sets the database handle when all is well", func(t *testing.T) {
		db := Postgres{}

		db.Open()
		defer db.Close()

		if db.DB == nil {
			t.Fail()
		}
	})
}

func TestClosingPostgresDatabase(t *testing.T) {
	t.Run("does nothing when no connection was opened", func(t *testing.T) {
		db := Postgres{}

		if err := db.Close(); err != nil {
			t.Fail()
		}
	})

	t.Run("fails when driver fails to close connection", func(t *testing.T) {
		mDriver.connector.conn.failOnClose = true

		db := Postgres{}
		db.Open()

		if err := db.Close(); err == nil {
			t.Fail()
		}

		mDriver.connector.conn.failOnClose = false
	})

	t.Run("resets the database handle when all is well", func(t *testing.T) {
		db := Postgres{}
		db.Open()

		if err := db.Close(); err != nil {
			t.Fail()
		}

		if db.DB != nil {
			t.Fail()
		}
	})
}

func TestBuildingPostgresDataSourceName(t *testing.T) {
	conf := Config{
		Host:     "host",
		Port:     "1234",
		User:     "user",
		Password: "password",
		Name:     "name",
	}

	t.Run("returns data source name without password when it is blank in configuration", func(t *testing.T) {
		expectedDataSourceName := fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s",
			conf.Host,
			conf.Port,
			conf.Name,
			conf.User,
		)

		db := Postgres{
			db: db{conf},
		}

		dataSourceName := db.buildDataSourceName()
		if strings.Compare(expectedDataSourceName, dataSourceName) != 0 {
			t.Fail()
		}
	})

	t.Run("returns data source name with password", func(t *testing.T) {
		expectedDataSourceName := fmt.Sprintf(
			"host=%s port=%s dbname=%s user=%s password=%s",
			conf.Host,
			conf.Port,
			conf.Name,
			conf.User,
			conf.Password,
		)

		db := Postgres{
			db: db{conf},
		}

		dataSourceName := db.buildDataSourceName()
		if strings.Compare(expectedDataSourceName, dataSourceName) != 0 {
			t.Fail()
		}
	})
}

type mockConnection struct {
	failOnClose bool
	failOnPing  bool
}

func (mc *mockConnection) Prepare(query string) (driver.Stmt, error) {
	return nil, nil
}

func (mc *mockConnection) Close() error {
	if mc.failOnClose {
		return fmt.Errorf("failed to close connectin")
	}

	return nil
}

func (mc *mockConnection) Begin() (driver.Tx, error) {
	return nil, nil
}

func (mc *mockConnection) Ping(_ context.Context) error {
	if mc.failOnPing {
		return fmt.Errorf("failed to ping")
	}

	return nil
}

type mockConnector struct {
	conn *mockConnection
}

func (mc *mockConnector) Connect(_ context.Context) (driver.Conn, error) {
	return mc.conn, nil
}

func (mc *mockConnector) Driver() driver.Driver {
	return nil
}

type mockDriver struct {
	failOnOpen bool
	connector  *mockConnector
}

func (md *mockDriver) Open(name string) (driver.Conn, error) {
	if md.failOnOpen {
		return nil, fmt.Errorf("failed to open connection")
	}

	return nil, nil
}

func (md *mockDriver) OpenConnector(name string) (driver.Connector, error) {
	if md.failOnOpen {
		return nil, fmt.Errorf("failed to open connector")
	}

	return md.connector, nil
}
