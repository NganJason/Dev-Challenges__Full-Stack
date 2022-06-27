import React, { useContext, useState, useEffect } from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ThemeContext } from "./hooks/themeContext";
import { NewService } from "../src/service/service";

import Home from "./components/Home/Home";
import Signup from "./components/Signup/Signup"
import EditInfo from "./components/EditInfo/EditInfo";
import Info from "./components/Info/Info";

import "./styles/main.scss"
import { useUserInfo } from "./hooks/useUserInfo";
import ProtectedRoute from "./components/CustomRoute/ProtectedRoute";
import AuthRoute from "./components/CustomRoute/AuthRoute";

function App() {
  const [isAuth, setIsAuth] = useState(false);
  const { isDarkTheme } = useContext(ThemeContext);
  const { userInfo, setUser, updateUserInfo, fetchLatestUserInfo } = useUserInfo();
  
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
              element={<Signup isSignup={true} setIsAuth={setIsAuth} setUser={setUser}/>}
            />
            <Route
              path="login"
              element={<Signup isSignup={false} setIsAuth={setIsAuth} setUser={setUser}/>}
            />
          </Route>

          <Route
            element={<ProtectedRoute isAuth={isAuth} setIsAuth={setIsAuth} setUser={setUser}/>}
          >
            <Route path="/" exact element={<Home userInfo={userInfo} fetchLatestUserInfo={fetchLatestUserInfo}/>}>
              <Route exact path="" element={<Info userInfo={userInfo} />} />
              <Route
                path="edit-info"
                element={
                  <EditInfo userInfo={userInfo} updateUserInfo={updateUserInfo} fetchLatestUserInfo={fetchLatestUserInfo}/>
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
