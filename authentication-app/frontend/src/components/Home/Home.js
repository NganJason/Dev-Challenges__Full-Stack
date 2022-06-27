import React, { useState, useEffect } from "react";

import { Outlet } from "react-router-dom";

import Nav from "../Nav/Nav"

function Home({userInfo, fetchLatestUserInfo}) {
  const [showModal, setShowModal] = useState(false);

  useEffect(() => {
    fetchLatestUserInfo(userInfo.user_id)
  })

  const toggleModal = () => {
    setShowModal(!showModal);
  };

    return (
      <div className="home">
        <Nav showModal={showModal} toggleModal={toggleModal} userInfo={userInfo}/>
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