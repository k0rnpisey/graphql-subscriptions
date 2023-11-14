package main

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"gqlgen-subscriptions/graph"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/go-chi/chi"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:9999/"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
			return webSocketInit(ctx, initPayload)
		},
	})
	srv.Use(extension.Introspection{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func webSocketInit(ctx context.Context, initPayload transport.InitPayload) (context.Context, *transport.InitPayload, error) {
	// Get the token from payload
	any := initPayload["authToken"]
	token, ok := any.(string)
	if !ok || token == "" {
		return nil, nil, errors.New("authToken not found in transport payload")
	}

	// Perform token verification and authentication...
	userId := "john.doe" // e.g. userId, err := GetUserFromAuthentication(token)

	// put it in context
	ctxNew := context.WithValue(ctx, "username", userId)

	return ctxNew, nil, nil
}
