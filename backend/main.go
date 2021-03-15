package main

import (
	"backend/responses"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"backend/config"
	"backend/posts"
	"backend/users"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type Token struct {
	ID       int
	Username string
	Email    string
	Role     string
	jwt.StandardClaims
}

func main() {
	dbUrl := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable",
		os.Getenv("BLOG_DB_USER"),
		os.Getenv("BLOG_DB_PASSWORD"),
		os.Getenv("BLOG_DB_NAME"),
	)

	config.InitDB(dbUrl)

	router := httprouter.New()

	router.GET("/api/posts", posts.GetPosts)
	router.GET("/api/posts/:id", posts.GetPost)
	router.POST("/api/posts", TokenAuth(posts.AddPost))
	router.PUT("/api/posts/:id", TokenAuth(posts.UpdatePost))
	router.DELETE("/api/posts/:id", TokenAuth(posts.DeletePost))
	router.GET("/api/users/:id/posts", TokenAuth(posts.GetUserPosts))
	router.PATCH("/api/posts/:id/published", TokenAuth(posts.UpdatePostPublishedStatus))
	router.PUT("/api/posts/:id/views", posts.AddView)

	router.POST("/api/login", users.Login)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	srv := &http.Server{
		Handler:      c.Handler(router),
		Addr:         "127.0.0.1:5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server started at http://127.0.0.1:5000")

	log.Fatal(srv.ListenAndServe())
}

func TokenAuth(handler httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			resp := responses.ErrorResponse{Err: "No token provided"}
			responses.SendError(w, resp, 403)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			resp := responses.ErrorResponse{Err: "Invalid token"}
			responses.SendError(w, resp, 403)
			return
		}

		tokenPart := splitted[1]
		tk := &Token{}
		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("BLOG_TOKEN_SECRET")), nil
		})

		if err != nil {
			resp := responses.ErrorResponse{Err: "Invalid token"}
			responses.SendError(w, resp, 403)
			return
		}

		if !token.Valid {
			resp := responses.ErrorResponse{Err: "Invalid token"}
			responses.SendError(w, resp, 403)
			return
		}

		user := users.User{
			ID:       tk.ID,
			Username: tk.Username,
			Email:    tk.Email,
			Role:     tk.Role,
		}

		log.Printf("Request from %v", user.Email)

		ctx := context.WithValue(r.Context(), "user", user)
		r = r.WithContext(ctx)

		handler(w, r, ps)
	}
}
