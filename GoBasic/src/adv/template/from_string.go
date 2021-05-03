package main

import (
	"fmt"
	"log"
	"text/template"
	"time"

	//"html/template"
	"os"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}
func main() {
	//.代表传入的参数
	t1, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		panic(err)
	} else {
		err = t1.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")
		//"html/template" 包结果对<做转义为&lt;
		//"text/template" 包结果对<不做处理
	}
	//管道
	t1, err = template.New("piple").Parse(`{{ 12.3456 | printf "%.2f"  }} `)
	if err != nil {
		panic(err)
	} else {
		err = t1.Execute(os.Stdout, "")
	}

	//with后面的要用双引号包起，表示这块区域的.表示这个值
	t1, err = template.New("alias").Parse(`cat is {{.}}, {{ with "arg" }}  temp arg  is {{.}} {{end}} `)
	if err != nil {
		panic(err)
	} else {
		err = t1.Execute(os.Stdout, "cat")
	}
	//变量
	t1, err = template.New("vari").Parse(`{{ range $key,$val:= . }} key= {{$key}} ,val={{$val}} {{end}} `)
	if err != nil {
		panic(err)
	} else {
		kvs := map[string]string{"a": "apple", "b": "banana"}
		err = t1.Execute(os.Stdout, kvs)
	}
	//if 有某个变量  gt
	t1, err = template.New("vari").Parse(`  
	{{ if . }}
      Number has value 
    {{ else }}
      Number is  empty
	{{ end }} 
	
	{{if gt .Age 18}}
	<p>hello, old man, {{.Name}}</p>
	{{else}}
	<p>hello,young man, {{.Name}}</p>
	{{end}} 
	`)
	if err != nil {
		panic(err)
	} else {
		type Person struct {
			Name string
			Age  int
		}
		p := Person{Name: "safly", Age: 30}
		err = t1.Execute(os.Stdout, p)
	}
	//模板 自定义函数
	funcMap := template.FuncMap{"fdate": formatDate} //自己的函数
	tFunc := template.New("customFUnc").Funcs(funcMap)
	tFunc, _ = tFunc.Parse("The date/time is {{ . | fdate }}")
	tFunc.Execute(os.Stdout, time.Now())

	/*block用类似于template语法的使用，但是不同的是block会有一个默认值，而template没有默认值
	<body>{{block "content" .}}This is the default body.{{end}}</body>
	如果你的content模板没有任何匹配的定义，将会显示默认的内容
	*/

	const tpl = `
	<!DOCTYPE html>
	<html>
		<head>
			<meta charset="UTF-8">
			<title>{{.Title}}</title>
		</head>
		<body>
			{{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
		</body>
	</html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	t, err := template.New("webpage").Parse(tpl)
	check(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My page",
		Items: []string{
			"My photos",
			"My blog",
		},
	}

	err = t.Execute(os.Stdout, data)
	check(err)

	noItems := struct {
		Title string
		Items []string
	}{
		Title: "My another page",
		Items: []string{},
	}

	err = t.Execute(os.Stdout, noItems)
	check(err)

	//--转义
	const s = `"Fran & Freddie's Diner" <tasty@example.com>`
	v := []interface{}{`"Fran & Freddie's Diner"`, ' ', `<tasty@example.com>`}

	fmt.Println(template.HTMLEscapeString(s))
	template.HTMLEscape(os.Stdout, []byte(s))
	fmt.Fprintln(os.Stdout, "")
	fmt.Println(template.HTMLEscaper(v...))

	fmt.Println(template.JSEscapeString(s)) //& 变 \u开头，‘ 变 \' ，< 变 \u开头
	template.JSEscape(os.Stdout, []byte(s))
	fmt.Fprintln(os.Stdout, "")
	fmt.Println(template.JSEscaper(v...)) //空格

	fmt.Println(template.URLQueryEscaper(v...)) //%形式

}
