import React from "react";
import FacebookLogin from "react-facebook-login/dist/facebook-login-render-props"; 
import facebookLogo from "../../assets/Facebook.svg";

import { NewService } from "../../service/service.js";

function Facebook({
  setIsAuth,
  setUser,
}) {
    const responseHandler = (r) => {
      const source =
        window.location.href.match(/\?state=(.*)/) &&
        window.location.href.match(/\?state=(.*)/)[1];
      
      if (source !== null && !source.includes("facebookdirect")) {
        return
      }

      if (r.accessToken === null || r.accessToken === undefined) {
        return;
      }

      const url = window.location.href;
      const newUrl = url.split("?code=");
      window.history.pushState({}, null, newUrl[0]);

      let s = NewService();
      s.FacebookLogin(r.accessToken)
        .then(function(resp) {
          setUser(resp);
          setIsAuth(true);
        })
        .catch(function(error) {
          console.log(error.message);
        })
    }

    return (
      <FacebookLogin
        appId="432017455175888"
        fields="name,email,picture"
        scope="public_profile,user_friends"
        callback={responseHandler}
        icon={facebookLogo}
        render={(renderProps) => (
          <img
            onClick={renderProps.onClick}
            src={facebookLogo}
            alt="facebook-icon"
          />
        )}
      />
    );
}

export default Facebook