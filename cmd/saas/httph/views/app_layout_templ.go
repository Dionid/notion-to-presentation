// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package views

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/Dionid/notion-to-presentation/cmd/saas/httph/views/partials"
)

func AppLayout() templ.Component {
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
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col h-full\"><header class=\"w-full flex items-center justify-center max-w-6xl mx-auto\"><div class=\"w-full m-6 gap-2 flex\"><a href=\"/app\" class=\"h-12 flex justify-center items-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = partials.Logo().Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a><script>\n                        window.addEventListener(\"load\", () => {\n                            if (window.location.pathname === \"/app\") {\n                                const el = document.getElementById(\"presentation-link\").classList\n                                el.add(\"btn-primary\")\n                                el.remove(\"btn-ghost\")\n                            } else if (window.location.pathname === \"/app/my-profile\") {\n                                const el = document.getElementById(\"my-profile-link\").classList\n                                el.add(\"btn-primary\")\n                                el.remove(\"btn-ghost\")\n                            }\n                        })\n                    </script><div class=\"ml-auto\"><a href=\"/app\" class=\"btn btn-ghost\" id=\"presentation-link\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" fill=\"currentColor\" class=\"size-4\"><path fill-rule=\"evenodd\" d=\"M4 2a1.5 1.5 0 0 0-1.5 1.5v9A1.5 1.5 0 0 0 4 14h8a1.5 1.5 0 0 0 1.5-1.5V6.621a1.5 1.5 0 0 0-.44-1.06L9.94 2.439A1.5 1.5 0 0 0 8.878 2H4Zm1 5.75A.75.75 0 0 1 5.75 7h4.5a.75.75 0 0 1 0 1.5h-4.5A.75.75 0 0 1 5 7.75Zm0 3a.75.75 0 0 1 .75-.75h4.5a.75.75 0 0 1 0 1.5h-4.5a.75.75 0 0 1-.75-.75Z\" clip-rule=\"evenodd\"></path></svg> Presentations</a> <a href=\"/app/my-profile\" class=\"btn btn-ghost\" id=\"my-profile-link\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" fill=\"currentColor\" class=\"size-4\"><path fill-rule=\"evenodd\" d=\"M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0Zm-5-2a2 2 0 1 1-4 0 2 2 0 0 1 4 0ZM8 9c-1.825 0-3.422.977-4.295 2.437A5.49 5.49 0 0 0 8 13.5a5.49 5.49 0 0 0 4.294-2.063A4.997 4.997 0 0 0 8 9Z\" clip-rule=\"evenodd\"></path></svg> My profile</a><script>\n                            window.window.addEventListener(\"load\", function () {\n                                document.getElementById(\"logout-button\").addEventListener(\"click\", async () => {\n                                    await window.cookieStore.delete(\"pb_auth\")\n                                    window.location.href = \"\"\n                                })\n                            })\n                        </script><button class=\"btn btn-ghost\" id=\"logout-button\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 16 16\" fill=\"currentColor\" class=\"size-4\"><path fill-rule=\"evenodd\" d=\"M14 4.75A2.75 2.75 0 0 0 11.25 2h-3A2.75 2.75 0 0 0 5.5 4.75v.5a.75.75 0 0 0 1.5 0v-.5c0-.69.56-1.25 1.25-1.25h3c.69 0 1.25.56 1.25 1.25v6.5c0 .69-.56 1.25-1.25 1.25h-3c-.69 0-1.25-.56-1.25-1.25v-.5a.75.75 0 0 0-1.5 0v.5A2.75 2.75 0 0 0 8.25 14h3A2.75 2.75 0 0 0 14 11.25v-6.5Zm-9.47.47a.75.75 0 0 0-1.06 0L1.22 7.47a.75.75 0 0 0 0 1.06l2.25 2.25a.75.75 0 1 0 1.06-1.06l-.97-.97h7.19a.75.75 0 0 0 0-1.5H3.56l.97-.97a.75.75 0 0 0 0-1.06Z\" clip-rule=\"evenodd\"></path></svg> Logout</button></div></div></header><div class=\"w-full mx-auto flex max-w-6xl p-6 h-full\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><footer class=\"w-full flex items-center justify-between max-w-6xl mx-auto p-6\"><div>Made by <a href=\"https://github.com/Dionid\" target=\"_blank\">David Shekunts</a> with ♥️</div><div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ.Raw("@2024").Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></footer></div>")
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
