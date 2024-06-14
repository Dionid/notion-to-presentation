export const propOrUndefined = (name, prop) => {
  return prop ? `${name}: ${prop};` : "";
};

export const formPresentationCss = (props) => {
  const {
    customizedSlides,
    mainFont,
    mainFontSize,
    headingFont,
    headingFontWeight,
    heading1Size,
    heading2Size,
    heading3Size,
    heading4Size,
    backgroundColor,
    mainColor,
    headingColor,
    headingTextAlign,
    contentTextAlign,
  } = props;

  const customSlideStyles = Object.keys(customizedSlides)
    .map((key) => {
      const value = customizedSlides[key];

      return `
.reveal .slides section:nth-child(${key}) h1, .reveal .slides section:nth-child(${key}) h2, .reveal .slides section:nth-child(${key}) h3, .reveal .slides section:nth-child(${key}) h4 {
    text-align: ${value.headingTextAlign};
    color: ${value.headingColor};
}

.reveal .slides section:nth-child(${key}) h1 {
    font-size: ${value.heading1Size}px;
}

.reveal .slides section:nth-child(${key}) h2 {
    font-size: ${value.heading2Size}px;
}

.reveal .slides section:nth-child(${key}) h3 {
    font-size: ${value.heading3Size}px;
}

.reveal .slides section:nth-child(${key}) h4 {
    font-size: ${value.heading4Size}px;
}

.reveal .slides section:nth-child(${key}) {
    color: ${value.mainColor};
}
`;
    })
    .join("\n");

  return `:root {
    --r-main-font: ${mainFont};
    --r-main-font-size: ${mainFontSize}px;
    --r-heading-font: ${headingFont};
    --r-heading-font-weight: ${headingFontWeight};
    --r-heading1-size: ${heading1Size}px;
    --r-heading2-size: ${heading2Size}px;
    --r-heading3-size: ${heading3Size}px;
    --r-heading4-size: ${heading4Size}px;
    ${propOrUndefined("--r-background-color", backgroundColor)}
    ${propOrUndefined("--r-main-font", mainColor)}
    ${propOrUndefined("--r-heading-color", headingColor)}
}

.reveal .slides h1, .reveal .slides h2, .reveal .slides h3, .reveal .slides h4 {
    text-align: ${headingTextAlign};
}

.reveal .slides .content {
    text-align: ${contentTextAlign};
}

${customSlideStyles}
`;
};

export const formFontsUrls = (name) => {
  return `https://fonts.googleapis.com/css2?family=${name.replace(
    /\s+/g,
    "+"
  )}:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&display=swap`;
};
