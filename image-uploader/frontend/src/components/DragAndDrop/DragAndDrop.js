import React, { useState } from "react";
import Text from "../Text/Text";
import DragIcon from "../../images/dragIcon.svg";

function DragAndDrop({ onFileSubmit }) {
  const [dragging, setDragging] = useState(false);
  const [dragCounter, setDragCounter] = useState(0);

  const handleDragIn = (e) => {
    e.preventDefault();
    e.stopPropagation();

    setDragCounter(dragCounter + 1);

    if (e.dataTransfer.items && e.dataTransfer.items.length > 0) {
      setDragging(true);
    }
  };

  const handleDragOut = (e) => {
    e.preventDefault();
    e.stopPropagation();

    setDragCounter(dragCounter - 1);
    if (dragCounter === 0) {
      setDragging(false);
    }
  };

  const handleDrop = (e) => {
    e.preventDefault();
    e.stopPropagation();

    setDragCounter(0);
    setDragging(false);

    if (e.dataTransfer.files && e.dataTransfer.files.length > 0) {
      onFileSubmit(e.dataTransfer.files[0]);
      e.dataTransfer.clearData();
    }
  };

  const handleDragOver = (e) => {
    e.preventDefault();
    e.stopPropagation();
  };

  return (
    <div
      className="drag-and-drop"
      onDragOver={handleDragOver}
      onDragEnter={handleDragIn}
      onDragLeave={handleDragOut}
      onDrop={handleDrop}
    >
      <div className="drag-and-drop__bg">
        <img src={DragIcon}/>

        <Text size="0.8rem" align="center" color="tertiary">
          Drag & Drop your image here
        </Text>
      </div>
    </div>
  );
}

export default DragAndDrop;
