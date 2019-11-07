package server

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jordan-wright/http-boilerplate/server/api/v1"
	"github.com/jordan-wright/unindexed"
)

// HelloWorld - образец обработчика

func HelloWorld(w http.ResponsWriter, r *http.Request) {
	fmt.Fprintf(w, "HelloWorld")
}

// NewRouter возвращает новый обработчик HTTP, который реализует маршруты основного сервера

func NewRouter() http.Handler {
	router := chi.NewRouter()

	// Настройте наше промежуточное программное обеспечение на разумные значения по умолчанию

	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.DefaultCompress)
	router.Use(middleware.Timeout(60 * time.Second))

	// Настройте наши корневые обработчики

	router.Get("/", HelloWorld)

	// Настройте наш API

	router.Mount("/api/v1/", v1.NewRouter())

	// Настройте статическую подачу файлов

	staticPath, _ := filepath.Abs("../../static/")
	fs := http.FileServer(unindexed.Dir(staticPath))
	router.Handle("/*", fs)

	return router
}
