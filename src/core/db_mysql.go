package core

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Conn_MySQL struct {
	Db  *sql.DB
	Err string
}

var (
	instance *Conn_MySQL
	once     sync.Once
)

// Conexión a la base de datos MySQL
func GetDBPool() *Conn_MySQL {
	once.Do(func() {
		// Cambia "books_db" por el nombre correcto de tu base de datos
		dsn := "root:noobmaster69@tcp(localhost:3306)/project1"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error al conectar con MySQL: %v", err)
		}

		if err = db.Ping(); err != nil {
			log.Fatalf("Error al hacer ping a MySQL: %v", err)
		}

		instance = &Conn_MySQL{
			Db:  db,
			Err: "",
		}
	})

	return instance
}

// Ejecuta una consulta preparada con argumentos
func (c *Conn_MySQL) ExecutePreparedQuery(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := c.Db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}

// Obtiene filas de una consulta y devuelve el error si ocurre
func (c *Conn_MySQL) FetchRows(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := c.Db.Query(query, args...)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return nil, err
	}
	return rows, nil
}
