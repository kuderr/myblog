package posts

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func AddPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	postId, err := addPost(&post)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	resp := ResponsePost{"Post successfully created", postId}
	err = sendData(w, 201, resp)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	posts, err := allPostsShorten()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}

	err = sendData(w, 200, posts)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	post, err := getPostData(postId)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = sendData(w, 200, post)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	var post Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = updatePost(postId, &post)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	resp := Response{"Post successfully updated"}
	err = sendData(w, 200, resp)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = deletePost(postId)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	resp := Response{"Post successfully deleted"}
	err = sendData(w, 200, resp)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func GetUserPosts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	posts, err := getUserPosts(userId)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = sendData(w, 200, posts)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}

func UpdatePostPublishedStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	postId, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	var post Post
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400), 400)
		return
	}

	err = updatePostPublishedStatus(postId, post.Published)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	resp := Response{"Post successfully updated"}
	err = sendData(w, 200, resp)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
	}
}
