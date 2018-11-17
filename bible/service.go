package bible

// Service map requirements
type Service struct {
	repository Repository
}

// NewService returns a service layer
func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

// GetVerse get one verse by book, chapter and verse
func (s Service) GetVerse(book, chapterID, verseID string) (Verse, error) {
	verse, err := s.repository.getVerse(book, chapterID, verseID)
	return verse, err
}
