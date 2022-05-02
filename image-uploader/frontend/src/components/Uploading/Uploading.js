import React from "react"
import LoadingBar from "../LoadingBar/LoadingBar";
import Text from "../Text/Text"

function Uploading() {
  return (
    <div className="uploading bg-secondary">
      <Text size="1.2rem" align="left" primary>
        Uploading...
      </Text>

      <LoadingBar width="90%" mgTop="1.5"/>
    </div>
  );
}

export default Uploading;
