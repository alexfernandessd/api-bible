package bible

// Verse contains elements from a verse
type Verse struct {
	ID        int
	Version   string
	Testament int
	Book      int
	Chapter   int
	Verse     int
	Text      string
}

type Chapter struct {
	Verses []Verse
}
