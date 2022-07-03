import React, {useEffect} from "react";
import { Outlet } from "react-router-dom";

function AuthRoute({ isAuth, children }) {
  useEffect(() => {
    if (isAuth) {
        const url = "/";
        window.history.pushState({}, null, url);
        window.location.reload(true);
    }
  }, [isAuth]);

  return children ? children : <Outlet />;
}

export default AuthRoute;
