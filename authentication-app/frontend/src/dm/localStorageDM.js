export class localStorageDM {
  constructor(key) {
    this.key = key;
  }

  save(data) {
    localStorage.setItem(this.key, JSON.stringify(data));
  }

  get() {
    let data = localStorage.getItem(this.key);
    if (data === "undefined" || data === null) {
      return undefined;
    }

    return JSON.parse(data);
  }
}
