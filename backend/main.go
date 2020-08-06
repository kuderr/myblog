package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"backend/config"
	"backend/posts"
	"backend/users"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	dbUrl := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable",
		os.Getenv("BLOG_DB_USER"),
		os.Getenv("BLOG_DB_PASSWORD"),
		os.Getenv("BLOG_DB_NAME"),
	)

	config.InitDB(dbUrl)

	router := httprouter.New()
	router.POST("/posts", posts.AddPost)
	router.GET("/posts", posts.GetPosts)
	router.GET("/posts/:id", posts.GetPost)
	router.PUT("/posts/:id", posts.UpdatePost)
	router.DELETE("/posts/:id", posts.DeletePost)
	router.GET("/users/:id/posts", posts.GetUserPosts)
	router.PATCH("/posts/:id/published", posts.UpdatePostPublishedStatus)

	router.POST("/login", users.Login)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowCredentials: true,
	})

	srv := &http.Server{
		Handler:      c.Handler(router),
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server starting at http://127.0.0.1:5000")

	log.Fatal(srv.ListenAndServe())
}

func TokenAuth(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token := r.Header.Get("Token")
		if token == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), 401)
			return
		}

		handler(w, r, ps)
	}
}
