import React from "react";
import lightThemeLogo from "../../assets/devchallenges.svg";
import darkThemeLogo from "../../assets/devchallenges-light.svg";
import Text from "../Text/Text";

function Nav() {
  return (
    <div className="nav">
      <img className="nav__logo" src={lightThemeLogo} />

      <div className="nav__profile">
        <img src="https://64.media.tumblr.com/65390415cadbd069148850d7cdcbedd1/tumblr_ozgyzr929N1qjoleso1_1280.jpg" />

        <Text size="0.8rem" bd="700">
          Jason Ngan
        </Text>

        <span class="material-icons md-36 primary">arrow_drop_down</span>
      </div>
    </div>
  );
}

export default Nav;
