package main

import (
	"html/template"
	"log"
	"os"
)

func main()  {
	const tmpl = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`

	t := template.Must(template.New("escape").Parse(tmpl))

	data := struct {
		A string
		B template.HTML
	}{"<b>Hello</b>", "<b>Hello</b>"}
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}


/**


 */