import React, { useContext } from "react"
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { ThemeContext } from "./hooks/themeContext";

import Home from "./components/Home/Home";
import Signup from "./components/Signup/Signup"
import EditInfo from "./components/EditInfo/EditInfo";
import Info from "./components/Info/Info";

import "./styles/main.scss"
import { useUserData } from "./hooks/useUserData";

function App() {
  const { isDarkTheme } = useContext(ThemeContext);

  const { userData, editUserData } = useUserData();
  return (
    <div className={`App bg-primary ${isDarkTheme ? "dark" : ""}`}>
      <BrowserRouter>
        <Routes>
          <Route path="signup" element={<Signup isSignup={true} />} />
          <Route path="login" element={<Signup isSignup={false} />} />

          <Route path="/" exact element={<Home />}>
            <Route exact path="" element={<Info userData={userData} />} />
            <Route
              path="edit-info"
              element={
                <EditInfo userData={userData} setUserData={editUserData} />
              }
            />
          </Route>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
