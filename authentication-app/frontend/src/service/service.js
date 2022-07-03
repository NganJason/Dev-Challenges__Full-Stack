import axios from "axios"

export const NewService = () => {
    let baseURL = process.env.REACT_APP_BACKEND_API;
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
            console.log(error);
            throw error;
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
        throw error;
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
        throw error;
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
        throw error;
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
        throw error
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
        throw error;
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
          throw error;
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
        throw error;
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
          throw error;
        })
    }
}