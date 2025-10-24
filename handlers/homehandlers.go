package handlers

import (
	"log"
	"net/http"

	u "webapp/utils"
)

/* Index Home */
func Home(w http.ResponseWriter, r *http.Request) {

	log.Println("Home: called")
	u.Render(w, "templates/Home.html", nil)
}
