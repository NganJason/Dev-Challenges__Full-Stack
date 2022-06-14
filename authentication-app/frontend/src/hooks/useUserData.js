import { useState } from "react";
import { initUserDataHandler } from "../handlers/userHandler";

let userDataHandler = initUserDataHandler()

export const useUserData = () => {
    const [userData, setUserData] = useState(userDataHandler.getUserData());

    const editUserData = (d) => {
        setUserData({...userDataHandler.setUserData(d)})
    }

    return { userData, editUserData };
}