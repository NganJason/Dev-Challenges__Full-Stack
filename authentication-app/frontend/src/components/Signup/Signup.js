import React, { useState } from "react";
import Text from "../Text/Text"
import Button from "../Button/Button"

import devLogo from "../../assets/devchallenges.svg"
import githubLogo from "../../assets/Github.svg"
import twitterLogo from "../../assets/Twitter.svg";
import googleLogo from "../../assets/Google.svg";
import facebookLogo from "../../assets/Facebook.svg";

function Signup() {
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

  return (
    <div className="signup">
      <img className="signup__logo" src={devLogo} />
      <Text size="1.1rem" bd="600" mgTop="1">
        Join thousands of learners from around the world
      </Text>

      <Text size="0.9rem" bd="400" mgTop="1">
        Master web development by making real-life projects. There are multiple
        paths for you to choose.
      </Text>

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
        mgTop="1"
        mgBtm="1.5"
        onClick={onSubmit}
      >
        Start coding now
      </Button>
      <Text bd="400" size="0.9rem" color="secondary" align="center">
        or continue with these social profile
      </Text>

      <div className="social-icons">
        <img src={googleLogo} />
        <img src={facebookLogo} />
        <img src={twitterLogo} />
        <img src={githubLogo} />
      </div>

      <Text bd="400" size="0.9rem" align="center" color="secondary">
        Already a member? <a>Login</a>
      </Text>
    </div>
  );
}

export default Signup;
