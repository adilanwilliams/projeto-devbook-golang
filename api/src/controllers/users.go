package controllers

import "net/http"

// CreateUser insert a new user in database.
func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Creating user."))
}

// UpdateUser update a user existend in database.
func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Updating user."))
}

// DeleteUser delete a user existend in database.
func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deleting user."))
}

// UpdateUser fetch all users in database.
func GetAllUsers(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Getting all users."))
}