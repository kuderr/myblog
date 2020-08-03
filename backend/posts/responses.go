package posts

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Msg string `json:"msg"`
}

type ResponsePost struct {
	Msg    string `json:"msg"`
	PostId int    `json:"postId"`
}

func sendData(w http.ResponseWriter, statusCode int, resp interface{}) error {
	respData, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respData)

	return nil
}
