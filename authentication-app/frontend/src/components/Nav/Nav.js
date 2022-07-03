import React, { useContext } from "react";
import { ThemeContext } from "../../hooks/themeContext";
import { Link } from "react-router-dom";

import Text from "../Text/Text";
import Button from "../Button/Button";
import DropdownModal from "./DropdownModal";

import lightThemeLogo from "../../assets/devchallenges.svg";
import darkThemeLogo from "../../assets/devchallenges-light.svg";

function Nav({showModal, toggleModal, userInfo}) {
  const { toggleIsDarkTheme, isDarkTheme } = useContext(ThemeContext);

  return (
    <div className="nav">
      <Link to="/" style={{ width: "auto" }}>
        <img
          className="nav__logo"
          alt="logo"
          src={isDarkTheme ? darkThemeLogo : lightThemeLogo}
        />
      </Link>

      <div className="nav__profile">
        <Text size="0.8rem" bd="700">
          {userInfo["username"] ? userInfo["username"] : "New User"}
        </Text>

        <span className="material-icons md-36 primary" onClick={toggleModal}>
          arrow_drop_down
        </span>

        <DropdownModal toggleModal={toggleModal} show={showModal} />

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
