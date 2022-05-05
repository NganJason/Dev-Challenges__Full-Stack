import FormData from "form-data";
import axios from "axios"

const ImgServiceBaseUrl = "http://localhost:8082/api/";

class ImgService {
  constructor() {
    this.instance = axios.create({
      baseURL: ImgServiceBaseUrl,
    });
  }

  uploadImg(file) {
    var formData = new FormData();
    formData.append("image", file);

    try {
      let res = this.instance.post("img/upload", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });

      return res
    } catch(e) {
      return null
    }
  }
}

export default ImgService