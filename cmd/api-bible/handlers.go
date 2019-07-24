package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexfernandessd/api-bible/bible"
	"github.com/go-chi/chi"
)

func versionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: must be implemented...
		// json.NewEncoder(w).Encode()
	}
}

func getBooksByTestament(svc bible.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: must be implemented...
		// json.NewEncoder(w).Encode()
	}
}

func getVersesByBook(svc bible.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO: must be implemented...
		// json.NewEncoder(w).Encode()
	}
}

func getRandomVerse(svc bible.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, t *http.Request) {
		verse, _ := svc.GetRandomVerse()
		json.NewEncoder(w).Encode(verse)
	}
}

func getVersesByChapter(svc bible.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := chi.URLParam(r, urlParamBook)
		if book == "" {
			fmt.Println("erro", book)
		}

		chapterID := chi.URLParam(r, urlParamChapterID)
		if chapterID == "" {
			fmt.Println("erro", chapterID)
		}

		verses, _ := svc.GetVerses(book, chapterID)

		json.NewEncoder(w).Encode(verses)
	}
}

func getVerse(svc bible.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		book := chi.URLParam(r, urlParamBook)
		if book == "" {
			fmt.Println("book param can not be empty: ", book)
			//TODO: error wirter
		}

		chapterID := chi.URLParam(r, urlParamChapterID)
		if chapterID == "" {
			fmt.Println("chapter param can not be empty: ", chapterID)
			//TODO: error wirter
		}

		verseID := chi.URLParam(r, urlParamVerseID)
		if verseID == "" {
			fmt.Println("verse param can not be empty: ", verseID)
			//TODO: error wirter
		}

		verse, err := svc.GetVerse(book, chapterID, verseID)
		if err != nil {
			fmt.Println("error on get verse: ", err)
			//TODO: error wirter
		}

		json.NewEncoder(w).Encode(verse)
	}
}
