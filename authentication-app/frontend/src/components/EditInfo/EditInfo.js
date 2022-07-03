import React, { useState, useEffect } from "react";

import Text from "../Text/Text";
import Button from "../Button/Button"
import { Link } from "react-router-dom";

function EditInfo({
  userInfo, 
  updateUserInfo,
  fetchLatestUserInfo,
}) {
  const [userInfoInput, setUserInfoInput] = useState({...userInfo})

  useEffect(() => {
    fetchLatestUserInfo(userInfo.user_id);
  }, [fetchLatestUserInfo, userInfo.user_id]);

  const onInputChange = (e) => {
    let id = e.target.id

    if (!(id in userInfoInput)) {
      return
    }

    let newData = {...userInfoInput}
    newData[id] = e.target.value
    setUserInfoInput(newData)
  }

  const onFormSubmit = () => {
    updateUserInfo(userInfoInput).then(() => {
      routeChange();
    }).catch((err) => {
      console.log(err)
    });
    
  }

  const routeChange = () => {
    const url = "/";
    window.history.pushState({}, null, url);
    window.location.reload(true);
  };

  return (
    <div className="edit-info">
      <div className="edit__link">
        <span class="material-icons md-18 inline link">arrow_back_ios</span>
        <Link to="/">Back</Link>
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
          <Text size="0.8rem" bd="500" tertiary>
            Username
          </Text>
          <input
            id="username"
            placeholder="Enter your name..."
            type="text"
            value={userInfoInput["username"]}
            onChange={onInputChange}
          />

          <Text size="0.8rem" bd="500" tertiary>
            Bio
          </Text>
          <input
            id="bio"
            placeholder="Enter your bio..."
            type="text"
            value={userInfoInput["bio"]}
            onChange={onInputChange}
          />

          <Text size="0.8rem" bd="500" tertiary>
            Phone
          </Text>
          <input
            id="phone"
            placeholder="Enter your phone..."
            type="text"
            value={userInfoInput["phone"]}
            onChange={onInputChange}
          />

          <Text size="0.8rem" bd="500" tertiary>
            Email
          </Text>
          <input
            id="email"
            placeholder="Enter your email..."
            type="text"
            value={userInfoInput["email"]}
            onChange={onInputChange}
          />

          <Button
            border="8px"
            width="5rem"
            height="2.3rem"
            mgTop="1.5"
            onClick={onFormSubmit}
          >
            Save
          </Button>
        </div>
      </div>
    </div>
  );
}

export default EditInfo;
