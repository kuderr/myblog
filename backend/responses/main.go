package responses

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Err string `json:"error"`
}

type Response struct {
	Msg string `json:"msg"`
}

type ResponseToken struct {
	Msg   string `json:"message"`
	Token string `json:"token"`
}

type ResponsePost struct {
	Msg    string `json:"message"`
	PostId int    `json:"postId"`
}

func SendData(w http.ResponseWriter, resp interface{}, statusCode int) error {
	respData, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respData)

	return nil
}

func SendError(w http.ResponseWriter, resp interface{}, statusCode int) {
	respData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(respData)
}
