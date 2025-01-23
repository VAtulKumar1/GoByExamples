package chapters

import (
	"html/template"
	"os"
)

func TextTemplate() {

	t1 := template.New("t1")
	_, err := t1.Parse("value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"go",
		"rust",
		"C++",
		"C#",
	})

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"jane doe"})

	t1.Execute(os.Stdout, map[string]string{
		"Name": "Atul",
	})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")

	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range : {{range .}}{{.}} {{end}}\n")

	t4.Execute(os.Stdout, []string{
		"go",
		"rust",
		"c++",
		"c#",
	})
}
