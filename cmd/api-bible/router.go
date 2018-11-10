package main

import (
	"net/http"

	"github.com/alexfernandessd/api-bible/bible"
	"github.com/go-chi/chi"
)

const (
	urlParamTestamentID = "testamentID"
	urlParamBook        = "book"
	urlParamChapterID   = "chapterID"
	urlParamVerseID     = "verseID"
)

func createServerHandler(service *bible.Service) http.Handler {
	router := chi.NewRouter()

	// router.Get("/version", versionHandler)

	router.Route("/testament", func(router chi.Router) {
		router.Get("/{testamentID}", getBooksByTestament(service))
	})

	router.Route("/book", func(router chi.Router) {
		router.Get("/{book}", getChapterByBook(service))
		router.Get("/{book}/chapter/{chapterID}", getVersesByChapter(service))
		router.Get("/{book}/chapter/{chapterID}/verse/{verseID}", getVerse(service))
	})

	return router
}
