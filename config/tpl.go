package config

import (
	"html/template"
)

var TPL *template.Template

func init() {
	/*funcMap := template.FuncMap{
		"replace": strings.Replace,
	}*/
	//TPL = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*.gohtml"))
	TPL = template.Must(template.ParseGlob("templates/*.gohtml"))
}