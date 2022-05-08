import React from "react";

import Text from "../Text/Text";
import Button from "../Button/Button"

function EditInfo() {
  return (
    <div className="edit-info">
      <div className="edit__link">
        <span class="material-icons md-18 inline link">arrow_back_ios</span>
        <a href="http://localhost:3001/">Back</a>
      </div>

      <div className="edit-dashboard">
        <section>
          <Text size="1.5rem" bd="400" align="left" mgTop="1.5">
            Personal Info
          </Text>
          <Text size="0.9rem" bd="300" align="left" mgTop="0.5">
            Basic info, like your name and photo
          </Text>
        </section>

        <div className="edit__content">
          <img src="https://64.media.tumblr.com/65390415cadbd069148850d7cdcbedd1/tumblr_ozgyzr929N1qjoleso1_1280.jpg" />

          <Text size="0.8rem" bd="500" tertiary>
            Name
          </Text>
          <input placeholder="Enter your name..." type="text" />

          <Text size="0.8rem" bd="500" tertiary>
            Bio
          </Text>
          <input placeholder="Enter your bio..." type="text" />

          <Text size="0.8rem" bd="500" tertiary>
            Phone
          </Text>
          <input placeholder="Enter your phone..." type="text" />

          <Text size="0.8rem" bd="500" tertiary>
            Email
          </Text>
          <input placeholder="Enter your email..." type="text" />

          <Text size="0.8rem" bd="500" tertiary>
            Password
          </Text>
          <input placeholder="Enter your password..." type="text" />

          <Button border="8px" width="5rem" height="2.3rem" mgTop="1.5">
            Save
          </Button>
        </div>
      </div>
    </div>
  );
}

export default EditInfo;
