import React from "react";
import { Navigate, Outlet } from "react-router-dom";

function ProtectedRoute({
    isAuth, 
    redirectPath="/login",
    children,
}) {
  if (!isAuth) {
      return <Navigate to={redirectPath} replace />;
  }

  return children ? children : <Outlet />
}

export default ProtectedRoute;
