import React, { useContext, useState, useEffect } from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ThemeContext } from "./hooks/themeContext";

import Home from "./components/Home/Home";
import Signup from "./components/Signup/Signup"
import EditInfo from "./components/EditInfo/EditInfo";
import Info from "./components/Info/Info";

import "./styles/main.scss"
import { useUserData } from "./hooks/useUserData";
import ProtectedRoute from "./components/ProtectedRoute/ProtectedRoute";

function App() {
  const [isAuth, setIsAuth] = useState(false);
  const { isDarkTheme } = useContext(ThemeContext);
  const { userData, editUserData } = useUserData();
  
  useEffect(() => {
    const code =
      window.location.href.match(/\?code=(.*)/) &&
      window.location.href.match(/\?code=(.*)/)[1];

    // if (code) {
    //   setIsAuth(true)
    // }
  }, []);

  return (
    <div className={`App bg-primary ${isDarkTheme ? "dark" : ""}`}>
      <BrowserRouter>
        <Routes>
          <Route path="signup" element={<Signup isSignup={true} />} />
          
          <Route element={<ProtectedRoute isAuth={isAuth} />}>
            <Route path="/" exact element={<Home />}>
              <Route exact path="" element={<Info userData={userData} />} />
              <Route
                path="edit-info"
                element={
                  <EditInfo userData={userData} setUserData={editUserData} />
                }
              />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
