package processor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/service"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/vo"
)

func GithubLoginProcessor(w http.ResponseWriter, r *http.Request) {
	var req vo.GithubLoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("Received access code=%s", req.AccessCode)

	s := service.NewGithubService()
	resp, err := s.Login(req.AccessCode, "")
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(resp)
}
