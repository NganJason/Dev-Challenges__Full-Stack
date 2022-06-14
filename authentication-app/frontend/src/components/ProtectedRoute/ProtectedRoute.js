import React from "react";
import { Outlet } from "react-router-dom";
import Signup from "../Signup/Signup";

function ProtectedRoute({
    isAuth,
    setIsAuth,
    children,
}) {
  if (!isAuth) {
    return <Signup isSignup={false} setIsAuth={setIsAuth}/>
  }

  return children ? children : <Outlet />
}

export default ProtectedRoute;
