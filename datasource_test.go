package orm

import "testing"

func TestNew(t *testing.T) {
	// Test with default options
	ds := New()
	if ds.User != "" {
		t.Errorf("Expected User to be empty, but got %s", ds.User)
	}
	if ds.Password != "" {
		t.Errorf("Expected Password to be empty, but got %s", ds.Password)
	}
	if ds.Net != "tcp" {
		t.Errorf("Expected Net to be tcp, but got %s", ds.Net)
	}
	if ds.Host != "localhost" {
		t.Errorf("Expected Host to be localhost, but got %s", ds.Host)
	}
	if ds.Port != 3306 {
		t.Errorf("Expected Port to be 3306, but got %d", ds.Port)
	}
	if ds.DBName != "" {
		t.Errorf("Expected DBName to be empty, but got %s", ds.DBName)
	}
	if ds.Driver != MySQLDriver {
		t.Errorf("Expected Driver to be MySQLDriver, but got %s", ds.Driver)
	}

	// Test with custom options
	ds = New(
		SetUser("testuser"),
		SetPassword("testpass"),
		SetNet("unix"),
		SetHost("/var/run/mysql.sock"),
		SetPort(1234),
		SetDBName("testdb"),
	)
	if ds.User != "testuser" {
		t.Errorf("Expected User to be testuser, but got %s", ds.User)
	}
	if ds.Password != "testpass" {
		t.Errorf("Expected Password to be testpass, but got %s", ds.Password)
	}
	if ds.Net != "unix" {
		t.Errorf("Expected Net to be unix, but got %s", ds.Net)
	}
	if ds.Host != "/var/run/mysql.sock" {
		t.Errorf("Expected Host to be /var/run/mysql.sock, but got %s", ds.Host)
	}
	if ds.Port != 1234 {
		t.Errorf("Expected Port to be 1234, but got %d", ds.Port)
	}
	if ds.DBName != "testdb" {
		t.Errorf("Expected DBName to be testdb, but got %s", ds.DBName)
	}
	if ds.Driver != MySQLDriver {
		t.Errorf("Expected Driver to be MySQLDriver, but got %s", ds.Driver)
	}
}

func TestDSN(t *testing.T) {
	ds := &DataSource{
		User:     "testuser",
		Password: "testpassword",
		Net:      "tcp",
		Host:     "localhost",
		Port:     3306,
		DBName:   "testdb",
		Params:   map[string]string{"charset": "utf8mb4"},
		Driver:   MySQLDriver,
	}

	expected := "testuser:testpassword@tcp(localhost:3306)/testdb?charset=utf8mb4"
	actual := ds.DSN()

	if actual != expected {
		t.Errorf("DSN() returned %s, expected %s", actual, expected)
	}
}
