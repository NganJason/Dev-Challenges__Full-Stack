import React from "react";
import Info from "../Info/Info";

import Nav from "../Nav/Nav"

function Home() {
    return (
        <div className="home">
            <Nav/>
            <Info/>          
        </div>
    );
}

export default Home;