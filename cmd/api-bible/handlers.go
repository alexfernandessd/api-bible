package main

import (
	"encoding/json"
	"net/http"

	"github.com/alexfernandessd/api-bible/bible"
	"github.com/go-chi/chi"
)

// getBooksByTestament
func getBooksByTestament(svc *bible.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// json.NewEncoder(w).Encode()
	}
}

// getChapterByBook
func getChapterByBook(svc *bible.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// json.NewEncoder(w).Encode()
	}
}

// getVersesByChapte
func getVersesByChapter(svc *bible.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// json.NewEncoder(w).Encode()
	}
}

// getVerse
func getVerse(svc *bible.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := chi.URLParam(r, urlParamBook)
		chapterID := chi.URLParam(r, urlParamChapterID)
		verseID := chi.URLParam(r, urlParamVerseID)

		verse, _ := svc.GetVerse(book, chapterID, verseID)

		json.NewEncoder(w).Encode(verse)
	}
}
