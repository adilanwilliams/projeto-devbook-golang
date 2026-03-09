package models

// UserPassword represents the data required to change a user's password.
// It contains the current password for validation and the new password
// that will replace the existing one.
type UserPassword struct {
	Current string `json:"current"`
	New     string `json:"new"`
}