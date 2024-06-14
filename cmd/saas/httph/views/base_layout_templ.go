// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BaseLayout() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html data-theme=\"bumblebee\"><head><title>N2P.dev - Notion to presentation</title><meta name=\"title\" content=\"N2P.dev - Notion to presentation\"><meta name=\"description\" content=\"Create interactive presentations from your Notion pages\"><meta property=\"og:site_name\" content=\"N2P\"><meta property=\"og:title\" content=\"N2P.dev - Notion to presentation\"><meta property=\"og:description\" content=\"Create interactive presentations from your Notion pages\"><meta property=\"og:url\" content=\"n2p.dev\"><meta property=\"og:image\" content=\"https://n2p.dev/public/n2p_link_preview.png\"><meta property=\"og:type\" content=\"website\"><meta name=\"twitter:card\" content=\"summary_large_image\"><meta property=\"twitter:domain\" content=\"n2p.dev\"><meta property=\"twitter:url\" content=\"https://n2p.dev/\"><meta name=\"twitter:title\" content=\"Notion to presentation\"><meta name=\"twitter:description\" content=\"Create interactive presentations from your Notion pages\"><meta name=\"twitter:image\" content=\"https://n2p.dev/public/n2p_link_preview.png\"><link rel=\"stylesheet\" href=\"/public/reveal/reveal.css\"><link rel=\"stylesheet\" href=\"/public/reveal/plugin/highlight/monokai.css\"><link rel=\"stylesheet\" href=\"/public/style.css\"><link rel=\"preconnect\" href=\"https://fonts.googleapis.com\"><link rel=\"preconnect\" href=\"https://fonts.gstatic.com\" crossorigin><link href=\"https://fonts.googleapis.com/css2?family=Inter:wght@100..900&amp;family=Roboto:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900&amp;display=swap\" rel=\"stylesheet\"><link rel=\"apple-touch-icon\" sizes=\"57x57\" href=\"/public/apple-icon-57x57.png\"><link rel=\"apple-touch-icon\" sizes=\"60x60\" href=\"/public/apple-icon-60x60.png\"><link rel=\"apple-touch-icon\" sizes=\"72x72\" href=\"/public/apple-icon-72x72.png\"><link rel=\"apple-touch-icon\" sizes=\"76x76\" href=\"/public/apple-icon-76x76.png\"><link rel=\"apple-touch-icon\" sizes=\"114x114\" href=\"/public/apple-icon-114x114.png\"><link rel=\"apple-touch-icon\" sizes=\"120x120\" href=\"/public/apple-icon-120x120.png\"><link rel=\"apple-touch-icon\" sizes=\"144x144\" href=\"/public/apple-icon-144x144.png\"><link rel=\"apple-touch-icon\" sizes=\"152x152\" href=\"/public/apple-icon-152x152.png\"><link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/public/apple-icon-180x180.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"192x192\" href=\"/public/android-icon-192x192.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"32x32\" href=\"/public/favicon-32x32.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"96x96\" href=\"/public/favicon-96x96.png\"><link rel=\"icon\" type=\"image/png\" sizes=\"16x16\" href=\"/public/favicon-16x16.png\"><link rel=\"manifest\" href=\"/public/manifest.json\"><meta name=\"msapplication-TileColor\" content=\"#ffffff\"><meta name=\"msapplication-TileImage\" content=\"/public/ms-icon-144x144.png\"><meta name=\"theme-color\" content=\"#ffffff\"></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</body><script defer src=\"https://unpkg.com/htmx.org@1.9.12\" integrity=\"sha384-ujb1lZYygJmzgSwoxRggbCHcjc0rB2XoQrxeTUQyRjrOnlCoYta87iKBWq3EsdM2\" crossorigin=\"anonymous\"></script><script defer src=\"https://cdn.jsdelivr.net/npm/vue@3.4.27/dist/vue.global.min.js\"></script><script defer src=\"/public/reveal/plugin/notes/notes.js\"></script><script defer src=\"/public/reveal/plugin/highlight/highlight.js\"></script><script defer src=\"/public/reveal/reveal.js\"></script><link rel=\"stylesheet\" href=\"/public/custom.css\"></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
