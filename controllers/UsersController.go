package controllers

import "net/http"

type UsersController struct{}

func (controller *UsersController) NewUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("new user"))
}
