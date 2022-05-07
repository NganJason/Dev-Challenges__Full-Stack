import React, { useState } from "react";
import Text from "../Text/Text"
import Button from "../Button/Button"

import lightThemeDevLogo from "../../assets/devchallenges.svg"
import darkThemeDevLogo from "../../assets/devchallenges-light.svg";
import githubLogo from "../../assets/Github.svg"
import twitterLogo from "../../assets/Twitter.svg";
import googleLogo from "../../assets/Google.svg";
import facebookLogo from "../../assets/Facebook.svg";

function Signup({isSignup, isDarkTheme}) {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const onUsernameChange = (e) => {
    setUsername(e.target.value)
  }

  const onPasswordChange = (e) => {
    setPassword(e.target.value)
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
        <img src={googleLogo} alt="google-icon" />
        <img src={facebookLogo} alt="facebook-icon" />
        <img src={twitterLogo} alt="twitter-icon" />
        <img src={githubLogo} alt="github-icon" />
      </div>

      <Text bd="400" size="0.9rem" align="center" color="secondary">
        {getFormLink()}
      </Text>
    </div>
  );
}

export default Signup;
