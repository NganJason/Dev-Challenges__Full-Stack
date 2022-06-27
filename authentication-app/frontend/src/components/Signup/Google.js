import React, { useEffect } from "react";
import jwt_decode from "jwt-decode";
import { NewService } from "../../service/service.js";

import googleLogo from "../../assets/Google.svg";

const Google_Client_ID =
  "347600384407-ne3gjt0942m016ciuu08vdlnhbtn9183.apps.googleusercontent.com";

function Google({
  setIsAuth,
  setUser,
}) {
    useEffect(() => {
      /* global google */
      google.accounts.id.initialize({
        client_id: Google_Client_ID,
        callback: handleCallbackResponse,
        ux_mode: "redirect",
      });
    }, []);

    const handleCallbackResponse = (resp) => {
      let user = jwt_decode(resp.credential);

      let s = NewService();

      s.GoogleLogin(user.email, user.sub)
        .then(function(resp) {
          setUser(resp);
          setIsAuth(true);
        })
        .catch(function(error){
          console.log(error.message)
        })
    };

    return (
      <img
        id="signin"
        src={googleLogo}
        alt="google-icon"
        onClick={() => {
          google.accounts.id.prompt((notification)=>{console.log(notification)});
        }}
      />
    );
}

export default Google

