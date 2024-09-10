// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package main

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func hello() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><title>snwy.me</title><style type=\"text/css\">\n\tbody {\n    width: 50%;\n    margin: 0 auto;\n    background: #151515;\n    font-family: \"Space Mono\", monospace;\n    font-size: 19px;\n    color: #fff;\n    line-height: 2.0;\n    display: table;\n}\n\nhtml,body\n{\n    height: 100%;\n}\n\nhtml::after {\n    content:'';\n    position: absolute;\n    top: 0;\n    width:100%;\n    height: 100%;\n    background-image: linear-gradient(rgba(0,0,0,0.4) 1px, transparent 1px);\n    background-size: 2px 2px;\n    background-repeat: repeat;\n    pointer-events: none;\n}\n\n@media screen and (max-width: 1300px) {\n    body {\n        width: 90%;\n    }\n}\n\nh1 {\n    font-size: 2.5em;\n    margin: 0;\n    padding: 0;\n}\n\nheader {\n    padding: 0 20px 20px 0;\n    display: flex;\n    align-items: center;\n}\n\nul {\n    padding-left: 0;\n    margin-top: 0;\n    list-style: none;\n}\n\na {\n    color: #fff;\n    text-decoration: white underline 2px;\n}\n\nhr {\n    height: 4px;\n    background: #fff;\n    opacity: 25%;\n    border: none;\n}\n\n#header_block {\n    margin-left: 5%;\n}\n\n#banner {\n    height: auto;\n    width: 35%;\n}\n\n#content {\n    height: 100%;\n    display: table-cell;\n    vertical-align: middle;\n    padding: 0;\n    margin: 0;\n}\n\t</style><style>\n    @import url('https://fonts.googleapis.com/css2?family=Space+Mono&display=swap');\n  </style><link rel=\"stylesheet\" href=\"./index.css\"></head><body><div id=\"content\"><header><img id=\"banner\" src=\"banner.jpg\"><div id=\"header_block\"><h1>snwy.me</h1><nav><ul><li>[<a href=\"https://github.com/snoglobe\">github</a>]</li><li>[<a href=\"https://x.com/snwy_me\">x</a>]</li><li>[<a href=\"https://www.youtube.com/@snwydotme\">youtube</a>]</li></ul></nav></div></header><hr><p>hey, i'm Mason Meirs - a.k.a. snwy. i'm a developer interested in programming language design and theory. i'm also interested in distributed systems, machine learning, and the future of computing.</p><p>i dropped out of high school in 2024 and got my GED to puruse a career.</p><p>i've built <a href=\"https://aurora-lang.org\">aurora</a>, a language built to be an automation tool that replaces the languages that came before it.</p><p></p><hr><p>contact me:</p><ul><li>[email: <a href=\"mailto:snwy@snwy.me\">snwy@snwy.me</a>]</li><li>[discord: snwy]</li></ul><hr><p>thanks for visiting!</p></div></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
