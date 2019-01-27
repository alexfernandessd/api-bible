package bible

// Service map requirements
type Service struct {
	database Database
}

// NewService returns a service layer
func NewService(db Database) *Service {
	return &Service{database: db}
}

// GetVerse get one verse by book, chapter and verse
func (s Service) GetVerse(book, chapterID, verseID string) (Verse, error) {
	var verse Verse
	err := s.database.getVerse(book, chapterID, verseID, &verse)
	return verse, err
}

// GetVerses get all verses by book and chapter
func (s Service) GetVerses(book, chapterID string) ([]Verse, error) {
	var verses []Verse
	err := s.database.getVerses(book, chapterID, &verses)
	return verses, err
}
