package users

import (
	"backend/responses"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		resp := responses.ErrorResponse{Err: "Invalid post data"}
		responses.SendError(w, resp, 400)
		return
	}

	compareUser, err := login(&user)
	if err != nil {
		log.Println(err)
		resp := responses.ErrorResponse{Err: "User not found"}
		responses.SendError(w, resp, 404)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(compareUser.Password), []byte(user.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		log.Println(err)
		resp := responses.ErrorResponse{Err: "Incorrect password"}
		responses.SendError(w, resp, 403)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       compareUser.ID,
		"username": compareUser.Username,
		"email":    compareUser.Email,
		"fullName": compareUser.FullName,
		"avatar":   compareUser.Avatar,
		"role":     compareUser.Role,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("BLOG_TOKEN_SECRET")))

	resp := responses.ResponseToken{Msg: "Successfully logged in", Token: tokenString}
	err = responses.SendData(w, resp, 200)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		resp := responses.ErrorResponse{Err: "Invalid post data"}
		responses.SendError(w, resp, 400)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = create(user.Email, string(hashedPassword))
	if err != nil {
		log.Println(err)
		resp := responses.ErrorResponse{Err: "User with provided email already exist"}
		responses.SendError(w, resp, 409)
		return
	}

	resp := responses.Response{Msg: "Successfully created"}
	responses.SendData(w, resp, 201)
}
