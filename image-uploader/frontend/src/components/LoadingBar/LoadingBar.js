import React from "react";
import Text from "../Text/Text";

function LoadingBar({
  width,
  height,
  loading,
  mgTop,
  mgBtm,
  mgLeft,
  mgRight,
}) {
  const LoadingBarStyles = {
    width: `${width ? `${width}` : "200px"}`,
    marginTop: `${mgTop ? `${mgTop}rem` : ""}`,
    marginBottom: `${mgBtm ? `${mgBtm}rem` : ""}`,
    marginLeft: `${mgLeft ? `${mgLeft}rem` : ""}`,
    marginRight: `${mgRight ? `${mgRight}rem` : ""}`,
  };

  const LoadingStyles = {
    width: `${loading ? `${loading}` : "30%"}`,
    height: `${height ? `${height}` : "8px"}`,
  };

  return (
    <div style={LoadingBarStyles} className="loading-bar">
      <div className="bar">
        <div style={LoadingStyles} className="loading"></div>
      </div>
    </div>
  );
}

export default LoadingBar;
