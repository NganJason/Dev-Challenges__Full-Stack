import { defaultUserInfo } from "../data/data";
import { localStorageDM } from "../dm/localStorageDM";
import { NewService } from "../service/service";

export const initUserInfoHandler = () => {
    let dm = new localStorageDM("user_info")
    let service = NewService()
    let userInfoHandler = new UserInfo(dm, service);

    return userInfoHandler;
}

class UserInfo {
  constructor(dm, userService) {
    this.dm = dm;
    this.service = userService;
    this.userInfo = this.getDefaultUserInfo();
  }

  getUserInfo() {
    return this.userInfo;
  }

  setUserInfo(d) {
    this.userInfo = d;

    this.dm.save(d);
    return this.userInfo;
  }

  updateUserInfo(d) {
    let closure = this 

    this.service
      .UpdateUserInfo(d)
      .then(function () {
        closure.setUserInfo(d);
      })
      .catch(function (error) {
        console.log(error);
      });

    return this.userInfo;
  }

  fetchLatestUserInfo(userID) {
    if (userID === null || userID === undefined) {
      return this.userInfo;
    }

    let closure = this
    this.service
      .GetUserInfo(userID)
      .then(function (resp) {
        closure.setUserInfo(resp);
      })
      .catch(function (err) {
        console.log(err);
      });

    return this.userInfo;
  }

  getDefaultUserInfo() {
    let userInfo = this.dm.get();
    if (userInfo === undefined) {
      return defaultUserInfo;
    }

    return userInfo;
  }
}