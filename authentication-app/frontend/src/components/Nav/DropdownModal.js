import React from "react"
import { Link } from "react-router-dom";
import { NewService } from "../../service/service";

import Text from "../Text/Text";

function DropdownModal({show, toggleModal}) {
   const stopPropagation = (e) => {
     e.stopPropagation();
   };

   const logout = () => {
    toggleModal();

    let s = NewService()
    window.localStorage.clear();
    s.Logout().then(() => {
      const url = window.location.origin + "/auth/login";
      window.history.pushState({}, null, url);
      window.location.reload(true);
    }).catch((err) => {
      console.log(err)
    })

    
   }

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

        <div className="bg-hover" onClick={logout}>
          <span className="material-icons inline md-24 tertiary">logout</span>
          <Text size="0.8rem" color="alert">
            Logout
          </Text>
        </div>
      </div>
    );
}

export default DropdownModal