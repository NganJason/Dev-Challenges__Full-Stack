import React, { useRef } from "react";
import Text from "../Text/Text"
import Button from "../Button/Button";
import DragAndDrop from "../DragAndDrop/DragAndDrop";

function Upload({ onFileSubmit }) {
  const inputRef = useRef(null);

  const onImageClick = () => {
    inputRef.current.click();
  };

  return (
    <div className="upload bg-secondary">
      <Text size="1.3rem" align="center">
        Upload your image
      </Text>

      <Text size="0.8rem" align="center" mgTop="1.5" color="secondary">
        File should be JPEG, PNG ...
      </Text>

      <div className="drag-box bg-tertiary">
        <DragAndDrop onFileSubmit={onFileSubmit} />
      </div>

      <Text size="0.8rem" align="center" mgTop="1.5" color="tertiary">
        Or
      </Text>

      <Button
        width="40%"
        height="2.5rem"
        border="8px"
        mgTop="1"
        onClick={onImageClick}
      >
        <Text size="0.8rem" color="light">
          Choose a file
        </Text>
        <input
          ref={inputRef}
          className="img__upload-btn"
          type="file"
          id="photo"
          name="photo"
          accept="image/png, image/jpeg"
          onChange={(e) => {
            onFileSubmit(e.target.files[0]);
          }}
        />
      </Button>
    </div>
  );
}

export default Upload;
