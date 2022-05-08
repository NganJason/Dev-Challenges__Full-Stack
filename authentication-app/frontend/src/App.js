import React, { useState } from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Button from "./components/Button/Button";
import Home from "./components/Home/Home";
import Signup from "./components/Signup/Signup"
import EditInfo from "./components/EditInfo/EditInfo";
import Info from "./components/Info/Info";

import "./styles/main.scss"

function App() {
  const [isDarkTheme, setIsDarkTheme] = useState(false)

  const toggleTheme = () => {
    setIsDarkTheme(!isDarkTheme)
  }
  return (
    <div className={`App bg-primary ${isDarkTheme ? "dark" : ""}`}>
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
          
          <Route path="/" exact element={<Home />}>
            <Route exact path="" element={<Info />} />
            <Route path="edit-info" element={<EditInfo />} />
          </Route>
        </Routes>
      </BrowserRouter>

      <Button
        onClick={toggleTheme}
        width="2rem"
        height="2rem"
        align="right"
        round
        color="inherit"
      >
        <span class="material-icons theme-toggler">
          {isDarkTheme ? "dark_mode" : "light_mode"}
        </span>
      </Button>
    </div>
  );
}

export default App;
