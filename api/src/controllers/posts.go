package controllers

import (
	"devbook/src/authentication"
	"devbook/src/models"
	"devbook/src/services"
	"devbook/src/utils/response"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	post.AuthorID = userToken
	post.ID, err = service.SavePost(post)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusCreated, response.Response{
		Success: true,
		Data:    post,
	})

}

// FindPostByID retrieves a post by its unique identifier.
func FindPostByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	service, err := services.NewPostService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	post, err := service.FindPostByID(postID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data: post,
	})
}

// FindUserFeed handles the request to retrieve the authenticated user's feed.
// The feed includes posts created by the user and by the users they follow.
func FindUserFeed(w http.ResponseWriter, r *http.Request){
	userID := r.Context().Value(authentication.UserIDToken).(uint64)

	service, err := services.NewPostService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	posts, err := service.FindUserFeed(userID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data: posts,
	})
}