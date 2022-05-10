import React, { useState } from "react";

import { Outlet } from "react-router-dom";

import Nav from "../Nav/Nav"

function Home() {
  const [showModal, setShowModal] = useState(false);

  const toggleModal = () => {
    setShowModal(!showModal);
  };

    return (
      <div className="home">
        <Nav showModal={showModal} toggleModal={toggleModal} />
        <div
          onClick={() => {
            setShowModal(false);
          }}
        >
          <Outlet />
        </div>
      </div>
    );
}

export default Home;