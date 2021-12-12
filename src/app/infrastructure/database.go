package infrastructure

import (
	"log"
	"fmt"
	"os"
	"database/sql"
	_  "github.com/go-sql-driver/mysql"
	"github.com/taise-hub/shellgame/src/app/interfaces/database"
)

var (
	conf  = &Config{
		User:     getenv("MYSQL_USER", "shellgame"),
		Password: getenv("MYSQL_PASSWORD", "password"),
		Server:   getenv("MYSQL_HOST", "127.0.0.1"),
		Port:     3306,
		DBName:   getenv("MYSQL_DATABASE", "shellgame"),
	}
)

type Config struct {
	User     string
	Password string
	Server   string
	Port     int
	DBName   string
}

type SqlHandler struct {
	Conn *sql.DB
}

func getenv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	} else {
		return val
	}
}

func (c *Config) setUpDest() string {
	protocol := fmt.Sprintf("tcp(%s:%d)", c.Server, c.Port)
	return fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True", c.User, c.Password, protocol, c.DBName)
}

func NewSqlHandler() *SqlHandler {
	conn, err := sql.Open("mysql", conf.setUpDest())
	if err != nil {
		panic(err)
	}
	handler := new(SqlHandler)
	handler.Conn = conn
	return handler
}

func (h *SqlHandler) Query(query string, args ...interface{}) (database.Rows, error) {
	rows, err := h.Conn.Query(query, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return rows, nil
}

func (h *SqlHandler) Exec(query string, args ...interface{}) (database.Result, error) {
	result, err := h.Conn.Exec(query, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}