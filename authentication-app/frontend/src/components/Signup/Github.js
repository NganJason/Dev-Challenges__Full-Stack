import React, { useEffect } from "react"

import githubLogo from "../../assets/Github.svg";
import { NewService } from "../../service/service.js"

function Github({
  setIsAuth,
  setUser,
  setErrorAlert,
}) {
    const onGithub = () => {
      let CLIENT_ID = "85de73f0c04a2f06d9d5";
      let REDIRECT_URI = window.location.href + "?source=github";
      let url = `https://github.com/login/oauth/authorize?client_id=${CLIENT_ID}&scope=user&redirect_uri=${REDIRECT_URI}`;
      window.location.href = url;
    };

    useEffect(() => {
        const source =
          window.location.href.match(/\?source=(.*)/) &&
          window.location.href.match(/\?source=(.*)/)[1];

        const code =
          window.location.href.match(/\&code=(.*)/) &&
          window.location.href.match(/\&code=(.*)/)[1];

        let isSourceGithub = false
        if (source !== null) {
          isSourceGithub = source.includes("github");
        }

        if (code === "" || code === null || !isSourceGithub) {
            return
        }

        const url = window.location.href;
        const newUrl = url.split("?source=");
        window.history.pushState({}, null, newUrl[0]);

        let s = NewService();

        s.GithubLogin(code)
          .then(function(resp) {
            setUser(resp)
            setIsAuth(true);
          })
          .catch(function(err) {
            console.log(err);
            setErrorAlert(err);
          });
    }, [])

    return <img src={githubLogo} alt="github-icon" onClick={onGithub} />;
}

export default Github