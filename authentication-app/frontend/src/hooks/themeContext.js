import React, { useState } from "react"

export const ThemeContext = React.createContext()

export function ThemeProvider({children}) {
    const [isDarkTheme, setIsDarkTheme] = useState(false)

    const toggleIsDarkTheme = () => {
        setIsDarkTheme(!isDarkTheme)
    }

    const value = {toggleIsDarkTheme, isDarkTheme};

    return (
    <ThemeContext.Provider value={value}>{children}</ThemeContext.Provider>
    );
}