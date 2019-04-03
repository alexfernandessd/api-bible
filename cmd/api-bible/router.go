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

func createServerHandler(service bible.Service) http.Handler {
	router := chi.NewRouter()

	router.Get("/version", versionHandler())

	router.Get("/testaments/{testament}/books", getBooksByTestament(service))

	router.Route("/books/{book}", func(router chi.Router) {
		router.Get("/verses", getVersesByBook(service))
		router.Route("/chapters", func(router chi.Router) {
			router.Route("/{chapterID}", func(router chi.Router) {
				router.Get("/verses", getVersesByChapter(service))
				router.Get("/verses/{verseID}", getVerse(service))
			})
		})
	})

	return router
}
