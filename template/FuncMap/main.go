package main

import (
	"html/template"
	"os"
)

type T struct {
	HTML template.HTML
	ATTR template.HTMLAttr
	URL  template.URL
	JS   template.JS
	CSS  template.CSS
}

func main() {
	data := T{
		HTML: `<div>test div</div>`,
		ATTR: `selected="selected"`,
		URL:  `https://upload.wikimedia.org/wikipedia/commons/5/53/Google_%22G%22_Logo.svg`,
		CSS:  `font-size: 15px`,
		JS:   `console.log("hello world")`,
	}

	_ = template.Must(template.New("Template").Parse(`
        {{.HTML}}
        <option {{.ATTR}} style="{{.CSS}}">test</option>
        <script>{{.JS}}</script>
        <img src="{{.URL}}">
    `)).Execute(os.Stdout, data)

	funcMap := template.FuncMap{
		"attr": func(s string) template.HTMLAttr { return template.HTMLAttr(s) },
		"safe": func(s string) template.HTML { return template.HTML(s) },
	}

	_ = template.Must(
		template.New("Template").Funcs(funcMap).Parse(`
			<option {{ .Attr | attr }}>test</option>
			{{ .Html | safe }}
			`),
	).Execute(os.Stdout, map[string]string{"Attr": `selected="selected"`, "Html": `<option selected="selected">option</option>`})
}
