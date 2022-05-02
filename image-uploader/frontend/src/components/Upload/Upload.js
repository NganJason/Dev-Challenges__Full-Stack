import React from "react";
import Text from "../Text/Text"
import DragIcon from "../../images/dragIcon.svg"
import Button from "../Button/Button";

function Upload() {
  return (
    <div className="upload bg-secondary">
      <Text size="1.3rem" align="center">
        Upload your image
      </Text>

      <Text size="0.8rem" align="center" mgTop="1.5" color="secondary">
        File should be JPEG, PNG ...
      </Text>

      <div className="drag-box bg-tertiary">
        <img src={DragIcon} />
        <Text size="0.8rem" align="center" color="tertiary">
          Drag & Drop your image here
        </Text>
      </div>

      <Text size="0.8rem" align="center" mgTop="1.5" color="tertiary">
        Or
      </Text>

      <Button width="40%" height="2.5rem" border="8px" mgTop="1">
        Choose a file
      </Button>
    </div>
  );
}

export default Upload;
