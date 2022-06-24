import React, { useContext, useState, useEffect } from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ThemeContext } from "./hooks/themeContext";
import { NewService } from "../src/service/service";

import Home from "./components/Home/Home";
import Signup from "./components/Signup/Signup"
import EditInfo from "./components/EditInfo/EditInfo";
import Info from "./components/Info/Info";

import "./styles/main.scss"
import { useUserData } from "./hooks/useUserData";
import ProtectedRoute from "./components/CustomRoute/ProtectedRoute";
import AuthRoute from "./components/CustomRoute/AuthRoute";

function App() {
  const [isAuth, setIsAuth] = useState(false);
  const { isDarkTheme } = useContext(ThemeContext);
  const { userData, editUserData } = useUserData();
  
  useEffect(() => {
    let s = NewService();

    s.VerifyAuth()
      .then(function(resp) {
        setIsAuth(resp.data.is_auth)
      })
      .catch(function(error) {
        console.log(error)
      })
  }, [])

  return (
    <div className={`App bg-primary ${isDarkTheme ? "dark" : ""}`}>
      <BrowserRouter>
        <Routes>
          <Route path="/auth" element={<AuthRoute isAuth={isAuth} />}>
            <Route
              path="signup"
              element={<Signup isSignup={true} setIsAuth={setIsAuth} />}
            />
            <Route
              path="login"
              element={<Signup isSignup={false} setIsAuth={setIsAuth} />}
            />
          </Route>

          <Route
            element={<ProtectedRoute isAuth={isAuth} setIsAuth={setIsAuth} />}
          >
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
