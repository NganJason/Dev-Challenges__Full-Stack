import React from "react";

function Button({
  color,
  align,
  border,
  txtAlign,
  size,
  mgTop,
  mgBtm,
  mgLeft,
  mgRight,
  round,
  active,
  inline,
  height,
  width,
  children,
  ...props
}) {
  const styles = {
    marginTop: `${mgTop ? `${mgTop}rem` : ""}`,
    marginBottom: `${mgBtm ? `${mgBtm}rem` : ""}`,
    marginLeft: `${mgLeft ? `${mgLeft}rem` : ""}`,
    marginRight: `${mgRight ? `${mgRight}rem` : ""}`,
    height: `${height ? height : ""}`,
    width: `${width ? width : ""}`,
    borderRadius: `${border ? border : ""}`,
  };

  const classNames = `
        btn 
        ${color ? color : "primary"}
        ${align ? align : "center"}
        ${txtAlign ? `txt_${txtAlign}` : "txt_center"}
        ${size ? `btn_${size}` : "btn_md"}
        ${round ? "btn_round" : ""}
        ${active ? "btn_active" : ""}
        ${inline ? "btn_inline" : ""}
    `;

  return (
    <div className={classNames} style={styles} {...props}>
      {children}
    </div>
  );
}

export default Button;
