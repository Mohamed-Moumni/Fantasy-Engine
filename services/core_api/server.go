package main

import (
	"core_api/graph"
	"core_api/graph/resolvers"
	"core_api/internal/auth"
	"core_api/internal/config"
	"log"
	"net/http"
	"os"
	"pkg/database"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/vektah/gqlparser/v2/ast"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	godotenv.Load(".core_api.env")

	if port == "" {
		port = defaultPort
	}

	auth.InitGoogleAuth()
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	
	db := database.Init()
	
	// database.AutoMigrate(db)

	// os.Exit(1)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &resolvers.Resolver{
		DB: db,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	mux := http.NewServeMux()
	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/auth/google/callback", http.HandlerFunc(auth.GoogleCallback))
	mux.Handle("/query", auth.AuthMiddleware()(srv))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},  // Allow all origins
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},  // Allow all headers
		AllowCredentials: false,          // Must be false when using "*"
	})

	handler := c.Handler(mux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
