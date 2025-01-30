package infrastructure

import (
	"fmt"
	"log"
	"project1/src/books/domain"
	"project1/src/books/domain/entities"
	"project1/src/core"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() domain.IBook {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(book entities.Book) (entities.Book, error) {
	query := "INSERT INTO books (title, price) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, book.Title, book.Price)
	if err != nil {
		log.Printf("Error al insertar libro: %v", err)
		return entities.Book{}, err
	}

	lastInsertID, _ := result.LastInsertId()
	book.ID = int32(lastInsertID)

	return book, nil
}

func (mysql *MySQL) FindAll() ([]entities.Book, error) {
	query := "SELECT id, title, price FROM books"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		return nil, err
	}
	defer rows.Close()

	var books []entities.Book
	for rows.Next() {
		var book entities.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Price); err != nil {
			log.Printf("Error al escanear libro: %v", err)
			return nil, err
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (mysql *MySQL) FindByID(id int) (entities.Book, error) {
	query := "SELECT id, title, price FROM books WHERE id = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		log.Printf("Error al ejecutar la consulta para ID %d: %v", id, err)
		return entities.Book{}, err
	}
	defer rows.Close()

	var book entities.Book
	if rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Price)
		if err != nil {
			return entities.Book{}, fmt.Errorf("error al escanear libro: %v", err)
		}
		return book, nil
	}

	return entities.Book{}, fmt.Errorf("libro no encontrado")
}

func (mysql *MySQL) Update(id int, book entities.Book) error {
	query := "UPDATE books SET title = ?, price = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, book.Title, book.Price, id)
	if err != nil {
		log.Printf("Error al modificar libro con ID %d: %v", id, err)
		return fmt.Errorf("error al modificar libro: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		log.Printf("No se encontró libro con ID %d para modificar", id)
		return fmt.Errorf("no se encontró el libro con ID %d", id)
	}

	log.Printf("Libro con ID %d modificado correctamente", id)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Printf("Error al eliminar libro con ID %d: %v", id, err)
		return fmt.Errorf("error al eliminar libro: %v", err)
	}
	return nil
}
