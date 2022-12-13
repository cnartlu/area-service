package database

import (
	"bytes"
	"errors"
	"net"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func ensureHavePort(addr string, defaultPort int) string {
	if _, _, err := net.SplitHostPort(addr); err != nil {
		return net.JoinHostPort(addr, strconv.Itoa(defaultPort))
	}
	return addr
}

func ParseDSN(c *Config) string {
	var buf bytes.Buffer
	hasParam := false

	switch strings.ToLower(c.GetDriver()) {
	case "postgres", "psg":
		// Set default network if empty
		if c.GetNet() == "" {
			c.Net = "tcp"
		}

		// Set default address if empty
		if c.GetAddress() == "" {
			switch c.GetNet() {
			case "tcp":
				c.Address = "127.0.0.1:3306"
			case "unix":
				c.Address = "/tmp/mysql.sock"
			default:
				panic(errors.New("default addr for network '" + c.GetNet() + "' unknown"))
			}
		} else if c.GetNet() == "tcp" {
			c.Address = ensureHavePort(c.GetAddress(), 3306)
		}

		// if c.GetHostname() != "" {
		// 	buf.Write([]byte("host="))
		// 	buf.WriteString(c.GetHostname())
		// 	buf.Write([]byte(" "))
		// }
		// if c.GetHostport() > 0 {
		// 	buf.Write([]byte("port="))
		// 	buf.WriteRune(rune(c.GetHostport()))
		// 	buf.Write([]byte(" "))
		// }
		if c.GetUsername() != "" {
			buf.Write([]byte("user="))
			buf.WriteString(c.GetUsername())
			buf.Write([]byte(" "))
		}
		if c.GetPassword() != "" {
			buf.Write([]byte("password="))
			buf.WriteString(c.GetPassword())
			buf.Write([]byte(" "))
		}
		if c.GetDatabase() != "" {
			buf.Write([]byte("dbname="))
			buf.WriteString(c.GetDatabase())
			buf.Write([]byte(" "))
		}

		if len(c.GetParams()) > 0 {
			for k, v := range c.GetParams() {
				v := strings.TrimSpace(v)
				buf.WriteString(k)
				buf.Write([]byte("="))
				buf.WriteString(v)
				buf.Write([]byte(" "))
			}
		}
	default:
		// Set default network if empty
		if c.GetNet() == "" {
			c.Net = "tcp"
		}

		// Set default address if empty
		if c.GetAddress() == "" {
			switch c.GetNet() {
			case "tcp":
				c.Address = "127.0.0.1:3306"
			case "unix":
				c.Address = "/tmp/mysql.sock"
			default:
				panic(errors.New("default addr for network '" + c.GetNet() + "' unknown"))
			}
		} else if c.GetNet() == "tcp" {
			c.Address = ensureHavePort(c.GetAddress(), 3306)
		}

		// [username[:password]@]
		if len(c.GetUsername()) > 0 {
			buf.WriteString(c.GetUsername())
			if len(c.GetPassword()) > 0 {
				buf.WriteByte(':')
				buf.WriteString(c.GetPassword())
			}
			buf.WriteByte('@')
		}

		// [protocol[(address)]]
		if len(c.GetNet()) > 0 {
			buf.WriteString(c.GetNet())
			if len(c.GetAddress()) > 0 {
				buf.WriteByte('(')
				buf.WriteString(c.GetAddress())
				buf.WriteByte(')')
			}
		}

		// /dbname
		buf.WriteByte('/')
		buf.WriteString(c.GetDatabase())

		// [?param1=value1&...&paramN=valueN]
		if len(c.GetParams()) > 0 {
			var params []string
			for param := range c.GetParams() {
				params = append(params, param)
			}
			sort.Strings(params)
			for _, param := range params {
				var value = url.QueryEscape(c.Params[param])
				buf.Grow(1 + len(param) + 1 + len(value))
				if !hasParam {
					hasParam = true
					buf.WriteByte('?')
				} else {
					buf.WriteByte('&')
				}
				buf.WriteString(param)
				buf.WriteByte('=')
				buf.WriteString(value)
			}
		}
	}

	return buf.String()

}
