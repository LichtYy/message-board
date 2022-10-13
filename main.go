package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type User struct {
	username string
	password string
}

func index(w http.ResponseWriter, r *http.Request) {
	// 首页
	fmt.Println("进入首页：")
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
	fmt.Println(w.Header(), r.Body, r.Header)
	fmt.Println(r.URL)
	// if err != nil {
	// 	t.Execute(w, nil)
	// } else {
	// 	log.Fatal("首页加载失败")
	// }
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Println("----------index end----------")
}

func login(w http.ResponseWriter, r *http.Request) {
	// 登录
	fmt.Println("进入login:")
	fmt.Println("method:", r.Method) // 获取请求的方法
	fmt.Println(w.Header(), r.Body)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		// 请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		user := User{strings.Join(r.Form["username"], ""), strings.Join(r.Form["password"], "")}
		if !judgePassword(user) {
			// 账号密码错误，不予登录
			fmt.Fprintf(w, "你输入的是：%s %s", user.username, user.password)
			fmt.Fprintf(w, "密码错误（应为123）")
		}
	}
	fmt.Println("----------login end----------")
}

func judgePassword(user User) bool {
	// 判断密码是否正确（这里需要连接数据库了）
	if user.password != "123" {
		return false
	}
	return true
}

func account(w http.ResponseWriter, r *http.Request) {
	// 个人账户
	fmt.Println("进入个人账户页面：")
	fmt.Println(w.Header(), r.Body)
	r.ParseForm() // 解析url传递的参数，对于POST则解析响应包的主体(request body)
	// 注意：如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) // 这些信息时输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	t, _ := template.ParseFiles("account.html")
	// w.Write([]byte("Hello " + r.Form["username"][0] + " !"))

	t.Execute(w, nil) // 这个写入到w的是输出到客户端的
	fmt.Println("----------account end----------")
}

func about(w http.ResponseWriter, r *http.Request) {
	// 进入附加信息页面
	fmt.Println("进入about：")
	t, _ := template.ParseFiles("about.html")
	t.Execute(w, nil)
	fmt.Println("----------about end----------")
}

func main() {
	http.HandleFunc("/", index)      // 设置访问的路由
	http.HandleFunc("/login", login) // 设置访问的路由
	http.HandleFunc("/account", account)
	http.HandleFunc("/about", about)
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
