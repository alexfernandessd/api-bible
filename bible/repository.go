package bible

import (
	"fmt"
)

// Repository map methods
type Repository interface {
	getVerse(book, chapterID, verseID string) (Verse, error)
	getVerses(book, chapterID string) (Verse, error)
}

// RepositoryImpl map requiments interfaces
type RepositoryImpl struct {
	db Database
}

// NewRepository create a new repository
func NewRepository(db Database) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r RepositoryImpl) getVerse(book, chapterID, verseID string) (Verse, error) {
	var verse Verse
	err := r.db.getVerse(book, chapterID, verseID, &verse)
	return verse, err
}

func (r RepositoryImpl) getVerses(book, chapterID string) (Verse, error) {
	var verse Verse
	err := r.db.getVerses(book, chapterID, &verse)
	var chapter Chapter
	fmt.Println(chapter.Verses)
	return verse, err
}
