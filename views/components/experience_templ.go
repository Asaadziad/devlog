// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Experience() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"max-w-2xl mx-auto p-6 bg-gray-100 dark:bg-gray-800 rounded-lg shadow-md\"><h2 class=\"text-3xl font-bold text-gray-800 dark:text-white mb-6 text-center\">Experience</h2><div class=\"bg-white dark:bg-gray-700 p-6 rounded-lg shadow-sm\"><div class=\"flex justify-between items-center mb-2\"><h3 class=\"text-xl font-semibold text-gray-800 dark:text-white\">PriorIT</h3><span class=\"text-sm text-gray-600 dark:text-gray-300\">May 2024 - Present</span></div><p class=\"text-gray-700 dark:text-gray-300 mb-2\">Software Engineer</p><p class=\"text-gray-600 dark:text-gray-400 text-sm\">Currently working on innovative software solutions at PriorIT, contributing to cutting-edge projects and expanding my skillset in a dynamic environment.</p></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
