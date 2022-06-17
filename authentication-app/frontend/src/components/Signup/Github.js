import React, { useEffect } from "react"

import githubLogo from "../../assets/Github.svg";
import { NewService } from "../../service/service.js"

function Github({setIsAuth}) {
    const onGithub = () => {
      let CLIENT_ID = "85de73f0c04a2f06d9d5";
      let REDIRECT_URI = "http://localhost:3001/";
      let url = `https://github.com/login/oauth/authorize?client_id=${CLIENT_ID}&scope=user&redirect_uri=${REDIRECT_URI}`;
      window.location.href = url;
    };

    useEffect(() => {
        const code =
          window.location.href.match(/\?code=(.*)/) &&
          window.location.href.match(/\?code=(.*)/)[1];

        if (code === "" || code === null) {
            return
        }

        const url = window.location.href;
        const newUrl = url.split("?code=");
        window.history.pushState({}, null, newUrl[0]);

        let s = NewService();

        s.GithubLogin(code)
          .then(function() {
            setIsAuth(true);
          })
          .catch(function(error) {
            console.log(error.message);
          });
    }, [])

    return <img src={githubLogo} alt="github-icon" onClick={onGithub} />;
}

export default Github