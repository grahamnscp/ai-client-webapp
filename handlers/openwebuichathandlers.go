package handlers

import (
	"log"
	"net/http"

	u "webapp/utils"
  "webapp/aiclient"
)

// Local structs
type aiRespStr struct {
	Success  bool
	Prompt   string
	Response string
}

// Global variables
var aiResp = aiRespStr{}
var aiChat = []aiRespStr{}


/* OpenWebuiChat handler */
func OpenWebuiChat(w http.ResponseWriter, r *http.Request) {

	log.Println("OpenWebuiChat: called")

	log.Println("OpenWebuiChat: method:", r.Method) //get request method
	if r.Method == "GET" {
		// Initial page render
		aiResp = aiRespStr{
			Success:  false,
			Prompt:   "",
			Response: "",
		}
		aiChat = append(aiChat, aiResp)

		u.Render(w, "templates/OpenWebuiChat.html", aiChat)

    // flush chat history ready for response
    aiChat = nil

		return
	}

  // form POST received
	r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
	prompt := r.FormValue("prompt")

	log.Println("OpenWebuiChat: received prompt:", prompt)

	// Call AI with prompt ##############
  aiResponse := aiclient.AIClientQuery(prompt)

	log.Println("OpenWebuiChat: received ai response:", aiResponse)

	// populate the ui form struct
	aiResp = aiRespStr{
		Success:  true,
		Prompt:   prompt,
		Response: aiResponse,
	}
	aiChat = append(aiChat, aiResp)

	// render response
	log.Println("OpenWebuiChat: aiChat:", aiChat)
	u.Render(w, "templates/OpenWebuiChat.html", aiChat)
}
