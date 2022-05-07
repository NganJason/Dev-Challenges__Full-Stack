import React from "react";
import Text from "../Text/Text";
import Button from "../Button/Button"

function Info() {
  return (
    <div className="info">
      <section>
        <Text size="2.5rem" bd="400" align="center" mgTop="1.5">
          Personal Info
        </Text>
        <Text size="1.2rem" bd="300" align="center" mgTop="1">
          Basic info, like your name and photo
        </Text>
      </section>

      <div className="info__board">
        <div className="board__header">
          <div className="header__content">
            <Text bd="400" size="1.5rem" color="primary">
              Profile
            </Text>
            <Text bd="500" size="0.8rem" color="secondary">
              Some info may be visible to other people
            </Text>
          </div>

          <Button color="secondary" border="12px">
            <Text bd="500" size="1rem" color="secondary">
              Edit
            </Text>
          </Button>
        </div>

        <div className="board__photo">
          <Text bd="500" size="0.9rem" color="secondary">
            Photo
          </Text>

          <img src="https://64.media.tumblr.com/65390415cadbd069148850d7cdcbedd1/tumblr_ozgyzr929N1qjoleso1_1280.jpg" />
        </div>

        <div className="board__name">
          <Text bd="500" size="0.9rem" color="secondary">
            Name
          </Text>

          <Text>Jason Ngan</Text>
        </div>

        <div className="board__bio">
          <Text bd="500" size="0.9rem" color="secondary">
            Bio
          </Text>

          <Text>I am a big fans of software developer</Text>
        </div>

        <div className="board__phone">
          <Text bd="500" size="0.9rem" color="secondary">
            Phone
          </Text>

          <Text>1234567</Text>
        </div>

        <div className="board__email">
          <Text bd="500" size="0.9rem" color="secondary">
            Email
          </Text>

          <Text>nganjason007@gmail.com</Text>
        </div>

        <div className="board__password">
          <Text bd="500" size="0.9rem" color="secondary">
            Password
          </Text>

          <Text>***********</Text>
        </div>
      </div>
    </div>
  );
}

export default Info;
