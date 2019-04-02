package bible

// Service map requirements
type Service struct {
	repository Repository
}

// NewService returns a service layer
func NewService(r Repository) *Service {
	return &Service{repository: r}
}

// GetVerse get one verse by book, chapter and verse
func (s Service) GetVerse(book, chapterID, verseID string) (Verse, error) {
	var verse Verse
	err := s.repository.getVerse(book, chapterID, verseID, &verse)
	return verse, err
}

// GetVerses get all verses by book and chapter
func (s Service) GetVerses(book, chapterID string) ([]Verse, error) {
	var verses []Verse
	err := s.repository.getVerses(book, chapterID, &verses)
	return verses, err
}
