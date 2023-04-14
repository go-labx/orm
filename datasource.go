package orm

import (
	"fmt"
)

type DriverName string

const (
	MySQLDriver DriverName = "mysql"
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

func SetUser(user string) Option {
	return func(d *DataSource) {
		d.User = user
	}
}

func SetPassword(password string) Option {
	return func(d *DataSource) {
		d.Password = password
	}
}

func SetNet(net string) Option {
	return func(d *DataSource) {
		d.Net = net
	}
}

func SetHost(host string) Option {
	return func(d *DataSource) {
		d.Host = host
	}
}

func SetPort(port int32) Option {
	return func(d *DataSource) {
		d.Port = port
	}
}

func SetDBName(dbname string) Option {
	return func(d *DataSource) {
		d.DBName = dbname
	}
}

func SetParams(params map[string]string) Option {
	return func(d *DataSource) {
		d.Params = params
	}
}

func New(options ...Option) *DataSource {
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

func (d *DataSource) DSN() string {
	switch d.Driver {
	case MySQLDriver:
		params := mapToString(d.Params)
		return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s", d.User, d.Password, d.Net, d.Host, d.Port, d.DBName, params)
	default:
		panic("unsupported driver")
	}
}
