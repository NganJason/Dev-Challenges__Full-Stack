import { useState } from "react";
import { initUserInfoHandler } from "../handlers/userHandler";

let userInfoHandler = initUserInfoHandler();

export const useUserInfo = () => {
    const [userInfo, setUserInfo] = useState(userInfoHandler.getUserInfo());

    const setUser = (d) => {
        setUserInfo({ ...userInfoHandler.setUserInfo(d) });
    }

    const updateUserInfo = async (d) => {
        return userInfoHandler.updateUserInfo(d).then((res) => {
            setUserInfo({...res})
        }).catch((err) => {
            throw err
        })        
    }

    const fetchLatestUserInfo = (userID) => {
        setUserInfo({ ...userInfoHandler.fetchLatestUserInfo(userID) });
    }

    return { userInfo, setUser, updateUserInfo, fetchLatestUserInfo };
}