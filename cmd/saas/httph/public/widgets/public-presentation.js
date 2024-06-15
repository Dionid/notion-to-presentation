import { formPresentationCss, formFontsUrls } from "./form-css.js";

window.addEventListener("load", function () {
  const component = document.querySelector("#public-presentation-component");

  if (!component) {
    return;
  }

  const data = JSON.parse(
    document.getElementById("presentation-data").textContent
  );

  if (!data) {
    alert("No data found");
    return;
  }

  const styleBlock = document.createElement("style");
  styleBlock.innerHTML = formPresentationCss({
    ...data.customizations.global,
    customizedSlides: data.customizations.customizedSlides,
  });

  component.appendChild(styleBlock);

  if (
    data.customizations.global.mainFont ===
    data.customizations.global.headingFont
  ) {
    const fontUrl = formFontsUrls(data.customizations.global.mainFont);

    const linkBlock = document.createElement("link");
    linkBlock.rel = "stylesheet";
    linkBlock.href = fontUrl;

    component.appendChild(linkBlock);
  } else {
    const mainFontUrl = formFontsUrls(data.customizations.global.mainFont);
    const headingFontUrl = formFontsUrls(data.customizations.global.mainFont);

    const mainFrontLinkBlock = document.createElement("link");
    mainFrontLinkBlock.rel = "stylesheet";
    mainFrontLinkBlock.href = mainFontUrl;

    const headingFrontLinkBlock = document.createElement("link");
    headingFrontLinkBlock.rel = "stylesheet";
    headingFrontLinkBlock.href = headingFontUrl;

    component.appendChild(mainFrontLinkBlock);
    component.appendChild(headingFrontLinkBlock);
  }
});
