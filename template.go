package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

// printf "%6.2f" 表示6位宽度2位精度
const templ = ` 
{{range .}}----------------------------------------
Name:   {{.Name}}
Price:  {{.Price | printf "%6.2f"}}
{{end}}`

var report = template.Must(template.New("report").Parse(templ))

type Book struct {
	Name  string
	Price float64
}

func main() {
	Data := []Book{ {"《三国演义》", 19.82}, {"《儒林外史》", 99.09}, {"《史记》", 26.89} }
	if err := report.Execute(os.Stdout, Data); err != nil {
		log.Fatal(err)
	}
}

func tHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("index.html.tmpl"))
	t.Execute(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", tHandler)
	http.ListenAndServe(":8080", nil) // 启动web服务
}
