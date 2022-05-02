import React from "react";
import DragIcon from "../../images/dragIcon.svg";
import Button from "../Button/Button";
import Text from "../Text/Text";

function Uploaded() {
  const copyToClipboard = () => {
    navigator.clipboard.writeText("Hello world")
  }
  return (
    <div className="uploaded bg-secondary">
      <span class="material-icons md-48 primary">check_circle</span>

      <Text size="1.5rem" align="center" primary>
        Uploaded Successfully
      </Text>

      <img
        className="uploaded__img"
        src={
          "https://images.pexels.com/photos/2662116/pexels-photo-2662116.jpeg?cs=srgb&dl=pexels-jaime-reimer-2662116.jpg&fm=jpg"
        }
      />

      <div className="uploaded__link bg-tertiary">
        <Text size="0.7rem">http://google.com</Text>

        <Button border="8px" width="25%" height="90%" align="right" onClick={copyToClipboard}>
          <Text size="0.7rem" color="light">
            Copy Link
          </Text>
        </Button>
      </div>
    </div>
  );
}

export default Uploaded;
