export default class LocalStorageWorker {
  private static localStorage = new LocalStorageWorker();
  static getLocalStorage = () => LocalStorageWorker.localStorage;

  localStorageSupported: boolean;

  constructor() {
    this.localStorageSupported =
      typeof window.localStorage !== "undefined" &&
      window.localStorage !== null;
  }

  add(key: string, value: any) {
    if (!this.localStorageSupported) return false;
    localStorage.setItem(key, value);
    return true;
  }

  get(key: string): any {
    if (!this.localStorageSupported) return null;
    return localStorage.getItem(key);
  }

  getListKey(): any {
    if (!this.localStorageSupported) return [];
    const lists = [] as string[];
    for (let i = 0; i < localStorage.length; i++) {
      lists.push(localStorage.key(i) || "");
    }
    return lists;
  }

  remove(key: string) {
    if (!this.localStorageSupported) return false;
    localStorage.removeItem(key);
    return true;
  }

  clear() {
    if (!this.localStorageSupported) return false;
    localStorage.clear();
    return true;
  }
}
