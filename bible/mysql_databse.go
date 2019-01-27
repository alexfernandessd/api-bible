package bible

import (
	"database/sql"
	"fmt"
	"log"
)

// Database map methods from database
type Database interface {
	getVerse(bookID, chapterID, verseID string, verse *Verse) error
	getVerses(bookID, chapterID string, verses *[]Verse) error
}

// MySQLDatabase contains a conection with SQL
type MySQLDatabase struct {
	db *sql.DB
}

// NewConnectionMySQL create a connection with a rds
func NewConnectionMySQL(config *Config) (*MySQLDatabase, error) {
	// Connection with AWS
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=false",
		config.AWSUser, config.AWSPassword, config.MySqlbEndpoint, config.AWSInstance,
	)
	dbConn, err := sql.Open("mysql", dnsStr)

	connErr := dbConn.Ping()
	if connErr != nil {
		log.Fatal("failed to connect on database: ", connErr)
	}

	return &MySQLDatabase{db: dbConn}, err
}

func (m MySQLDatabase) getVerse(book, chapterID, verseID string, verse *Verse) error {
	bookID, err := m.getBookByID(book)

	rows, err := m.db.Query("SELECT * FROM verses where book = ? and chapter = ? and verse = ?", bookID, chapterID, verseID)
	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	for rows.Next() {
		err = rows.Scan(&verse.ID, &verse.Version, &verse.Testament, &verse.Book, &verse.Chapter, &verse.Verse, &verse.Text)
	}
	if err != nil {
		fmt.Println("error: ", err)
	}

	return err
}

func (m MySQLDatabase) getVerses(book, chapterID string, verses *[]Verse) error {
	bookID, err := m.getBookByID(book)

	rows, err := m.db.Query("SELECT * FROM verses where book = ? and chapter = ?", bookID, chapterID)
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

func (m MySQLDatabase) getBookByID(book string) (int, error) {
	rows, err := m.db.Query("SELECT id FROM books where name = ?", book)
	var bookID int

	for rows.Next() {
		var id int
		err = rows.Scan(&id)

		if err == nil {
			bookID = id - 1
		}
	}
	return bookID, err
}
