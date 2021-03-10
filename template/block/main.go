package main

import (
	"os"
	"strings"
	"text/template"
)

func main() {
	const (
		master = `Names:
{{ block "list" . }}
	{{ "\n" }}
	{{ range . }}
		{{ println "-" . }}
	{{ end }}
{{ end }}`
		overlay = `{{ define "list" }}{{ join . ", " }}{{ end }}`
	)
	var (
		funcs     = template.FuncMap{"join": strings.Join}
		guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
	)
	masterTmpl, _ := template.New("master").Funcs(funcs).Parse(master)
	overlayTmpl, _ := template.Must(masterTmpl.Clone()).Parse(overlay)

	_ = masterTmpl.Execute(os.Stdout, guardians)
	_ = overlayTmpl.Execute(os.Stdout, guardians)
}
