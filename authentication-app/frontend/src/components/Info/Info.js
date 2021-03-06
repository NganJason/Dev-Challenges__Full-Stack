import React from "react";
import { useNavigate } from "react-router-dom";

import Text from "../Text/Text";
import Button from "../Button/Button"

function Info({userInfo}) {
  let navigate = useNavigate(); 

  const routeChange = () => {
    let path = `edit-info`;
    navigate(path);
  };

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

          <Button color="secondary" border="12px" onClick={routeChange}>
            <Text bd="500" size="1rem" color="secondary">
              Edit
            </Text>
          </Button>
        </div>

        <div className="board__name">
          <Text bd="500" size="0.9rem" color="secondary">
            Username
          </Text>

          <div className="field__content">
            <Text>{userInfo["username"] ? userInfo["username"] : "NA"}</Text>
          </div>
        </div>

        <div className="board__bio">
          <Text bd="500" size="0.9rem" color="secondary">
            Bio
          </Text>

          <div className="field__content">
            <Text>{userInfo["bio"] ? userInfo["bio"] : "NA"}</Text>
          </div>
        </div>

        <div className="board__phone">
          <Text bd="500" size="0.9rem" color="secondary">
            Phone
          </Text>

          <div className="field__content">
            <Text>{userInfo["phone"] ? userInfo["phone"] : "NA"}</Text>
          </div>
        </div>

        <div className="board__email">
          <Text bd="500" size="0.9rem" color="secondary">
            Email
          </Text>

          <div className="field__content">
            <Text>{userInfo["email"] ? userInfo["email"] : "NA"}</Text>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Info;
