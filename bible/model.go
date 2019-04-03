package bible

// Verse contains elements from a verse.
type Verse struct {
	ID        int    `json:"verseID,omitempty"`
	Version   string `json:"version,omitempty"`
	Testament int    `json:"testament,omitempty"`
	Book      int    `json:"book,omitempty"`
	Chapter   int    `json:"chapter,omitempty"`
	Verse     int    `json:"verse,omitempty"`
	Text      string `json:"text"`
}
