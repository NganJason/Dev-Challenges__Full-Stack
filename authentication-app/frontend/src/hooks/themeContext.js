import React, { useState } from "react"
import { initThemeHandler } from "../handlers/themeHandler"

export const ThemeContext = React.createContext()

const h = initThemeHandler()

export function ThemeProvider({children}) {
    const [isDarkTheme, setIsDarkTheme] = useState(h.getIsDarkTheme())

    const toggleIsDarkTheme = () => {
        setIsDarkTheme(h.setIsDarkTheme(!isDarkTheme))
    }

    const value = {toggleIsDarkTheme, isDarkTheme};

    return (
    <ThemeContext.Provider value={value}>{children}</ThemeContext.Provider>
    );
}