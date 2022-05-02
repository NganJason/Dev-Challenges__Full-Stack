import React from "react";

function Text({
  align,
  size,
  color,
  mgTop,
  mgBtm,
  mgLeft,
  mgRight,
  bd,
  inline,
  hover,
  cursor,
  children,
  ...props
}) {
  const styles = {
    textAlign: `${align ? align : "left"}`,
    fontSize: `${size ? `${size}` : "1rem"}`,
    fontWeight: `${bd ? `${bd}` : "500"}`,
    marginTop: `${mgTop ? `${mgTop}rem` : ""}`,
    marginBottom: `${mgBtm ? `${mgBtm}rem` : ""}`,
    marginLeft: `${mgLeft ? `${mgLeft}rem` : ""}`,
    marginRight: `${mgRight ? `${mgRight}rem` : ""}`,
  };

  const classNames = `
  text 
  ${color ? color : "primary"}
  ${inline ? "inline" : ""}
  ${hover ? "text_hover" : ""}
  ${cursor ? "text_cursor" : ""}
`;

  return (
    <p className={classNames} style={styles}>
      {children}
    </p>
  );
}

export default Text;
