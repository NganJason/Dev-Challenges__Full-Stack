import React, { useState, useContext } from "react";
import { ThemeContext } from "../../hooks/themeContext";
import GoogleLogin from "react-google-login";

import Text from "../Text/Text"
import Button from "../Button/Button"

import lightThemeDevLogo from "../../assets/devchallenges.svg"
import darkThemeDevLogo from "../../assets/devchallenges-light.svg";
import githubLogo from "../../assets/Github.svg"
import twitterLogo from "../../assets/Twitter.svg";
import googleLogo from "../../assets/Google.svg";
import facebookLogo from "../../assets/Facebook.svg";

const Google_Client_ID =
  "347600384407-76au7p6cbmgkb26fr26bp56o98ooks2e.apps.googleusercontent.com";

function Signup({isSignup}) {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const { isDarkTheme } = useContext(ThemeContext);

  const onUsernameChange = (e) => {
    setUsername(e.target.value)
  }

  const onPasswordChange = (e) => {
    setPassword(e.target.value)
  }

  const onGithub = (e) => {
    let CLIENT_ID = "85de73f0c04a2f06d9d5";
    let REDIRECT_URI = "http://localhost:3001/";
    let url = `https://github.com/login/oauth/authorize?client_id=${CLIENT_ID}&scope=user&redirect_uri=${REDIRECT_URI}`;
    window.location.href = url
  }

  const onGoogleSuccess = (res) => {
    console.log("success", res)
  }

  const onGoogleFail = (res) => {
    console.log("fail", res)
  }

  const onSubmit = () => {
    console.log(username, password)
  }

  const getFormTitle = () => {
    if (isSignup) {
      return (
        <div>
          <Text size="1.1rem" bd="600" mgTop="1">
            Join thousands of learners from around the world
          </Text>

          <Text size="0.9rem" bd="400" mgTop="1">
            Master web development by making real-life projects. There are
            multiple paths for you to choose.
          </Text>
        </div>
      );
    } else {
      return (
        <div>
          <Text size="1.1rem" bd="600" mgTop="1">
            Login
          </Text>
        </div>
      );
    }
  }

  const getFormLink = () => {
    if (isSignup) {
      return (
        <div>
          Already a member? <a href="http://localhost:3001/login">Login</a>
        </div>
      )
    } else {
      return (
        <div>
          Don't have an account yet?{" "}
          <a href="http://localhost:3001/signup">Register</a>
        </div>
      ); 
    }
  }

  return (
    <div className="signup">
      <img
        className="signup__logo"
        src={isDarkTheme ? darkThemeDevLogo : lightThemeDevLogo}
        alt="form-logo"
      />

      {getFormTitle()}

      <input
        type="text"
        placeholder="Email"
        onChange={onUsernameChange}
        value={username}
      ></input>

      <input
        type="password"
        placeholder="Password"
        onChange={onPasswordChange}
        value={password}
      ></input>

      <Button
        width="100%"
        height="2.5rem"
        border="8px"
        mgTop="1.5"
        mgBtm="1.5"
        onClick={onSubmit}
      >
        {isSignup ? "Start coding now" : "Login"}
      </Button>

      <Text bd="400" size="0.9rem" color="secondary" align="center">
        or continue with these social profile
      </Text>

      <div className="social-icons">
        <GoogleLogin
          clientId={Google_Client_ID}
          onSuccess={onGoogleSuccess}
          onFailure={onGoogleFail}
          cookiePolicy={"single_host_origin"}
          isSignedIn={true}
          render={(renderProps) => (
            <img
              src={googleLogo}
              alt="google-icon"
              onClick={renderProps.onClick}
            />
          )}
          buttonText="Login"
        />
        <img src={facebookLogo} alt="facebook-icon" />
        <img src={twitterLogo} alt="twitter-icon" />
        <img src={githubLogo} alt="github-icon" onClick={onGithub} />
      </div>
      <Text bd="400" size="0.9rem" align="center" color="secondary">
        {getFormLink()}
      </Text>
    </div>
  );
}

export default Signup;
