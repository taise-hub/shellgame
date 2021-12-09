package database

type SqlHandler interface {
	Exec(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Rows, error)
}

type Rows interface {
	Close() error
	Scan(dest ...interface{}) error
	Next() bool
}

type Result interface {
	LastInsertId() (int64, error)
}