package processor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/service"
	"github.com/NganJason/Dev-Challenges__Full-Stack/auth-app/internal/vo"
)

func GithubLoginProcessor(w http.ResponseWriter, r *http.Request) {
	var req vo.GithubLoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println(err.Error())
		return
	}

	s := service.NewGithubService()
	resp, err := s.Login(req.AccessCode, "")
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(resp)
}
