package bible

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Database map methods from database
type Database interface {
	getVerse(bookID, chapterID, verseID string, verse *Verse) error
}

type MySQLDatabase struct {
	db *sql.DB
}

// NewConnectionMySQL create a connection with a rds
func NewConnectionMySQL(config *Config) (*MySQLDatabase, error) {
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

	for rows.Next() {
		err = rows.Scan(&verse.ID, &verse.Version, &verse.Testament, &verse.Book, &verse.Chapter, &verse.Verse, &verse.Text)
		// fmt.Printf("%s %d.%d: %s", book, verse.Chapter, verse.Verse, verse.Text)
	}

	if err != nil {
		fmt.Println("error: ", err)
	}

	return err
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
