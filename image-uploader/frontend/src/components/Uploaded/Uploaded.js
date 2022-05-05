import React from "react";
import Button from "../Button/Button";
import Text from "../Text/Text";

function Uploaded({url}) {
  const copyToClipboard = () => {
    navigator.clipboard.writeText(url)
  }

  const shortenUrlLink = () => {
    if (url.length > 30) {
      return `${url.slice(0,30)}....`
    }

    return url
  }
  return (
    <div className="uploaded bg-secondary">
      <span class="material-icons md-48 primary">check_circle</span>

      <Text size="1.5rem" align="center" primary>
        Uploaded Successfully
      </Text>

      <img className="uploaded__img" src={url} />

      <div className="uploaded__link bg-tertiary">
        <Text size="0.7rem">
          {shortenUrlLink()}
        </Text>

        <Button
          border="8px"
          width="25%"
          height="90%"
          align="right"
          onClick={copyToClipboard}
        >
          <Text size="0.7rem" color="light">
            Copy Link
          </Text>
        </Button>
      </div>
    </div>
  );
}

export default Uploaded;
