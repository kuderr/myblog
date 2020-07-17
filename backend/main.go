package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"backend/models"

	"github.com/julienschmidt/httprouter"
)

type Env struct {
	db models.Datastore
}

func main() {
	dbUrl := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("BLOG_DB_USER"),
		os.Getenv("BLOG_DB_PASSWORD"),
		os.Getenv("BLOG_DB_NAME"),
	)

	db, err := models.InitDB(dbUrl)
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	router := httprouter.New()
	router.GET("/posts", env.getPosts)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

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

func (env *Env) getPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := env.db.AllPostsShorten()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	for _, post := range posts {
		fmt.Fprintf(w, "%d, %s, %s\n", post.ID, post.Title, post.Summary)
	}
}
