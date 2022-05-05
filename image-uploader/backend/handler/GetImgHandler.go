package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/NganJason/Dev-Challenges__Full-Stack/tree/master/image-uploader/vo"
)

func GetImgHandler(w http.ResponseWriter, r *http.Request) {
	paths, ok := r.URL.Query()[vo.GetImgRequestParam]
	if !ok {
		log.Println("cannot get param path from url")
		return
	}

	imgPath := getFilePath(paths[0])

	fileBytes, err := ioutil.ReadFile(imgPath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

func getFilePath(path string) string {
	return fmt.Sprintf("%s/%s", vo.ImgDir, path)
}