import {userData} from "../data/data"

export const initUserDataHandler = () => {
    console.log("init new handler")
    let userDataHandler = new UserData();

    return userDataHandler
}

class UserData {
    constructor() {
        this.userData = this.getDefaultUserData()
    }

    getUserData() {
        return this.userData
    }

    setUserData(d) {
        this.userData = d
        return this.userData
    }

    getDefaultUserData() {
        return userData
    }
}