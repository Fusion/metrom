// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "strconv"
import "metrom/models"

func MainBody(prefs *models.Preferences) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div id=\"app w-screen\"><div class=\"navbar bg-base-300\"><div class=\"flex-1 px-2 lg:flex-none\"><form id=\"subject-form\" class=\"\" hx-post=\"/hops\" hx-on::after-request=\"if(event.detail.successful) this.reset()\"><span id=\"actioncontainer\" hx-on::after-swap=\"sensibleResetPart2()\"><input class=\"input input-bordered\" type=\"text\" name=\"subject\" id=\"subject\" placeholder=\"subject\" autofocus> <button class=\"btn\" type=\"submit\" id=\"subject-submit\">Start</button></span></form></div><div class=\"flex flex-1 justify-end px-2\"><div tabindex=\"0\" role=\"button\" class=\"btn mx-2\" hx-get=\"/help\" hx-trigger=\"click\" hx-target=\"#helpwrapper\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"1em\" height=\"1em\" viewBox=\"0 0 24 24\"><path fill=\"currentColor\" d=\"M10 19h3v3h-3zm2-17c5.35.22 7.68 5.62 4.5 9.67c-.83 1-2.17 1.66-2.83 2.5C13 15 13 16 13 17h-3c0-1.67 0-3.08.67-4.08c.66-1 2-1.59 2.83-2.25C15.92 8.43 15.32 5.26 12 5a3 3 0 0 0-3 3H6a6 6 0 0 1 6-6\"></path></svg></div><div class=\"dropdown dropdown-end space-y-6\"><div tabindex=\"0\" role=\"button\" class=\"btn\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"1em\" height=\"1em\" viewBox=\"0 0 512 512\"><path fill=\"currentColor\" d=\"M495.9 166.6c3.2 8.7.5 18.4-6.4 24.6l-43.3 39.4c1.1 8.3 1.7 16.8 1.7 25.4s-.6 17.1-1.7 25.4l43.3 39.4c6.9 6.2 9.6 15.9 6.4 24.6c-4.4 11.9-9.7 23.3-15.8 34.3l-4.7 8.1c-6.6 11-14 21.4-22.1 31.2c-5.9 7.2-15.7 9.6-24.5 6.8l-55.7-17.7c-13.4 10.3-28.2 18.9-44 25.4l-12.5 57.1c-2 9.1-9 16.3-18.2 17.8c-13.8 2.3-28 3.5-42.5 3.5s-28.7-1.2-42.5-3.5c-9.2-1.5-16.2-8.7-18.2-17.8l-12.5-57.1c-15.8-6.5-30.6-15.1-44-25.4l-55.6 17.8c-8.8 2.8-18.6.3-24.5-6.8c-8.1-9.8-15.5-20.2-22.1-31.2l-4.7-8.1c-6.1-11-11.4-22.4-15.8-34.3c-3.2-8.7-.5-18.4 6.4-24.6l43.3-39.4c-1.1-8.4-1.7-16.9-1.7-25.5s.6-17.1 1.7-25.4l-43.3-39.4c-6.9-6.2-9.6-15.9-6.4-24.6c4.4-11.9 9.7-23.3 15.8-34.3l4.7-8.1c6.6-11 14-21.4 22.1-31.2c5.9-7.2 15.7-9.6 24.5-6.8l55.7 17.7c13.4-10.3 28.2-18.9 44-25.4l12.5-57.1c2-9.1 9-16.3 18.2-17.8C227.3 1.2 241.5 0 256 0s28.7 1.2 42.5 3.5c9.2 1.5 16.2 8.7 18.2 17.8l12.5 57.1c15.8 6.5 30.6 15.1 44 25.4l55.7-17.7c8.8-2.8 18.6-.3 24.5 6.8c8.1 9.8 15.5 20.2 22.1 31.2l4.7 8.1c6.1 11 11.4 22.4 15.8 34.3zM256 336a80 80 0 1 0 0-160a80 80 0 1 0 0 160\"></path></svg></div><ul tabindex=\"0\" id=\"preferences\" class=\"dropdown-content bg-base-100 rounded-box z-50 mt-4 w-72 p-2 shadow\" hx-get=\"/preferences\" hx-trigger=\"load\" hx-target=\"#preferences\"></ul></div></div></div><div class=\"flex flex-col items-center justify-center\"><div class=\"overflow-x-auto w-screen\"><table class=\"table table-pin-rows w-screen bg-transparent border-collapse\"><!-- head --><thead><tr><th class=\"px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap\">Hop</th><th class=\"px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap\">Host</th><th class=\"px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap\">Loss</th><th class=\"px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap\">Latency</th><th class=\"px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap\">Min</th><th class=\"px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap\">Max</th><th class=\"px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap\">Jitter</th><th class=\"px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap\">Min</th><th class=\"px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap\">Max</th></tr></thead> <tbody id=\"hop-list\" class=\"divide-y divide-gray-100\" hx-target=\"#hop-list\" hx-get=\"/hops\" hx-trigger=\"every 2s\" hx-on::trigger=\"console.log(event)\"></tbody></table></div><!-- --></div></div><div id=\"modalwrapper\" hx-on::after-swap=\"sensibleReset()\"></div><div id=\"helpwrapper\"></div><script>\n    function sensibleReset(e) {\n      htmx.ajax('GET', '/hops/resetsearch', '#actioncontainer')\n    }\n\n    // Wait until the server-side is not busy canceling anymore\n    function sensibleResetPart2(e) {\n      if (document.getElementById('subject-submit').disabled) {\n        setTimeout(() => {\n          htmx.ajax('GET', '/hops/resetsearch', '#actioncontainer')\n        }, 1000)\n      }\n    }\n\n    function themeUpdateFrontend(theme) {\n      document.querySelector('html').setAttribute('data-theme', theme);\n    }\n\n    function themeUpdateBackend(theme) {\n      var body = {\n        theme: theme\n      };\n      htmx.ajax(\n        'POST',\n        '/savetheme',\n        {\n          values: {\n            theme: theme\n          }\n        })\n    }\n\n    function themeSelect(e) {\n      var cb = document.getElementById('themeselector');\n      var theme = 'light';\n      if(cb.checked != false) {\n        theme = 'dark';\n      }\n      themeUpdateFrontend(theme);\n      themeUpdateBackend(theme);\n    }\n\n    htmx.onLoad(function(elt) {\n      for (const pref of ['maxhops','timeout','probes','jitter']) {\n        const output = document.getElementById(pref + \"feedback\");\n        const input = document.getElementById(pref + \"slider\");\n        output.textContent = input.value; \n        input.addEventListener(\"input\", function() {\n          output.textContent = this.value;\n        });\n      }\n    });\n    //htmx.logAll();\n  </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if prefs.Theme == "dark" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n      themeUpdateFrontend(\"dark\");\n    </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\n      themeUpdateFrontend(\"light\");\n    </script>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func HopTable(Hops []models.FrontendHop) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		for _, val := range Hops {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<tr><td class=\"border-t-0 px-4 align-middle text-xs font-medium whitespace-nowrap p-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(val.Hop)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 162, Col: 110}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"border-t-0 px-4 align-middle text-xs font-medium whitespace-nowrap p-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(val.Host)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 163, Col: 111}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</td><td class=\"border-t-0 px-4 align-middle text-xs font-medium whitespace-nowrap p-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(val.Loss)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 164, Col: 111}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("%</td><td class=\"border-t-0 px-4 align-middle text-xs font-medium whitespace-nowrap p-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(val.LatencyAvg)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 165, Col: 117}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("ms</td><td class=\"border-t-0 px-4 align-middle text-xs font-medium whitespace-nowrap p-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(val.LatencyMin)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 166, Col: 117}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("ms</td><td class=\"border-t-0 px-4 align-middle text-xs font-medium whitespace-nowrap p-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(val.LatencyMax)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 167, Col: 117}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("ms</td><td class=\"border-t-0 px-4 align-middle text-xs font-medium whitespace-nowrap p-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 string
			templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(val.JitterAvg)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 168, Col: 116}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("ms</td><td class=\"border-t-0 px-4 align-middle text-xs font-medium whitespace-nowrap p-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(val.JitterMin)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 169, Col: 116}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("ms</td><td class=\"border-t-0 px-4 align-middle text-xs font-medium whitespace-nowrap p-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(val.JitterMax)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 170, Col: 116}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("ms</td></tr>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func Preferences(prefs *models.Preferences) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var12 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var12 == nil {
			templ_7745c5c3_Var12 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<li><form><label class=\"flex cursor-pointer\"><span class=\"label-text w-full\">Theme</span> <svg xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><circle cx=\"12\" cy=\"12\" r=\"5\"></circle> <path d=\"M12 1v2M12 21v2M4.2 4.2l1.4 1.4M18.4 18.4l1.4 1.4M1 12h2M21 12h2M4.2 19.8l1.4-1.4M18.4 5.6l1.4-1.4\"></path></svg> <input type=\"checkbox\" id=\"themeselector\" class=\"toggle theme-controller\" value=\"synthwave\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if prefs.Theme == "dark" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" onclick=\"themeSelect()\"> <svg xmlns=\"http://www.w3.org/2000/svg\" width=\"20\" height=\"20\" viewBox=\"0 0 24 24\" fill=\"none\" stroke=\"currentColor\" stroke-width=\"2\" stroke-linecap=\"round\" stroke-linejoin=\"round\"><path d=\"M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z\"></path></svg></label></form></li><li><form><label class=\"flex label cursor-pointer\"><span class=\"label-text w-full\">Hostnames</span> <input type=\"checkbox\" id=\"cb-resolve\" name=\"cb-resolve\" class=\"checkbox\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if prefs.Resolve {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" checked")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-post=\"/hops/toggleresolve\" hx-trigger=\"change\"></label></form></li><div class=\"collapse bg-base-200\"><input type=\"checkbox\"><div class=\"collapse-title font-medium\">Advanced</div><div class=\"collapse-content\"><li><form><label class=\"flex cursor-pointer\"><span class=\"label-text w-full\">Max hops (<span id=\"maxhopsfeedback\"></span>)</span> <input id=\"maxhopsslider\" name=\"maxhopsslider\" type=\"range\" min=\"10\" max=\"100\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var13 string
		templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(prefs.MaxHops))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 241, Col: 140}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"range range-info range-xs\" hx-post=\"/hops/savemaxhops\" hx-trigger=\"change\"></label></form></li><li><form><label class=\"flex cursor-pointer\"><span class=\"label-text w-full\">Probe timeout (<span id=\"timeoutfeedback\"></span>)</span> <input id=\"timeoutslider\" name=\"timeoutslider\" type=\"range\" min=\"3\" max=\"10\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var14 string
		templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(prefs.Timeout))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 249, Col: 138}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"range range-info range-xs\" hx-post=\"/hops/savetimeout\" hx-trigger=\"change\"></label></form></li><li><form><label class=\"flex cursor-pointer\"><span class=\"label-text w-full\">Probes / poke (<span id=\"probesfeedback\"></span>)</span> <input id=\"probesslider\" name=\"probesslider\" type=\"range\" min=\"1\" max=\"10\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var15 string
		templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(prefs.ProbeCount))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 257, Col: 139}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"range range-info range-xs\" hx-post=\"/hops/saveprobes\" hx-trigger=\"change\"></label></form></li><li><form><label class=\"flex cursor-pointer\"><span class=\"label-text w-full\">Jitter samples (<span id=\"jitterfeedback\"></span>)</span> <input id=\"jitterslider\" name=\"jitterslider\" type=\"range\" min=\"4\" max=\"10\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var16 string
		templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(prefs.JitterSamples))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 265, Col: 142}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"range range-info range-xs\" hx-post=\"/hops/savejitter\" hx-trigger=\"change\"></label></form></li></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Modal(title string, message string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var17 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var17 == nil {
			templ_7745c5c3_Var17 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<dialog id=\"modal\" class=\"modal modal-open\"><div class=\"modal-box\"><h3 class=\"text-lg font-bold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var18 string
		templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 276, Col: 40}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><p class=\"py-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var19 string
		templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(message)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 277, Col: 28}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"modal-action\"><form method=\"dialog\"><button class=\"btn\" hx-trigger=\"click\" hx-get=\"/closemodal\" hx-target=\"#modal\" hx-swap=\"outerHTML\">Close</button></form></div></div></dialog>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func OOBButton(action string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var20 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var20 == nil {
			templ_7745c5c3_Var20 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if action == "search" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"input input-bordered\" type=\"text\" name=\"subject\" id=\"subject\" placeholder=\"subject\" disabled> <button class=\"btn\" type=\"submit\" id=\"subject-submit\">Stop</button> <span class=\"loading loading-bars inline-block align-middle\"></span>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else if action == "busy" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"input input-bordered\" type=\"text\" name=\"subject\" id=\"subject\" placeholder=\"subject\" disabled> <button class=\"btn cfr-busy\" type=\"submit\" id=\"subject-submit\" disabled>Stopping</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input class=\"input input-bordered\" type=\"text\" name=\"subject\" id=\"subject\" placeholder=\"subject\" autofocus> <button class=\"btn\" type=\"submit\" id=\"subject-submit\">Start</button>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		return templ_7745c5c3_Err
	})
}

