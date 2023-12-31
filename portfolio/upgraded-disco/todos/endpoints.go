package todos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/common"
	"server/config"
	"server/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	logger "github.com/rs/zerolog/log"
)

func handleRequests() {
	router := routes()

	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{config.FetchConfig().SERVERORIGIN})
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"})
	methods := handlers.AllowedMethods([]string{"POST"})
	// ttl := handlers.MaxAge(3600)

	logger.Info().Msg(fmt.Sprintf("Starting server on port: %v", config.FetchConfig().Ports.Todos))
	logger.Info().Err(http.ListenAndServe(fmt.Sprintf(":%v", config.FetchConfig().Ports.Todos), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func routes() *mux.Router {
	routeHandler := middleware.RouteHandler{}
	r := mux.NewRouter()
	r.Use(routeHandler.InternalOriginMiddleware)
	r.HandleFunc("/list", returnTodos).Methods("POST")
	r.HandleFunc("/new", reqHandler).Methods("POST")  // add todo to mongo, add new issue to respective repo
	r.HandleFunc("/done", reqHandler).Methods("POST") // delete todo from mongo, change status on github
	r.NotFoundHandler = http.HandlerFunc(common.InvalidEndpoint)
	return r
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	var body RequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch body.Action {
	case "list":
		returnTodos(w, r)
	case "add":
		// AddTodo(w, r)
	case "done":
		// MarkDone(w, r)
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
	}
}

func StartServer() {
	handleRequests()
}
