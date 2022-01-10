package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
)

func main() {

	http.HandleFunc("/userInfo", userInfo) //Use the default DefaultServeMux.

	http.HandleFunc("/healthz", healthz) //Use the default DefaultServeMux.
	// • ListenAndService
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
	// • 定义 handle 处理函数

}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func userInfo(w http.ResponseWriter, r *http.Request) {

	//打印访问日志->先实现 之后拆出作为filter  包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	fmt.Printf("the client ip is %s \n", r.Host)
	//获取request header，推给response header
	var tk = r.Header.Get("testKey")
	if tk == "" {
		tk = "lilei is ready to code"
	}
	//获取version 推给header
	var goversion string = runtime.Version()

	w.Header().Add("testKey", tk)
	w.Header().Add("goversion", goversion)

	//setcookie
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Programming",
		HttpOnly: true,
	}
	//w.Header().Set("Set-Cookie",c1.String())
	//w.Header().Add("Set-Cookie",c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2) //三种设置cookie的方法

	//http 返回码设置有http包的const提供
	w.WriteHeader(http.StatusOK)
	//http.StatusText获取http code 对应的描述
	fmt.Printf("httpCode is %d and httpStatus is %s", http.StatusOK, http.StatusText(http.StatusOK))
	io.WriteString(w, "200")
}

//需要把本章节学习的内容放进去  包依赖部分引入verdor 对依赖的lib进行适度修改  实现支持json的接口
