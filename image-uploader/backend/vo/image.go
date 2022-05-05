package vo

const (
	ImgFormKey = "image"
	ImgDir = "image"
)

const GetImgRequestParam = "path"

type UploadImgResponse struct {
	Url string `json:"url"`
	DebugMsg string `json:"debug_msg"`
}