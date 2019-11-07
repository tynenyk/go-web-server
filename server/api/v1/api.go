package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"string"

	"log"

	"github.com/go-chi/chi"
)

// ValidBearer - жестко закодированный токен на предъявителя для демонстрационных целей

const ValidBearer = "123456"

// HelloResponse - это представление JSON для настроенного сообщения

type HelloResponse struct {
	Message string `json: "message"`
}

func jsonResponse(w http.ResponseWriter, data interface{}, c int) {
	dj, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)

}

// HelloWorld возвращает простой «Hello World!» сообщение

func HelloWorld (w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Hello World!",
	}
	jsonResponse(w, response, http.StatusOK)
}

// HelloName возвращает персонализированное сообщение JSON

func HelloName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	response := HelloResponse{
		Message: fmt.Sprintf("Hello %s!", name),
	}
	jsonResponse(w, response, http.StatusOK)
}

// RequireAuthentication - это пример обработчика промежуточного программного обеспечения, который проверяет наличие
// жестко закодированный токен на предъявителя. Это может быть использовано для проверки сеансовых файлов cookie, JWT
// и более.

func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFanc(func(w http.ResponseWriter, r *http.Request){
		
		// Убедитесь, что указан заголовок авторизации

		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		// 	Предполагая, что всё прошло гладко, мы можем выполнить аутентифицированный обработчик

		next.ServeHTTP(w, r)
	})
}

// NewRouter возвращает обработчик HTTP, который реализует маршруты для API

func NewRouter() http.Handler {
	r := chi.NewRouter()

	r.Use(RequireAuthentication)

	// Зарегистрировать маршруты API

	r.Get("/", HelloWorld)
	r.Get("/{name", HelloName)

	return r
}
}