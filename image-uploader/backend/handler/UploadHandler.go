package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/NganJason/Dev-Challenges__Full-Stack/tree/master/image-uploader/util"
	"github.com/NganJason/Dev-Challenges__Full-Stack/tree/master/image-uploader/vo"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	var resp vo.UploadImgResponse

	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile(vo.ImgFormKey)
	if err != nil {
		log.Println("error retrieving file")
		return
	}

	defer file.Close()

	if exist := isDirExist(vo.ImgDir); !exist {
		if err := os.Mkdir(vo.ImgDir, os.ModePerm); err != nil {
    	    log.Println(err)
			return
	    }	
	} 

	tempFile, err := ioutil.TempFile(vo.ImgDir, "*.png")
    if err != nil {
        log.Println(err)
		return
    }
    defer tempFile.Close()

    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
		return
    }    
    tempFile.Write(fileBytes)
	
	srvAddr := getServerAddress(r)
	fileAddr := getImgUrl(srvAddr, tempFile.Name())

	resp.Url = fileAddr
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("uploaded img successfully, %s\n", fileAddr)
	w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResp)
}

func isDirExist(path string) (bool) {
    _, err := os.Stat(path)
    if err == nil { 
		return true 
	}

    if os.IsNotExist(err) {
		return false 
	}

    return false
}

func getServerAddress(r *http.Request) string {
	ctx := r.Context()
	srvAddr := ctx.Value(http.LocalAddrContextKey).(net.Addr)

	s := strings.Replace(srvAddr.String(), util.LocalHostIP, util.LocalHostAddr, 1)
	return fmt.Sprintf("http://%s/", s)
}

func getImgUrl(serverAddr, path string) string {
	stringList := strings.Split(path, "/")
	imgPath := stringList[len(stringList) -1]

	return fmt.Sprintf("%sapi/img/get?%s=%s", serverAddr, vo.GetImgRequestParam, imgPath)
}
