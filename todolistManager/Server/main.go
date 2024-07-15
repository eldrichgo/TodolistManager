package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"server/graph"
	r "server/graph"
)

func initLogger() logger.Interface {
	logLevel := logger.Info
	f, _ := os.Create("gorm.log")
	newLogger := logger.New(
		log.New(
			io.MultiWriter(f, os.Stdout), "\r\n", log.LstdFlags), logger.Config{
			Colorful:                  true,
			LogLevel:                  logLevel,
			SlowThreshold:             time.Second,
			IgnoreRecordNotFoundError: true,
		})

	return newLogger
}

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=todolist port=5432 sslmode=prefer TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: initLogger()})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Migrate the schema
	//db.AutoMigrate(&models.Task{})

	// Initialize resolver
	resolver := &r.Resolver{Db: db}
	// Initialize schema
	schema := graph.NewExecutableSchema(graph.Config{Resolvers: resolver})
	h := handler.NewDefaultServer(schema)

	router := gin.Default()

	router.POST("/query", func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	})

	router.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL", "/query").ServeHTTP(c.Writer, c.Request)
	})

	log.Printf("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(router.Run(":8080"))
}
