package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main(){
	http.HandleFunc("/",myWeb)

	////指定相对路径
	//statciHandler := http.FileServer(http.Dir("./static"))
	////将/js/路径下的请求匹配到 ./static/js/下
	//http.Handle("/js/",statciHandler)
	http.Handle("/js/", http.FileServer(http.Dir("./static"))) //对应到 ./static/js
	http.Handle("/css/", http.FileServer(http.Dir("./static"))) //对应到 ./static/css
	http.Handle("/img/", http.FileServer(http.Dir("./static"))) //对应到 ./static/img
	http.Handle("/upload/", http.FileServer(http.Dir("./static"))) //对应到 ./static/upload

	fmt.Println("服务器正在启动，访问地址：http://localhost:8080")
	err:= http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Println("服务器错误：",err)
	}
}

func myWeb(w http.ResponseWriter,r *http.Request){

	_ = r.ParseForm()

	for k,v := range r.URL.Query(){
		fmt.Println("key:",k," value:",v[0])
	}

	for k,v := range r.PostForm{
		fmt.Println("key:", k, ", value:", v[0])
	}

	//t := template.New("index")
	//t.Parse("<div id='template-div'>Hi, {{.name}}, {{.someStr}}</div>")
	t,_ := template.ParseFiles("./templates/index.html")

	data := map[string]string{
		"name": "jack",
		"someStr":"欢迎使用",
	}
	_ = t.Execute(w, data)

	//fmt.Fprintf(w, "欢迎使用")
}
