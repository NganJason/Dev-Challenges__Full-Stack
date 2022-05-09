import React, { useContext } from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ThemeContext } from "./hooks/themeContext";

import Button from "./components/Button/Button";
import Home from "./components/Home/Home";
import Signup from "./components/Signup/Signup"
import EditInfo from "./components/EditInfo/EditInfo";
import Info from "./components/Info/Info";

import "./styles/main.scss"

function App() {
  const { isDarkTheme, toggleIsDarkTheme } = useContext(ThemeContext);

  return (
      <div className={`App bg-primary ${isDarkTheme ? "dark" : ""}`}>
        <BrowserRouter>
          <Routes>
            <Route
              path="signup"
              element={<Signup isSignup={true} />}
            />
            <Route
              path="login"
              element={<Signup isSignup={false} />}
            />

            <Route path="/" exact element={<Home />}>
              <Route exact path="" element={<Info />} />
              <Route path="edit-info" element={<EditInfo />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </div>
  );
}

export default App;
