import React, { useState } from "react";

import Text from "../Text/Text";
import Button from "../Button/Button"
import { Link, useNavigate } from "react-router-dom";

function EditInfo({userData, setUserData}) {
  const [clonedData, setClonedData] = useState({...userData})
  let navigate = useNavigate();

  const onInputChange = (e) => {
    let id = e.target.id

    if (!(id in clonedData)) {
      return
    }

    let newData = {...clonedData}
    newData[id] = e.target.value
    setClonedData(newData)
  }

  const onFormSubmit = () => {
    setUserData(clonedData)
    routeChange()
  }

  const routeChange = () => {
    let path = "/";
    navigate(path);
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
          <img src={clonedData["image"]} />

          <Text size="0.8rem" bd="500" tertiary>
            Name
          </Text>
          <input
            id="name"
            placeholder="Enter your name..."
            type="text"
            value={clonedData["name"]}
            onChange={onInputChange}
          />

          <Text size="0.8rem" bd="500" tertiary>
            Bio
          </Text>
          <input
            id="bio"
            placeholder="Enter your bio..."
            type="text"
            value={clonedData["bio"]}
            onChange={onInputChange}
          />

          <Text size="0.8rem" bd="500" tertiary>
            Phone
          </Text>
          <input
            id="phone"
            placeholder="Enter your phone..."
            type="text"
            value={clonedData["phone"]}
            onChange={onInputChange}
          />

          <Text size="0.8rem" bd="500" tertiary>
            Email
          </Text>
          <input
            id="email"
            placeholder="Enter your email..."
            type="text"
            value={clonedData["email"]}
            onChange={onInputChange}
          />

          <Text size="0.8rem" bd="500" tertiary>
            Password
          </Text>
          <input
            id="password"
            placeholder="Enter your password..."
            type="text"
            value={clonedData["password"]}
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
