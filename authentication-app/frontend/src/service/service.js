import axios from "axios"

export const NewService = () => {
    let baseURL = "http://localhost:8082/api/"
    let s = new Service(baseURL)
    
    return s
}

class Service {
    constructor(baseURL) {
        this.baseURL = baseURL
    }

    GithubLogin(accessCode) {
        let url = this.baseURL + "login/github"

        return axios
          .post(
            url,
            {
              access_code: accessCode,
            },
            { withCredentials: true }
          )
          .then(function (resp) {
            return resp;
          })
          .catch(function (error) {
            throw new Error(error.response.data.debug_msg);
          }); 
    }

    VerifyAuth() {
      let url = this.baseURL + "login/verify_auth"

      return axios.get(
        url, {
          withCredentials: true
        }
      ).then(function (resp) {
        return resp;
      }).catch(function (error) {
        throw new Error(error.response.data.debug_msg)
      })
    }
}