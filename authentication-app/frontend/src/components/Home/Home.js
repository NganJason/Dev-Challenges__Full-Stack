import React from "react";

import { Outlet } from "react-router-dom";

import Nav from "../Nav/Nav"

function Home() {
    return (
      <div className="home">
        <Nav />
        <Outlet />
      </div>
    );
}

export default Home;