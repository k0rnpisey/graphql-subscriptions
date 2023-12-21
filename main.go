package main

import (
	"context"
	"errors"
	gqlgenerrcode "github.com/99designs/gqlgen/graphql/errcode"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"gqlgen-subscriptions/graph"
	"gqlgen-subscriptions/graph/model"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// somewhere before creating the gqlgen 'router' object
	gqlgenerrcode.RegisterErrorType(gqlgenerrcode.ValidationFailed, gqlgenerrcode.KindUser)

	router := chi.NewRouter()
	// allow all origins
	router.Use(cors.AllowAll().Handler)

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	//router.Use(cors.New(cors.Options{
	//	AllowedOrigins:   []string{"http://localhost:5173/"},
	//	AllowCredentials: true,
	//	Debug:            true,
	//}).Handler)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		UserStore:                make(map[string]*model.User),
		PostStore:                make([]*model.Post, 0),
		NotificationStore:        make(map[string][]*model.Notification),
		NotificationSubscription: make(map[string]chan *model.Notification),
	}}))

	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	srv.Use(extension.Introspection{})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
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
