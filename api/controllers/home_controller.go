package controllers

import (
	"net/http"

	reponses "github.com/selvakreditbee/fullstack/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	reponses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}