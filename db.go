package orm

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/go-labx/orm/dialect"
	"github.com/go-labx/orm/session"

	"github.com/go-labx/orm/logger"
)

type DB struct {
	debug      bool
	db         *sql.DB
	dialect    dialect.Dialect
	dsn        string
	dataSource *DataSource
}

// NewDB creates a new ORM DB instance
func NewDB(d *DataSource) (*DB, error) {
	dsn := d.DSN()
	db, err := sql.Open(string(d.Driver), dsn)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		logger.Error(err)
		return nil, err
	}

	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(string(d.Driver))
	if !ok {
		err = fmt.Errorf("dialect %s Not Found", d.Driver)
		logger.Error(err)
		return nil, err
	}

	logger.Info("Connect database success")
	return &DB{
		db:         db,
		dialect:    dial,
		dsn:        dsn,
		dataSource: d,
	}, nil
}

// NewSession creates a new session with the current database connection
func (d *DB) NewSession() *session.Session {
	return session.New(d.db, d.dialect)
}

// EnableDebug sets the debug flag to true
func (d *DB) EnableDebug() {
	d.debug = true
}

// DisableDebug sets the debug flag to false
func (d *DB) DisableDebug() {
	d.debug = false
}

// Close closes the connection to the database
func (d *DB) Close() error {
	return d.db.Close()
}

// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle
// before being closed.
func (d *DB) SetConnMaxIdleTime(t time.Duration) {
	d.db.SetConnMaxIdleTime(t)
}

// SetConnMaxLifetime sets the maximum amount of time a connection may be reused
// for before being closed.
func (d *DB) SetConnMaxLifetime(t time.Duration) {
	d.db.SetConnMaxLifetime(t)
}

// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
func (d *DB) SetMaxIdleConns(n int) {
	d.db.SetMaxIdleConns(n)
}

// SetMaxOpenConns sets the maximum number of open connections to the database.
func (d *DB) SetMaxOpenConns(n int) {
	d.db.SetMaxOpenConns(n)
}

// Stats returns database statistics.
func (d *DB) Stats() sql.DBStats {
	return d.db.Stats()
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (d *DB) Ping() error {
	return d.db.Ping()
}

// PingContext verifies a connection to the database is still alive, establishing a connection if necessary.
func (d *DB) PingContext(ctx context.Context) error {
	return d.db.PingContext(ctx)
}

// Driver returns the database's underlying driver.
func (d *DB) Driver() driver.Driver {
	return d.db.Driver()
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d *DB) Exec(query string, args ...any) (sql.Result, error) {
	return d.NewSession().Raw(query, args...).Exec()
}

// ExecContext executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (d *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return d.NewSession().Raw(query, args...).ExecContext(ctx)
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (d *DB) Query(query string, args ...any) (*sql.Rows, error) {
	return d.NewSession().Raw(query, args...).Query()
}

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (d *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return d.NewSession().Raw(query, args...).QueryContext(ctx)
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (d *DB) QueryRow(query string, args ...any) *sql.Row {
	return d.NewSession().Raw(query, args...).QueryRow()
}

// QueryRowContext executes a query that is expected to return at most one row.
// QueryRowContext always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the rest.
func (d *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return d.NewSession().Raw(query, args...).QueryRowContext(ctx)
}

// Conn returns a single connection by either opening a new connection
// or returning an existing connection from the connection pool. Conn will
// block until either a connection is returned or ctx is canceled.
// Queries run on the same Conn will be run in the same database session.
//
// Every Conn must be returned to the database pool after use by
// calling Conn.Close.
func (d *DB) Conn(ctx context.Context) (*sql.Conn, error) {
	return d.db.Conn(ctx)
}

// IsTableExist checks if a table with the given name exists in the database.
// It returns true if the table exists, false otherwise.
func (d *DB) IsTableExist(name string) bool {
	if name == "" {
		return false
	}

	sqlStat := d.dialect.IsTableExistSQL(name)

	row := d.QueryRow(sqlStat)
	var tableName string
	_ = row.Scan(&tableName)

	return tableName == name
}

func (d *DB) DropTable(name string) bool {
	if name == "" {
		return false
	}

	sqlStat := d.dialect.DropTableSQL(name)

	_, err := d.Exec(sqlStat)
	if err != nil {
		logger.Error(err)
		return false
	}

	return true
}

func (d *DB) Version() (string, error) {
	sqlStat := d.dialect.VersionSQL()

	var version string
	row := d.QueryRow(sqlStat)
	err := row.Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}

func (d *DB) DataSourceName() string {
	return d.dsn
}

func (d *DB) DriverName() string {
	return string(d.dataSource.Driver)
}
