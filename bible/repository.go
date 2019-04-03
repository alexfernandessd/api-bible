package bible

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	dbDriver string = "mysql"
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

// NewConnectionMySQL create a connection with a MySQL.
func NewConnectionMySQL(dbUser, dbPassword, dbEndPoint, dbInstance, connectionString string) (*sql.DB, error) {
	dnsStr := fmt.Sprintf(connectionString,
		dbUser, dbPassword, dbEndPoint, dbInstance,
	)

	dbConn, err := sql.Open(dbDriver, dnsStr)

	connErr := dbConn.Ping()
	if connErr != nil {
		log.Fatal("failed to connect on database: ", connErr)
		return nil, connErr
	}

	return dbConn, err
}

func (r RepositoryImpl) getVerse(book, chapterID, verseID string, verse *Verse) error {
	start := time.Now()

	bookID, err := r.getBookByID(book)
	if err != nil {
		return err
	}

	row := r.db.QueryRow("SELECT version, text FROM verses where book = ? and chapter = ? and verse = ? LIMIT 1", *bookID, chapterID, verseID)

	fmt.Println("execution time of 2 querys: ", time.Since(start))

	err = row.Scan(&verse.Version, &verse.Text)
	if err != nil {
		fmt.Println("error on scan row: ", err)
		return err
	}

	return nil
}

func (r RepositoryImpl) getVerses(book, chapterID string, verses *[]Verse) error {
	bookID, err := r.getBookByID(book)

	rows, err := r.db.Query("SELECT * FROM verses where book = ? and chapter = ?", *bookID, chapterID)
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

func (r RepositoryImpl) getBookByID(book string) (*int, error) {
	row := r.db.QueryRow("SELECT id FROM books WHERE name = ? LIMIT 1", book)

	var bookID int
	err := row.Scan(&bookID)

	if err != nil && bookID == 0 {
		// TODO: Personal error
		return nil, errors.New("book not found")
	}

	if err != nil {
		return nil, err
	}

	return &bookID, nil
}
