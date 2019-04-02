package bible

// Service map methods from service.
type Service interface {
	GetVerse(book, chapterID, verseID string) (Verse, error)
	GetVerses(book, chapterID string) ([]Verse, error)
}

// ServiceImpl map requirements.
type ServiceImpl struct {
	repository Repository
}

// NewService returns a service layer
func NewService(r Repository) *ServiceImpl {
	return &ServiceImpl{repository: r}
}

// GetVerse get one verse by book, chapter and verse
func (s ServiceImpl) GetVerse(book, chapterID, verseID string) (Verse, error) {
	var verse Verse
	err := s.repository.getVerse(book, chapterID, verseID, &verse)
	return verse, err
}

// GetVerses get all verses by book and chapter
func (s ServiceImpl) GetVerses(book, chapterID string) ([]Verse, error) {
	var verses []Verse
	err := s.repository.getVerses(book, chapterID, &verses)
	return verses, err
}
