import React from "react";
import { Outlet } from "react-router-dom";
import Signup from "../Signup/Signup";

function ProtectedRoute({
    isAuth, 
    redirectPath="/login",
    children,
}) {
  if (!isAuth) {
    return <Signup isSignup={false}/>
  }

  return children ? children : <Outlet />
}

export default ProtectedRoute;
