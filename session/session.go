package session

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-labx/orm/logger"
)

// Session struct holds the database connection and the SQL query to be executed.
type Session struct {
	db      *sql.DB         // Database connection
	sql     strings.Builder // SQL query
	sqlArgs []interface{}   // Arguments for the SQL query
}

// New creates a new Session with the provided database connection.
func New(db *sql.DB) *Session {
	return &Session{db: db}
}

// Clear resets the SQL query and its arguments in the Session.
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlArgs = nil
}

// DB returns the database connection from the Session.
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw sets the SQL query and its arguments in the Session.
func (s *Session) Raw(sql string, args ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlArgs = append(s.sqlArgs, args...)
	return s
}

// Exec executes the SQL query in the Session and returns the result.
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	logger.Info(s.sql.String(), s.sqlArgs)
	if result, err = s.db.Exec(s.sql.String(), s.sqlArgs...); err != nil {
		logger.Error(err.Error())
	}
	return result, nil
}

// ExecContext executes the SQL query in the Session with a context and returns the result.
func (s *Session) ExecContext(ctx context.Context) (result sql.Result, err error) {
	defer s.Clear()
	logger.Info(s.sql.String(), s.sqlArgs)
	if result, err = s.db.ExecContext(ctx, s.sql.String(), s.sqlArgs...); err != nil {
		logger.Error(err.Error())
	}
	return result, nil
}

// Query executes the SQL query in the Session and returns the rows.
func (s *Session) Query() (result *sql.Rows, err error) {
	defer s.Clear()
	logger.Info(s.sql.String(), s.sqlArgs)
	if result, err = s.db.Query(s.sql.String(), s.sqlArgs...); err != nil {
		logger.Error(err.Error())
	}
	return result, nil
}

// QueryContext executes the SQL query in the Session with a context and returns the rows.
func (s *Session) QueryContext(ctx context.Context) (result *sql.Rows, err error) {
	defer s.Clear()
	logger.Info(s.sql.String(), s.sqlArgs)
	if result, err = s.db.QueryContext(ctx, s.sql.String(), s.sqlArgs...); err != nil {
		logger.Error(err.Error())
	}
	return result, nil
}

// QueryRow executes the SQL query in the Session and returns the first row.
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	logger.Info(s.sql.String(), s.sqlArgs)
	return s.db.QueryRow(s.sql.String(), s.sqlArgs...)
}

// QueryRowContext executes the SQL query in the Session with a context and returns the first row.
func (s *Session) QueryRowContext(ctx context.Context) *sql.Row {
	defer s.Clear()
	logger.Info(s.sql.String(), s.sqlArgs)
	return s.db.QueryRowContext(ctx, s.sql.String(), s.sqlArgs...)
}
