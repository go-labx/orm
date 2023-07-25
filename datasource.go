package orm

import (
	"fmt"
)

type DriverName string

const (
	MySQLDriver DriverName = MySQL
)

type DataSource struct {
	User     string
	Password string
	Net      string
	Host     string
	Port     int32
	DBName   string
	Params   map[string]string

	Driver DriverName
}

type Option func(*DataSource)

// SetUser sets the user for the DataSource
func SetUser(user string) Option {
	return func(d *DataSource) {
		d.User = user
	}
}

// SetPassword sets the password for the DataSource
func SetPassword(password string) Option {
	return func(d *DataSource) {
		d.Password = password
	}
}

// SetNet sets the network for the DataSource
func SetNet(net string) Option {
	return func(d *DataSource) {
		d.Net = net
	}
}

// SetHost sets the host for the DataSource
func SetHost(host string) Option {
	return func(d *DataSource) {
		d.Host = host
	}
}

// SetPort sets the port for the DataSource
func SetPort(port int32) Option {
	return func(d *DataSource) {
		d.Port = port
	}
}

// SetDBName sets the database name for the DataSource
func SetDBName(dbname string) Option {
	return func(d *DataSource) {
		d.DBName = dbname
	}
}

// SetParams sets the parameters for the DataSource
func SetParams(params map[string]string) Option {
	return func(d *DataSource) {
		d.Params = params
	}
}

// NewDataSource creates a new DataSource with the given options
func NewDataSource(options ...Option) *DataSource {
	datasource := &DataSource{
		User:     "",
		Password: "",
		Net:      "tcp",
		Host:     "localhost",
		Port:     3306,
		DBName:   "",
		Driver:   MySQLDriver,
	}

	for _, option := range options {
		option(datasource)
	}

	return datasource
}

// DSN returns the data source name for the DataSource
func (d *DataSource) DSN() string {
	switch d.Driver {
	case MySQLDriver:
		params := MapToString(d.Params)
		return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s", d.User, d.Password, d.Net, d.Host, d.Port, d.DBName, params)
	default:
		panic("unsupported driver")
	}
}
