// Package model ...
// this model file  holds structs that are to be used by the admin handler.
package model

// AdminLoginResponse to be usedby the admin response class
type AdminLoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Admin   *Admin `json:"admin"`
}

// LoginResponse to be usedby the admin response class
type LoginResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	User    interface{} `json:"user"`
}

// SimpleSuccessNotifier ...
type SimpleSuccessNotifier struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ShortSuccess
type ShortSuccess struct {
	Msg string `json:"msg"`
}
