import { localStorageDM } from "../dm/localStorageDM";

export const initThemeHandler = () => {
    let dm = new localStorageDM("is_dark_theme")
    let themeHandler = new ThemeHandler(dm)

    return themeHandler
}

class ThemeHandler {
    constructor(dm) {
        this.dm = dm;
        this.isDarkTheme = this.getDefaultIsDarkTheme()
    }

    setIsDarkTheme(isDark) {
        this.isDarkTheme = isDark

        this.dm.save(isDark)

        return this.isDarkTheme
    }
    
    getIsDarkTheme() {
        return this.isDarkTheme
    }

    getDefaultIsDarkTheme() {
        let defaultIsDark = this.dm.get();
        if (defaultIsDark === undefined) {
            return false
        }

        return defaultIsDark
    }
}