func Button(classes string, hxUrl string, hxTarget string, hxSwap string, method string, label string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var21 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var21 == nil {
			templ_7745c5c3_Var21 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button class=\"btn { classes }\" type=\"submit\" hx-target=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var22 string
		templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(hxTarget)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 345, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"click\" hx-swap=\"outerhtml transition:false\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if method == "get" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var23 string
			templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(hxUrl)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 349, Col: 17}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if method == "post" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-post=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var24 string
			templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(hxUrl)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 352, Col: 18}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var25 string
		templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs(label)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 354, Col: 9}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func TextInput(name string, placeholder string, method string, endpoint string, hxTrigger string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var26 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var26 == nil {
			templ_7745c5c3_Var26 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		var templ_7745c5c3_Var27 = []any{"input join-item input-bordered input-primary w-full max-w-xs px-4 py-3 my-8"}
		templ_7745c5c3_Err = templ.RenderCSSItems(ctx, templ_7745c5c3_Buffer, templ_7745c5c3_Var27...)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<input type=\"text\" name=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var28 string
		templ_7745c5c3_Var28, templ_7745c5c3_Err = templ.JoinStringErrs(name)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 360, Col: 14}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var28))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" placeholder=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var29 string
		templ_7745c5c3_Var29, templ_7745c5c3_Err = templ.JoinStringErrs(placeholder)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 361, Col: 28}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var29))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-target=\"#result\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if method == "get" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var30 string
			templ_7745c5c3_Var30, templ_7745c5c3_Err = templ.JoinStringErrs(endpoint)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 364, Col: 20}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var30))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if method == "post" {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-post=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var31 string
			templ_7745c5c3_Var31, templ_7745c5c3_Err = templ.JoinStringErrs(endpoint)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 367, Col: 21}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var31))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" hx-trigger=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var32 string
		templ_7745c5c3_Var32, templ_7745c5c3_Err = templ.JoinStringErrs(hxTrigger)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 369, Col: 25}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var32))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var33 string
		templ_7745c5c3_Var33, templ_7745c5c3_Err = templ.JoinStringErrs(templ.CSSClasses(templ_7745c5c3_Var27).String())
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 1, Col: 0}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var33))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func VersionComponent(Version string, UpdateText string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var34 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var34 == nil {
			templ_7745c5c3_Var34 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"stats shadow mt-auto\"><div class=\"stat\"><div class=\"stat-title\">Current App Version</div><div class=\"stat-value\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var35 string
		templ_7745c5c3_Var35, templ_7745c5c3_Err = templ.JoinStringErrs(Version)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 378, Col: 36}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var35))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"stat-desc\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var36 string
		templ_7745c5c3_Var36, templ_7745c5c3_Err = templ.JoinStringErrs(UpdateText)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/templates.templ`, Line: 379, Col: 38}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var36))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func Help() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var37 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var37 == nil {
			templ_7745c5c3_Var37 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<dialog id=\"modal\" class=\"modal modal-open\"><div class=\"modal-box w-11/12 max-w-full\"><h3 class=\"text-lg font-bold\">Help</h3><p class=\"py-2\"><span class=\"font-bold\">Preferences</span></p><p class=\"py-1 px-1\"><span class=\"font-bold\">Hostnames</span> When checked, reverse DNS queries will be performed to resolve the IP addresses being discovered.</p><p class=\"py-2\"><span class=\"font-bold\">Advanced preferences</span></p><p class=\"py-1 px-1\">The current default values are considered the \"sane default\"</p><p class=\"py-1 px-1\"><span class=\"font-bold\">Max Hops</span> Maximum number of hops to check. Stop when we find our final answer, or we reach this number of hops. <em>Default: 30</em></p><p class=\"py-1 px-1\"><span class=\"font-bold\">Probe timeout</span> How many seconds do we wait for a hop to respond. <em>Default: 3</em></p><p class=\"py-1 px-1\"><span class=\"font-bold\">Probes per poke</span> How many probes do we send during each round of hop testing. <em>Default: 3</em></p><div role=\"alert\" class=\"alert\"><svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"none\" viewBox=\"0 0 24 24\" class=\"stroke-info h-6 w-6 shrink-0\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z\"></path></svg> <span>Due to the nature of the underlying protocols, lowering the number of hops  may cause the app to report false packet loss.</span></div><p class=\"py-1 px-1\"><span class=\"font-bold\">Jitter samples</span> Number of samples to base recent jitter measurements on. <em>Default: 4</em></p><div class=\"modal-action\"><form method=\"dialog\"><button class=\"btn\" hx-trigger=\"click\" hx-get=\"/closemodal\" hx-target=\"#modal\" hx-swap=\"outerHTML\">Close</button></form></div></div></dialog>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
