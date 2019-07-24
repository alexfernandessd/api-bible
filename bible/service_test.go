package bible

import (
	"testing"
)

func TestService_GetVerse(t *testing.T) {
	type stubs struct {
		getVerse func(bookID, chapterID, verseID string, verse *Verse) error
	}
	type args struct {
		book      string
		chapterID string
		verseID   string
	}
	tests := []struct {
		name    string
		stubs   stubs
		args    args
		wantErr bool
	}{
		{
			"get verse with success",
			stubs{
				getVerse: func(bookID, chapterID, verseID string, verse *Verse) error {
					return nil
				},
			},
			args{
				book:      "genesis",
				chapterID: "1",
				verseID:   "1",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := repositoryMock{
				getVerseStub: tt.stubs.getVerse,
			}

			service := NewService(&repo)

			_, err := service.GetVerse(tt.args.book, tt.args.chapterID, tt.args.verseID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetVerse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

type repositoryMock struct {
	getVerseStub       func(bookID, chapterID, verseID string, verse *Verse) error
	getVersesStub      func(bookID, chapterID string, verses *[]Verse) error
	getRandomVerseStub func(verse *Verse) error
}

func (r *repositoryMock) getRandomVerse(verse *Verse) error {
	return r.getRandomVerseStub(verse)
}

func (r *repositoryMock) getVerse(bookID, chapterID, verseID string, verse *Verse) error {
	return r.getVerseStub(bookID, chapterID, verseID, verse)
}

func (r *repositoryMock) getVerses(bookID, chapterID string, verses *[]Verse) error {
	return r.getVersesStub(bookID, chapterID, verses)
}
