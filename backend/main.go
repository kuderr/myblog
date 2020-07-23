package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"backend/models"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type Response struct {
	Msg string `json:"msg"`
}

type Env struct {
	db models.Datastore
}

func main() {
	dbUrl := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable",
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
	router.POST("/posts", env.addPost)
	router.GET("/posts/:id", env.getPost)

	handler := cors.Default().Handler(router)

	srv := &http.Server{
		Handler:      handler,
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

func (env *Env) getPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := env.db.AllPostsShorten()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	respData, err_ := json.Marshal(posts)
	if err_ != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(respData)
}

func (env *Env) getPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	post, err := env.db.GetPostData(postId)
	switch {
	case err == sql.ErrNoRows:
		log.Println(err)
		http.NotFound(w, r)
		return
	case err != nil:
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	respData, err_ := json.Marshal(post)
	if err_ != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(respData)
}

func (env *Env) addPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = env.db.AddPost(&post)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	resp := Response{"Post successfully created"}
	respData, err_ := json.Marshal(resp)
	if err_ != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(respData)
}
