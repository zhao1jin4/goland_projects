package main

import (
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

func simpleHttp() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "测试, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	//http://localhost:8080/bar
	//http://localhost:8080/test
}
func index(writer http.ResponseWriter, request *http.Request) { //debug时这个方法可能会因超时被重复调用
	fmt.Println("---begin index")
	var t = template.New("layout")
	t = template.Must(t.ParseFiles("templates/index.html", "templates/private.navbar.html"))

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
	t.ExecuteTemplate(writer, "layout", data)
	fmt.Println("---end index")
}
func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("---begin login")
	var email = request.PostFormValue("email")
	fmt.Println("email=", email)

	cookie, err := request.Cookie("_session_id")
	if cookie != nil && err == nil {
		fmt.Println("_session_id=", cookie.Value)
	} else {
		cookie := http.Cookie{
			Name:     "_session_id",
			Value:    "123abc",
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie) //响应头并没有Set-Cookie

	}
	writer.Write([]byte("cookie 保存了"))
	//http.Redirect(writer, request, "index?id=123", 302) //客户端跳转带不了cookie?
	fmt.Println("---end login")
}

func upload(w http.ResponseWriter, r *http.Request) {
	var maxUploadSize int64 = 1024 * 1024 * 2 //2MB
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		fmt.Println("---FILE_TOO_BIG")
		panic(err)
	}
	var user_id = r.FormValue("user_id") //GET
	fmt.Println("---user_id=", user_id)
	fileType := r.PostFormValue("type")
	file, _, err := r.FormFile("uploadFile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)

	filetype := http.DetectContentType(fileBytes)
	if filetype != "image/jpeg" && filetype != "image/jpg" &&
		filetype != "image/gif" && filetype != "image/png" &&
		filetype != "application/pdf" {
		fmt.Println("INVALID_FILE_TYPE")
		return
	}

	fileName := "my_file_name"
	newPath := filepath.Join("d:/tmp", fileName)
	fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)
	newFile, err := os.Create(newPath)

	defer newFile.Close()
	if _, err := newFile.Write(fileBytes); err != nil {
		fmt.Printf("CANT_WRITE_FILE")
		return
	}
	w.Write([]byte("SUCCESS"))
}
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fileName := r.Form["filename"]            //filename  文件名
	path := "d:/tmp/"                         //文件存放目录
	fileNames := url.QueryEscape(fileName[0]) // 防止中文乱码
	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileNames+"\"")
	//---
	http.ServeFile(w, r, path+fileName[0])
	//---
	/*
		file, err := os.Open(path + fileName[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println("Read File Err:", err.Error())
		} else {
			w.Write(content)
		}
	*/
}
func templateHttp() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))                //相对当前文件的public目录放js,image等文件
	mux.Handle("/static/", http.StripPrefix("/static/", files)) //请求static转到public目录
	// http://localhost:8081/static/go-logo-blue.svg
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/upload", upload)
	mux.HandleFunc("/download", downloadHandler)
	// http://localhost:8081/index
	// http://localhost:8081/login
	// http://localhost:8081/upload
	//http://localhost:8081/download?filename=中文.txt  文件内容要为UTF-8
	server := &http.Server{
		Addr:           "0.0.0.0:8081",
		Handler:        mux,
		ReadTimeout:    time.Duration(3 * int64(time.Second)),
		WriteTimeout:   time.Duration(2 * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
func main() {
	//simpleHttp()
	templateHttp() //上传

	//restful
	//服务端 接受 文件 下载 未做？？

}
