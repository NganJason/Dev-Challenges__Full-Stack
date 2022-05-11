import React from "react"
import { Link } from "react-router-dom";

import Text from "../Text/Text";

function DropdownModal({show}) {
   const stopPropagation = (e) => {
     e.stopPropagation();
   };

    return (
      <div
        className={`dropdown-modal bg-primary ${show ? "show" : ""}`}
        onClick={stopPropagation}
      >
        <div className="bg-hover">
          <span className="material-icons inline md-24 secondary">
            account_circle
          </span>
          <Link to="/" style={{ textDecoration: "none" }}>
            <Text size="0.8rem" color="tertiary" inline>
              My Profile
            </Text>
          </Link>
        </div>

        <div className="bg-hover">
          <span className="material-icons inline md-24 secondary">people</span>
          <Text size="0.8rem" color="tertiary">
            Group Chat
          </Text>
        </div>

        <div className="bg-hover">
          <span className="material-icons inline md-24 tertiary">logout</span>
          <Text size="0.8rem" color="alert">
            Logout
          </Text>
        </div>
      </div>
    );
}

export default DropdownModal