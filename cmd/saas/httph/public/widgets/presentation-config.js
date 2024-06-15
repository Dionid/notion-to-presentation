import { formPresentationCss, formFontsUrls } from "./form-css.js";

window.addEventListener("load", function () {
  const { createApp } = Vue;

  const data = JSON.parse(
    document.getElementById("presentation-data").textContent
  );

  if (!data) {
    alert("No data found");
    return;
  }

  let oldHtml = data.html;

  // # Config component
  createApp({
    mounted() {
      // # Reveal
      const revealPresentation = new Reveal({
        hash: true,
        plugins: [RevealHighlight, RevealNotes],
        embedded: true,
      });
      revealPresentation.initialize();
    },
    updated() {
      if (oldHtml !== this.html) {
        oldHtml = this.html;

        // # Reveal
        const revealPresentation = new Reveal({
          hash: true,
          plugins: [RevealHighlight, RevealNotes],
          embedded: true,
        });

        revealPresentation.initialize();
      }
    },
    data() {
      const global = data.customizations.global || {};

      return {
        // # State
        loading: false,
        error: "",
        configExpanded: true,
        changed: false,
        // # Data
        id: data.id,
        title: data.title,
        description: data.description,
        customCss: data.customCss,
        notionPageUrl: data.notionPageUrl,
        html: data.html,
        theme: data.theme,
        public: data.public,
        mainFont: global.mainFont || "Inter",
        mainFontSize: global.mainFontSize || 30,
        headingFont: global.headingFont || "Inter",
        headingFontWeight: global.headingFontWeight || "400",
        heading1Size: global.heading1Size || 152,
        heading2Size: global.heading2Size || 84,
        heading3Size: global.heading3Size || 62,
        heading4Size: global.heading4Size || 40,
        backgroundColor: global.backgroundColor,
        mainColor: global.mainColor,
        headingColor: global.headingColor,
        headingTextAlign: global.headingTextAlign || "left",
        contentTextAlign: global.contentTextAlign || "left",
        customizedSlides: data.customizations.customizedSlides || {},
      };
    },
    watch: {
      changedData: {
        handler: function (val) {
          this.changed = true;
        },
        deep: true,
      },
    },
    computed: {
      publicLink() {
        return "https://n2p.dev/public-presentations/" + this.id;
      },
      changedData() {
        return [
          this.title,
          this.description,
          this.customCss,
          this.theme,
          this.public,
          this.mainFont,
          this.headingFontWeight,
          this.mainFontSize,
          this.headingFont,
          this.heading1Size,
          this.heading2Size,
          this.heading3Size,
          this.heading4Size,
          this.backgroundColor,
          this.mainColor,
          this.headingColor,
          this.headingTextAlign,
          this.contentTextAlign,
          this.customizedSlides,
          this.html,
        ];
      },
      themeUrl() {
        return `/public/reveal/theme/${this.theme}.css`;
      },
      formedCss() {
        return formPresentationCss(this);
      },
      mainFontGoogleLink() {
        return formFontsUrls(this.mainFont);
      },
      headingFontGoogleLink() {
        return formFontsUrls(this.headingFont);
      },
    },
    methods: {
      toggleChatWidget() {
        Tawk_API.toggleVisibility();
      },
      copyPublicLink() {
        navigator.clipboard.writeText(this.publicLink);
        alert("Copied to clipboard");
      },
      toggleConfigExpanded() {
        const el = document.getElementById("presentation-config");

        if (el.style.width && el.style.width[0] === "0") {
          el.style.width = "100%";
        } else {
          el.style.width = "0px";
        }
      },
      addCustomSlideStyles() {
        const slide = parseInt(
          prompt("Enter slide number to add custom styles")
        );
        if (!slide || Number.isNaN(slide) || slide <= 0) {
          alert("Invalid slide number");
          return;
        }

        if (this.customizedSlides[slide]) {
          alert("Custom styles already added for this slide");
          return;
        }

        this.customizedSlides[slide] = {
          headingTextAlign: this.headingTextAlign,
          contentTextAlign: this.contentTextAlign,
          mainFontSize: this.mainFontSize,
          heading1Size: this.heading1Size,
          heading2Size: this.heading2Size,
          heading3Size: this.heading3Size,
          heading4Size: this.heading4Size,
          mainColor: this.mainColor,
          headingColor: this.headingColor,
        };
      },
      deleteSlideStyles(slide) {
        const resp = confirm(
          `Are you sure you want to delete slide #${slide} styles?`
        );
        if (!resp) {
          return;
        }

        delete this.customizedSlides[slide];
      },
      goBack() {
        if (this.changed) {
          const resp = confirm(
            "You have unsaved changes. Are you sure you want to leave?"
          );
          if (!resp) {
            return;
          }
        }
        window.location.href = "/app";
      },
      async resync() {
        this.loading = true;

        const response = await fetch(
          "/app/presentation/" + this.id + "/resync",
          {
            method: "GET",
            headers: {
              "Content-Type": "application/json",
            },
          }
        );

        this.loading = false;

        const respJson = await response.json();

        if (response.ok) {
          this.html = respJson.result;
        } else {
          this.error = respJson.message;
        }
      },
      async save() {
        this.error = "";

        const data = {
          title: this.title,
          description: this.description,
          custom_css: this.customCss,
          theme: this.theme,
          public: this.public,
          html: this.html,
          customizations: {
            customizedSlides: this.customizedSlides,
            global: {
              mainFont: this.mainFont,
              mainFontSize: this.mainFontSize,
              headingFont: this.headingFont,
              headingFontWeight: this.headingFontWeight,
              heading1Size: this.heading1Size,
              heading2Size: this.heading2Size,
              heading3Size: this.heading3Size,
              heading4Size: this.heading4Size,
              backgroundColor: this.backgroundColor,
              mainColor: this.mainColor,
              headingColor: this.headingColor,
              headingTextAlign: this.headingTextAlign,
              contentTextAlign: this.contentTextAlign,
            },
          },
        };

        const response = await fetch(
          "/api/collections/presentation/records/" + this.id,
          {
            method: "PATCH",
            headers: {
              "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
          }
        );

        const respJson = await response.json();

        if (response.ok) {
          alert("Saved");
        } else {
          this.error = respJson.message;
        }

        this.changed = false;
      },
    },
  }).mount("#presentation-component");
});
