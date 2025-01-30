package infrastructure

import (
	"fmt"
	"log"
	"project1/src/core"
	"project1/src/reader/domain"
	"project1/src/reader/domain/entities"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() domain.IReader {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Safe(reader entities.Reader) (entities.Reader, error) {
	query := "INSERT INTO readers (name, age) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, reader.Name, reader.Age)
	if err != nil {
		log.Printf("Error al insertar lector: %v", err)
		return entities.Reader{}, err
	}

	lastInsertID, _ := result.LastInsertId()
	reader.ID = int32(lastInsertID)

	return reader, nil
}

func (mysql *MySQL) LocateAll() ([]entities.Reader, error) {
	query := "SELECT id, name, age FROM readers"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return nil, err
	}
	defer rows.Close()

	var readers []entities.Reader
	for rows.Next() {
		var reader entities.Reader
		if err := rows.Scan(&reader.ID, &reader.Name, &reader.Age); err != nil {
			log.Printf("Error al escanear lector: %v", err)
			return nil, err
		}
		readers = append(readers, reader)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return readers, nil
}

func (mysql *MySQL) LocateByID(id int) (entities.Reader, error) {
	query := "SELECT id, name, age FROM readers WHERE id = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		log.Printf("Error al ejecutar la consulta para ID %d: %v", id, err)
		return entities.Reader{}, err
	}
	defer rows.Close()

	var reader entities.Reader
	if rows.Next() {
		err := rows.Scan(&reader.ID, &reader.Name, &reader.Age)
		if err != nil {
			return entities.Reader{}, fmt.Errorf("error al escanear lector: %v", err)
		}
		return reader, nil
	}

	return entities.Reader{}, fmt.Errorf("lector no encontrado")
}

func (mysql *MySQL) Upgrade(id int, reader entities.Reader) error {
	query := "UPDATE readers SET name = ?, age = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, reader.Name, reader.Age, id)
	if err != nil {
		log.Printf("Error al modificar lector con ID %d: %v", id, err)
		return fmt.Errorf("error al modificar lector: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("No se encontró lector con ID %d para modificar", id)
		return fmt.Errorf("no se encontró el lector con ID %d", id)
	}

	log.Printf("Lector con ID %d modificado correctamente", id)
	return nil
}

func (mysql *MySQL) Erase(id int) error {
	query := "DELETE FROM readers WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("Error al eliminar lector con ID %d: %v", id, err)
		return fmt.Errorf("error al eliminar lector: %v", err)
	}
	return nil
}
