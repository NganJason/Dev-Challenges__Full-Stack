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
          .then(function(resp) {
            return resp.data.user_info;
          })
          .catch(function(error) {
            throw new Error(error.response.data.debug_msg);
          }); 
    }

    GoogleLogin(email, subID) {
      let url = this.baseURL + "login/google"

      return axios.post(
        url, 
        {
          email: email,
          sub_id: subID
        },
        { withCredentials: true }
      )
      .then(function(resp) {
        return resp.data.user_info;
      })
      .catch(function(error) {
        throw new Error(error.response.data.debug_msg)
      });
    }

    FacebookLogin(accessCode) {
      let url = this.baseURL + "login/facebook"

      return axios.post(
        url,
        {
          access_code: accessCode,
        },
        { withCredentials: true }
      )
      .then(function(resp) {
        return resp.data.user_info;
      })
      .catch(function(error) {
        throw new Error(error.response.data.debug_msg)
      })
    }

    DefaultSignup(username, password) {
      let url = this.baseURL + "signup"

      return axios.post(
        url,
        {
          username: username,
          password: password,
        },
        { withCredentials: true }
      )
      .then(function(resp) {
        return resp.data.user_info;
      })
      .catch(function(error) {
        throw new Error(error.response.data.debug_msg)
      })
    }

    DefaultLogin(username, password) {
      let url = this.baseURL + "login"

      return axios.post(
        url,
        {
          username: username,
          password: password,
        },
        { withCredentials: true }
      )
      .then(function(resp) {
        return resp.data.user_info;
      })
      .catch(function(error) {
        throw new Error(error.response.data.debug_msg)
      })
    }
    
    VerifyAuth() {
      let url = this.baseURL + "login/verify_auth"

      return axios.get(
        url,
        { withCredentials: true }
      ).then(function (resp) {
        return resp;
      }).catch(function (error) {
        throw new Error(error.response.data.debug_msg)
      })
    }

    UpdateUserInfo(req) {
      let url = this.baseURL + "user_info/update"

      let request = {
          user_id: req.user_id,
          username: req.username,
          bio: req.bio,
          phone: req.phone,
          email: req.email,
        }

      return axios
        .post(url, request, { withCredentials: true })
        .then(function (resp) {
          return resp.data.user_info;
        })
        .catch(function (error) {
          throw new Error(error.response.data.debug_msg);
        });
    }

    GetUserInfo(userID) {
      let url = this.baseURL + "user_info/get"

      return axios.post(
        url,
        {
          user_id: userID,
        },
        { withCredentials: true }
      ).then(function (resp) {
        return resp.data.user_info;
      }).catch(function (error) {
        throw new Error(error.response.data.debug_msg)
      })
    }

    Logout() {
      let url = this.baseURL + "logout"

      return axios.post(
        url,
        {},
        { withCredentials: true }
        ).then(function (resp) {
          return resp
        }).catch(function (error) {
          throw new Error(error.response.data.debug_msg)
        })
    }
}