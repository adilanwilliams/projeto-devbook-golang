package controllers

import (
	"devbook/src/authentication"
	"devbook/src/models"
	"devbook/src/services"
	"devbook/src/utils/response"
	"encoding/json"
	"errors"
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
		Data:    post,
	})
}

// FindUserFeed handles the request to retrieve the authenticated user's feed.
// The feed includes posts created by the user and by the users they follow.
func FindUserFeed(w http.ResponseWriter, r *http.Request) {
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
		Data:    posts,
	})
}

// UpdateUser handles the update of an existing user identified by postId.
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(authentication.UserIDToken).(uint64)
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

	if post.AuthorID != userID {
		err = errors.New("you are not allowed to update another user's post")
		response.ResponseError(w, http.StatusForbidden, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.ResponseError(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := json.Unmarshal(body, &post); err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	post.ID = postID

	err = service.UpdatePost(post)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{})

}

// DeletePost handles the removal of a post identified by postId.
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(authentication.UserIDToken).(uint64)
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

	if post.AuthorID != userID {
		err = errors.New("you are not allowed to delete another user's post")
		response.ResponseError(w, http.StatusForbidden, err)
		return
	}

	err = service.DeletePost(postID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusNoContent, response.Response{})
}

// FindPostsByUser handles the request to retrieve all posts created by a specific user.
// It reads the user ID from the request parameters and returns the user's posts in JSON format.
func FindPostsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, err)
		return
	}

	service, err := services.NewPostService()
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	posts, err := service.FindPostsByUser(userID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusOK, response.Response{
		Success: true,
		Data:    posts,
	})

}

func LikePost(w http.ResponseWriter, r *http.Request) {
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

	err = service.LikePost(postID)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, err)
		return
	}

	response.ResponseJSON(w, http.StatusNoContent, response.Response{})
}
