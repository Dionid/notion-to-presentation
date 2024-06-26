// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/Dionid/notion-to-presentation/libs/ntp/models"
	"github.com/Dionid/notion-to-presentation/libs/templu"
)

func PresentationPage(presentation *models.Presentation) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			templ_7745c5c3_Err = templ.JSONScript("presentation-data", presentation).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" <script>\n            window.addEventListener(\"load\", () => {\n                Tawk_API.customStyle = {\n                    visibility: {\n                        //for desktop only\n                        desktop: {\n                            position: 'tr', // bottom-right\n                            // xOffset: 15, // 15px away from right\n                            // yOffset: 40 // 40px up from bottom\n                        },\n                        // for mobile only\n                        // mobile: {\n                            // position: 'bl', // bottom-left\n                            // xOffset: 5, // 5px away from left\n                            // yOffset: 50 // 50px up from bottom\n                        // },\n                        // change settings of bubble if necessary\n                        // bubble: {\n                            // rotate: '0deg',\n                            // xOffset: -20,\n                            // yOffset: 0\n                        // }\n                    }\n                }\n            })\n        </script> <script type=\"module\" src=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(templu.PathWithVersion(ctx, "/public/widgets/presentation-config.js"))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `presentation.templ`, Line: 37, Col: 105}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></script> <div id=\"presentation-component\" class=\"w-full h-full flex flex-row-reverse\"><div class=\"w-full relative\"><div id=\"presentation-container\" class=\"w-full\" v-html=\"html\"></div><a href=\"n2p.dev\" target=\"_blank\" class=\"absolute left-5 bottom-5 text-xs \" style=\"color: #878787;\">made with n2p.dev</a></div><div id=\"presentation-config\" class=\"w-full max-w-sm h-full relative\"><label class=\"btn btn-circle swap swap-rotate absolute\" style=\"right: -60px; top: 10px;\"><input type=\"checkbox\" @click=\"toggleConfigExpanded\"> <svg class=\"swap-on w-5 h-5\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 20 20\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7.75 4H19M7.75 4a2.25 2.25 0 0 1-4.5 0m4.5 0a2.25 2.25 0 0 0-4.5 0M1 4h2.25m13.5 6H19m-2.25 0a2.25 2.25 0 0 1-4.5 0m4.5 0a2.25 2.25 0 0 0-4.5 0M1 10h11.25m-4.5 6H19M7.75 16a2.25 2.25 0 0 1-4.5 0m4.5 0a2.25 2.25 0 0 0-4.5 0M1 16h2.25\"></path></svg> <svg class=\"swap-off w-5 h-5\" aria-hidden=\"true\" xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 8 14\"><path stroke=\"currentColor\" stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M7 1 1.3 6.326a.91.91 0 0 0 0 1.348L7 13\"></path></svg></label><div class=\"flex flex-col bg-white border-r border-solid border-gray-300 shadow w-full max-w-sm h-full overflow-auto\"><form class=\"flex flex-col p-6\"><p class=\"font-semibold\">Global styles:</p><label for=\"link\" class=\"block mt-5\">Notion link:</label> <input type=\"text\" placeholder=\"Presentation link\" name=\"link\" class=\"mt-2 input input-bordered w-full\" readonly v-model=\"notionPageUrl\"><div class=\"mt-5 flex justify-between w-full content-center\"><label for=\"public\" class=\"\">Published:</label> <input type=\"checkbox\" class=\"toggle\" name=\"public\" v-model=\"public\"></div><div v-if=\"public\" class=\"join mt-2\"><input name=\"publicLink\" class=\"input input-bordered w-full join-item\" target=\"_blank\" readonly v-model=\"publicLink\"> <button type=\"button\" class=\"btn btn-primary join-item\" @click=\"copyPublicLink\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" fill=\"currentColor\" class=\"size-4\"><path d=\"M5.5 3.5A1.5 1.5 0 0 1 7 2h2.879a1.5 1.5 0 0 1 1.06.44l2.122 2.12a1.5 1.5 0 0 1 .439 1.061V9.5A1.5 1.5 0 0 1 12 11V8.621a3 3 0 0 0-.879-2.121L9 4.379A3 3 0 0 0 6.879 3.5H5.5Z\"></path> <path d=\"M4 5a1.5 1.5 0 0 0-1.5 1.5v6A1.5 1.5 0 0 0 4 14h5a1.5 1.5 0 0 0 1.5-1.5V8.621a1.5 1.5 0 0 0-.44-1.06L7.94 5.439A1.5 1.5 0 0 0 6.878 5H4Z\"></path></svg></button></div><label for=\"title\" class=\"block mt-5\">Title:</label> <input type=\"text\" placeholder=\"Presentation title\" name=\"title\" class=\"mt-2 input input-bordered w-full\" required v-model=\"title\"> <label for=\"description\" class=\"block mt-2\">Description:</label> <textarea type=\"text\" placeholder=\"Presentation description\" name=\"description\" class=\"mt-2 min-h-20 textarea textarea-bordered\" v-model=\"description\"></textarea> <label for=\"theme\" class=\"block mt-2\">Theme:</label> <select class=\"mt-2 select select-bordered w-full\" name=\"theme\" v-model=\"theme\"><option>beige</option> <option>black-contrast</option> <option>black</option> <option>blood</option> <option>dracula</option> <option>league</option> <option>moon</option> <option>night</option> <option>serif</option> <option>simple</option> <option>sky</option> <option>solarized</option> <option>white_contrast_compact_verbatim_headers</option> <option>white-contrast</option> <option>white</option></select> <label for=\"mainFont\" class=\"block mt-2\">Content font:</label> <select class=\"mt-2 select select-bordered w-full\" v-model=\"mainFont\" name=\"mainFont\"><option>Inter</option> <option>Arial</option> <option>Roboto</option> <option>Roboto Mono</option> <option>Tiny5</option> <option>Open Sans</option> <option>Montserrat</option> <option>Merriweather</option> <option>Lora</option></select> <label for=\"headingFont\" class=\"block mt-2\">Heading font:</label> <select class=\"mt-2 select select-bordered w-full\" v-model=\"headingFont\" name=\"headingFont\"><option>Inter</option> <option>Arial</option> <option>Roboto</option> <option>Roboto Mono</option> <option>Tiny5</option> <option>Open Sans</option> <option>Montserrat</option> <option>Merriweather</option> <option>Lora</option></select> <label for=\"headingTextAlign\" class=\"block mt-2\">Heading text align:</label> <select class=\"mt-2 select select-bordered w-full\" v-model=\"headingTextAlign\" name=\"headingTextAlign\"><option>left</option> <option>center</option> <option>right</option></select> <label for=\"contentTextAlign\" class=\"block mt-2\">Content text align:</label> <select class=\"mt-2 select select-bordered w-full\" v-model=\"contentTextAlign\" name=\"contentTextAlign\"><option>left</option> <option>center</option> <option>right</option></select> <label for=\"mainFontSize\" class=\"block mt-2\">Content text size:</label> <input type=\"text\" placeholder=\"Main text size\" name=\"mainFontSize\" class=\"mt-2 input input-bordered w-full\" required v-model=\"mainFontSize\"> <label for=\"headingFontWeight\" class=\"block mt-2\">Heading font weight:</label><div><input name=\"headingFontWeight\" type=\"range\" min=\"100\" max=\"900\" class=\"range\" step=\"100\" v-model=\"headingFontWeight\"><div class=\"w-full flex justify-between text-xs px-2 mt-2\"><span>100</span> <span>200</span> <span>300</span> <span>400</span> <span>500</span> <span>600</span> <span>700</span> <span>800</span> <span>900</span></div></div><label for=\"heading1Size\" class=\"block mt-2\">Heading 1 size:</label> <input type=\"text\" placeholder=\"Heading 1 text size\" name=\"heading1Size\" class=\"mt-2 input input-bordered w-full\" required v-model=\"heading1Size\"> <label for=\"heading2Size\" class=\"block mt-2\">Heading 2 size:</label> <input type=\"text\" placeholder=\"Heading 2 text size\" name=\"heading2Size\" class=\"mt-2 input input-bordered w-full\" required v-model=\"heading2Size\"> <label for=\"heading3Size\" class=\"block mt-3\">Heading 3 size:</label> <input type=\"text\" placeholder=\"Heading 3 text size\" name=\"heading3Size\" class=\"mt-2 input input-bordered w-full\" required v-model=\"heading3Size\"> <label for=\"heading4Size\" class=\"block mt-3\">Heading 4 size:</label> <input type=\"text\" placeholder=\"Heading 4 text size\" name=\"heading4Size\" class=\"mt-2 input input-bordered w-full\" required v-model=\"heading4Size\"> <label for=\"backgroundColor\" class=\"block mt-3\">Custom background color:</label> <input type=\"text\" placeholder=\"like #ffffff\" name=\"backgroundColor\" class=\"mt-2 input input-bordered w-full\" required v-model=\"backgroundColor\"> <label for=\"mainColor\" class=\"block mt-3\">Text color:</label> <input type=\"text\" placeholder=\"like #ffffff\" name=\"mainColor\" class=\"mt-2 input input-bordered w-full\" required v-model=\"mainColor\"> <label for=\"headingColor\" class=\"block mt-3\">Heading color:</label> <input type=\"text\" placeholder=\"like #ffffff\" name=\"headingColor\" class=\"mt-2 input input-bordered w-full\" required v-model=\"headingColor\"> <label for=\"custom-css\" class=\"block mt-2\">Custom CSS:</label> <textarea type=\"text\" placeholder=\"if you know what are you doing\" name=\"custom-css\" class=\"min-h-40 mt-2 textarea textarea-bordered\" v-model=\"customCss\"></textarea><div v-for=\"(value, key) in customizedSlides\" class=\"flex flex-col w-full mt-5\"><div class=\"divider\"></div><label class=\"block font-semibold\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.Raw("Slide #{{ key }}").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(":</label> <label class=\"block mt-2\">Heading text align:</label> <select class=\"mt-2 select select-bordered w-full\" v-model=\"customizedSlides[key].headingTextAlign\"><option>left</option> <option>center</option> <option>right</option></select> <label for=\"contentTextAlign\" class=\"block mt-2\">Content text align:</label> <select class=\"mt-2 select select-bordered w-full\" v-model=\"customizedSlides[key].contentTextAlign\"><option>left</option> <option>center</option> <option>right</option></select> <label for=\"mainFontSize\" class=\"block mt-2\">Content text size:</label> <input type=\"text\" placeholder=\"Main text size\" class=\"mt-2 input input-bordered w-full\" required v-model=\"customizedSlides[key].mainFontSize\"> <label for=\"heading1Size\" class=\"block mt-2\">Heading 1 size:</label> <input type=\"text\" placeholder=\"Heading 1 text size\" class=\"mt-2 input input-bordered w-full\" required v-model=\"customizedSlides[key].heading1Size\"> <label for=\"heading2Size\" class=\"block mt-2\">Heading 2 size:</label> <input type=\"text\" placeholder=\"Heading 2 text size\" class=\"mt-2 input input-bordered w-full\" required v-model=\"customizedSlides[key].heading2Size\"> <label for=\"heading3Size\" class=\"block mt-3\">Heading 3 size:</label> <input type=\"text\" placeholder=\"Heading 3 text size\" class=\"mt-2 input input-bordered w-full\" required v-model=\"customizedSlides[key].heading3Size\"> <label for=\"heading4Size\" class=\"block mt-3\">Heading 4 size:</label> <input type=\"text\" placeholder=\"Heading 4 text size\" class=\"mt-2 input input-bordered w-full\" required v-model=\"customizedSlides[key].heading4Size\"> <label for=\"headingColor\" class=\"block mt-3\">Heading color:</label> <input type=\"text\" placeholder=\"like #ffffff\" class=\"mt-2 input input-bordered w-full\" required v-model=\"customizedSlides[key].headingColor\"> <label for=\"mainColor\" class=\"block mt-3\">Text color:</label> <input type=\"text\" placeholder=\"like #ffffff\" class=\"mt-2 input input-bordered w-full\" required v-model=\"customizedSlides[key].mainColor\"> <button type=\"button\" class=\"mt-5 btn btn-ghost\" @click=\"deleteSlideStyles(key)\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.Raw("Delete slide #{{ key }} styles").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button><div class=\"divider\"></div></div><button type=\"button\" class=\"mt-5 btn btn-outlined\" @click=\"addCustomSlideStyles\">+ Custom slide styles</button><div v-if=\"error\" role=\"alert\" class=\"alert alert-error mt-5\"><svg @click=\"error = &#39;&#39;\" xmlns=\"http://www.w3.org/2000/svg\" class=\"stroke-current shrink-0 h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg> <span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.Raw("{{ error }}").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div><button type=\"button\" class=\"mt-5 btn\" :class=\"{&#39;btn-primary&#39;: changed }\" :disabled=\"!changed\" @click=\"save\">Save</button> <button type=\"button\" class=\"mt-2 btn\" @click=\"resync\">Resync with Notion</button> <button type=\"button\" class=\"mt-2 btn btn-ghost\" @click=\"goBack\">Go back</button> <button type=\"button\" class=\"mt-5 btn btn-ghost\" @click=\"toggleChatWidget\">Toggle chat widget</button></form><div v-if=\"loading\" class=\"absolute top-0 left-0 flex w-full h-full items-center justify-center bg-black bg-opacity-30 text-white text-2xl z-10\"><span class=\"loading loading-spinner loading-lg text-primary\"></span></div></div><link rel=\"stylesheet\" href=\"/public/reveal/theme/night.css\"><link rel=\"stylesheet\" v-bind:href=\"themeUrl\"><component is=\"style\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.Raw("{{ formedCss }}").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</component> <component is=\"style\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.Raw("{{ customCss }}").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</component><link rel=\"preconnect\" href=\"https://fonts.googleapis.com\"><link rel=\"preconnect\" href=\"https://fonts.gstatic.com\" crossorigin><link :href=\"mainFontGoogleLink\" rel=\"stylesheet\"><link :href=\"headingFontGoogleLink\" rel=\"stylesheet\"></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = BaseLayout(true).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
