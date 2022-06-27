import React from "react";
import { Outlet } from "react-router-dom";
import Signup from "../Signup/Signup";

function ProtectedRoute({
    isAuth,
    setIsAuth,
    setUser,
    children,
}) {
  if (!isAuth) {
    return <Signup isSignup={false} setIsAuth={setIsAuth} setUser={setUser}/>
  }

  return children ? children : <Outlet />
}

export default ProtectedRoute;
