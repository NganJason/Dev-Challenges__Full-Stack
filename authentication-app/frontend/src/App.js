import React, { useState } from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Button from "./components/Button/Button";
import Signup from "./components/Signup/Signup"

import "./styles/main.scss"

function App() {
  const [isDarkTheme, setIsDarkTheme] = useState(false)

  const toggleTheme = () => {
    setIsDarkTheme(!isDarkTheme)
  }
  return (
    <div className={`App bg-primary ${isDarkTheme ? "dark" : ""}`}>
      <Button
        onClick={toggleTheme}
        width="2rem"
        height="2rem"
        color="inherit"
        round
      >
        <span class="material-icons theme-toggler">
          {isDarkTheme ? "dark_mode" : "light_mode"}
        </span>
      </Button>

      <BrowserRouter>
        <Routes>
          <Route
            path="signup"
            element={<Signup isSignup={true} isDarkTheme={isDarkTheme} />}
          />
          <Route
            path="login"
            element={<Signup isSignup={false} isDarkTheme={isDarkTheme} />}
          />
          <Route
            path="/"
            exact
            element={<Signup isSignup={true} isDarkTheme={isDarkTheme} />}
          />
          
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
