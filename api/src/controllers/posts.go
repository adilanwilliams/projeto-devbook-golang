package controllers

import (
	"devbook/src/models"
	"devbook/src/services"
	"devbook/src/utils/response"
	"encoding/json"
	"io"
	"net/http"
)

// SavePost handles the creation of a new post.
func SavePost(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(body, &post); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	userToken := r.Context().Value("userID").(uint64)

	service, err := services.NewPostService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	post.ID, err = service.SavePost(post, userToken)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusCreated, response.Response{
		Success: true,
		Data: post,
	})

}
