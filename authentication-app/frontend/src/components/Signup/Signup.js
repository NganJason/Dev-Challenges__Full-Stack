import React, { useState, useContext } from "react";
import { ThemeContext } from "../../hooks/themeContext";

import Text from "../Text/Text"
import Button from "../Button/Button"
import Github from "./Github";

import lightThemeDevLogo from "../../assets/devchallenges.svg"
import darkThemeDevLogo from "../../assets/devchallenges-light.svg";

import Google from "./Google";
import Facebook from "./Facebook";
import { NewService } from "../../service/service";


function Signup({
  isSignup, 
  setIsAuth, 
  setUser,
}) {
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const [errorLog, setErrorLog] = useState("error")
  const [showError, setShowError] = useState(false)

  const { isDarkTheme } = useContext(ThemeContext);

  const onUsernameChange = (e) => {
    setUsername(e.target.value)
  }

  const onPasswordChange = (e) => {
    setPassword(e.target.value)
  }

  const onSubmit = () => {
    let s = NewService()
    const url = "http://localhost:3001";

    if (isSignup) {
      s.DefaultSignup(username, password)
      .then(function(resp) {
        setUser(resp);
        setIsAuth(true);
        
        window.history.pushState({}, null, url);
        window.location.reload(true);
      })
      .catch(function(err) { 
        console.log(err)
        setErrorAlert(err);
      })
    } else {
      s.DefaultLogin(username, password)
        .then(function (resp) {
          setUser(resp);
          setIsAuth(true);
          
          window.history.pushState({}, null, url);
          window.location.reload(true);
        })
        .catch(function (err) {
          console.log(err)
          setErrorAlert(err)
        });
    }

    clearFields()
  }

  const clearFields = () => {
    setUsername("")
    setPassword("")
  }

  const setErrorAlert = (err) => {
    if (err.response.status !== 401 && err.response.status !== 400) {
      return
    }

    let msg = err.response.data.debug_msg;
    msg = msg[0].toUpperCase() + msg.substring(1)
    setErrorLog(msg);
    setShowError(true);

    setInterval(() => {
      setShowError(false)
    }, 4000)
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
          Already a member? <a href="http://localhost:3001/auth/login">Login</a>
        </div>
      )
    } else {
      return (
        <div>
          Don't have an account yet?{" "}
          <a href="http://localhost:3001/auth/signup">Register</a>
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

      <div className={"msg " + (showError ? "alert-shown" : "alert-hidden")}>
        {errorLog}
      </div>
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
        <Google
          setIsAuth={setIsAuth}
          setUser={setUser}
          setErrorAlert={setErrorAlert}
        />
        <Facebook
          setIsAuth={setIsAuth}
          setUser={setUser}
          setErrorAlert={setErrorAlert}
        />
        <Github
          setIsAuth={setIsAuth}
          setUser={setUser}
          setErrorAlert={setErrorAlert}
        />
      </div>
      <Text bd="400" size="0.9rem" align="center" color="secondary">
        {getFormLink()}
      </Text>
    </div>
  );
}

export default Signup;
