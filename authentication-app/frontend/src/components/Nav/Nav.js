import React, { useState, useContext } from "react";
import { ThemeContext } from "../../hooks/themeContext";

import Text from "../Text/Text";
import Button from "../Button/Button";
import DropdownModal from "./DropdownModal";

import lightThemeLogo from "../../assets/devchallenges.svg";
import darkThemeLogo from "../../assets/devchallenges-light.svg";

function Nav({showModal, toggleModal}) {
  const { toggleIsDarkTheme, isDarkTheme } = useContext(ThemeContext);

  return (
    <div className="nav">
      <img
        className="nav__logo"
        src={isDarkTheme ? darkThemeLogo : lightThemeLogo}
      />

      <div className="nav__profile">
        <img src="https://64.media.tumblr.com/65390415cadbd069148850d7cdcbedd1/tumblr_ozgyzr929N1qjoleso1_1280.jpg" />

        <Text size="0.8rem" bd="700">
          Jason Ngan
        </Text>

        <span className="material-icons md-36 primary" onClick={toggleModal}>
          arrow_drop_down
        </span>

        <DropdownModal show={showModal} />

        <Button
          onClick={toggleIsDarkTheme}
          width="2rem"
          height="2rem"
          align="right"
          round
          color="inherit"
        >
          <span class="material-icons theme-toggler">
            {isDarkTheme ? "dark_mode" : "light_mode"}
          </span>
        </Button>
      </div>
    </div>
  );
}

export default Nav;
