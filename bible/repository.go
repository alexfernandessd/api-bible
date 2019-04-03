package bible

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

// Repository map methods from repository.
type Repository interface {
	getVerse(bookID, chapterID, verseID string, verse *Verse) error
	getVerses(bookID, chapterID string, verses *[]Verse) error
}

// RepositoryImpl respository implementation.
type RepositoryImpl struct {
	db *sql.DB
}

// NewRepository is the repository constructor.
func NewRepository(db *sql.DB) *RepositoryImpl {
	return &RepositoryImpl{
		db: db,
	}
}

// NewConnectionMySQL create a connection with a rds.
func NewConnectionMySQL(dbUser, dbPassword, dbEndPoint, dbInstance, connectionString string) (*sql.DB, error) {
	// Connection with AWS
	dnsStr := fmt.Sprintf(connectionString,
		dbUser, dbPassword, dbEndPoint, dbInstance,
	)

	dbConn, err := sql.Open("mysql", dnsStr)

	connErr := dbConn.Ping()
	if connErr != nil {
		log.Fatal("failed to connect on database: ", connErr)
	}

	return dbConn, err
}

func (r RepositoryImpl) getVerse(book, chapterID, verseID string, verse *Verse) error {
	bookID, err := r.getBookByID(book)
	if err != nil {
		return err
	}

	rows, err := r.db.Query("SELECT version, text FROM verses where book = ? and chapter = ? and verse = ?", bookID, chapterID, verseID)
	if err != nil {
		fmt.Println("error on execute query: ", err)
		return err
	}

	if !rows.Next() {
		return errors.New("chapter or verse not found")
	}

	for rows.Next() {
		err = rows.Scan(&verse.Version, &verse.Text)
	}
	if err != nil {
		fmt.Println("error on scan rows: ", err)
		return err
	}

	return nil
}

func (r RepositoryImpl) getVerses(book, chapterID string, verses *[]Verse) error {
	bookID, err := r.getBookByID(book)

	rows, err := r.db.Query("SELECT * FROM verses where book = ? and chapter = ?", bookID, chapterID)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	var verse Verse

	for rows.Next() {
		err = rows.Scan(
			&verse.ID,
			&verse.Version,
			&verse.Testament,
			&verse.Book,
			&verse.Chapter,
			&verse.Verse,
			&verse.Text,
		)
		*verses = append(*verses, verse)
	}

	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	return nil
}

func (r RepositoryImpl) getBookByID(book string) (int, error) {
	rows, err := r.db.Query("SELECT id FROM books where name = ?", book)
	var bookID int

	if !rows.Next() {
		return bookID, errors.New("book not found")
	}

	for rows.Next() {
		var id int
		err = rows.Scan(&id)

		if err == nil {
			bookID = id - 1
		}
	}
	return bookID, err
}
