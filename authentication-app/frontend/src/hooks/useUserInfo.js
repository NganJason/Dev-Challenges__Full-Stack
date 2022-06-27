import { useState } from "react";
import { initUserInfoHandler } from "../handlers/userHandler";

let userInfoHandler = initUserInfoHandler();

export const useUserInfo = () => {
    const [userInfo, setUserInfo] = useState(userInfoHandler.getUserInfo());

    const setUser = (d) => {
        setUserInfo({ ...userInfoHandler.setUserInfo(d) });
    }

    const updateUserInfo = (d) => {
        setUserInfo({ ...userInfoHandler.updateUserInfo(d) });
    }

    const fetchLatestUserInfo = (userID) => {
        setUserInfo({ ...userInfoHandler.fetchLatestUserInfo(userID) });
    }

    return { userInfo, setUser, updateUserInfo, fetchLatestUserInfo };
}