package db

import (
	"strings"

	"github.com/go-sql-driver/mysql"
)

// Dialect names for external usage.
const (
	MySQL    = "mysql"
	SQLite   = "sqlite3"
	Postgres = "postgres"
	Gremlin  = "gremlin"
)

// Dsn It is the name that applications use to request a connection to an ODBC Data Source. In other words, it is a symbolic name that represents the ODBC connection.
type Dsn interface {
	Source() string
	Dialect() string
}

type Source struct {
	source  string
	dialect string
}

func (s Source) Source() string {
	return s.source
}

func (s Source) Dialect() string {
	return s.dialect
}

func NewSource(c *Config) Dsn {
	source := Source{
		dialect: MySQL,
		source:  "/",
	}
	if c.GetDriver() != "" {
		source.dialect = c.GetDriver()
	}
	if c.GetSource() != "" {
		source.source = c.GetSource()
	} else {
		i := strings.Builder{}
		switch strings.ToLower(c.GetDriver()) {
		case Postgres:
			if c.GetHostname() != "" {
				i.Write([]byte("host="))
				i.WriteString(c.GetHostname())
				i.Write([]byte(" "))
			}
			if c.GetHostport() > 0 {
				i.Write([]byte("port="))
				i.WriteRune(rune(c.GetHostport()))
				i.Write([]byte(" "))
			}
			if c.GetUsername() != "" {
				i.Write([]byte("user="))
				i.WriteString(c.GetUsername())
				i.Write([]byte(" "))
			}
			if c.GetPassword() != "" {
				i.Write([]byte("password="))
				i.WriteString(c.GetPassword())
				i.Write([]byte(" "))
			}
			if c.GetDatabase() != "" {
				i.Write([]byte("dbname="))
				i.WriteString(c.GetDatabase())
				i.Write([]byte(" "))
			}
			for k, v := range c.GetParams() {
				v := strings.TrimSpace(v)
				i.WriteString(k)
				i.Write([]byte("="))
				i.WriteString(v)
				i.Write([]byte(" "))
			}
			source.source = i.String()
		case MySQL, SQLite, Gremlin:
			fallthrough
		default:
			c := mysql.Config{
				User:    c.GetUsername(),
				Passwd:  c.GetPassword(),
				Addr:    c.GetHostname(),
				DBName:  c.GetDatabase(),
				Timeout: c.GetTimeout().AsDuration(),
				Params:  c.GetParams(),
			}
			source.source = c.FormatDSN()
		}
	}

	return &source
}
