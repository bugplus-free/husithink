package models

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"husithink/sessions"
	_ "husithink/sessions/memory"
)
var GlobalSessions *sessions.Manager

func init() {
	GlobalSessions, _ = sessions.NewManager("memory", "gosessionid", 3600)
	go GlobalSessions.GC()
}


func SayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)

	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	if r.Method == "GET" {
		timestamp := strconv.Itoa(time.Now().Nanosecond())
		hashWr := md5.New()
		hashWr.Write([]byte(timestamp))
		token := fmt.Sprintf("%x", hashWr.Sum(nil))
		//如果是请求login页的话就给他
		t, _ := template.ParseFiles("views/begin.gtpl")
		log.Println(t.Execute(w, token))
	} else {
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		fmt.Fprintf(w, "该功能正在开发中...")
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	r.ParseForm()
	h := md5.New()
	salt:="astaxie%^7&8888"
	io.WriteString(h,salt+time.Now().String())
	token:=fmt.Sprintf("%x",h.Sum(nil))
	if r.Form.Get("token")!=token{
		t, _ := template.ParseFiles("views/login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
		return 
	}
	sess.Set("token",token)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, sess.Get("username"))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		user := Userinfo{
			UserName: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}
		if If_In_Sqlite3(&user) {
			fmt.Println("用户登录成功")
			t, _ := template.ParseFiles("views/submit.gtpl")
			w.Header().Set("Content-Type", "text/html")
			t.Execute(w, sess.Get("username"))
		} else {
			sess.Set("username", r.Form["username"])
			http.Redirect(w, r, "/", 302)
		}
		// }
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		// template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}
func Submit(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		//如果是请求login页的话就给他
		// t, _ := template.ParseFiles("views/submit.gtpl")
		// log.Println(t.Execute(w, nil))
		t, err := template.New("foo").Parse(`{{define "T"}}{{.}}{{end}}`)
		err = t.ExecuteTemplate(w, "T", template.HTML("<h1>千里之行，从这里开始呀</h1>"))
		if err != nil {
			fmt.Fprintf(w, "error")
		}

	} else {
		fmt.Fprintf(w, "该功能未开发...")
	}
}
func Enroll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		//如果是请求login页的话就给他
		t, _ := template.ParseFiles("views/enroll.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		user := Userinfo{
			UserName: r.Form.Get("username"),
			Email:    r.Form.Get("email"),
			Password: r.Form.Get("password"),
		}
		if !If_Add_Sqlite3(&user) {
			fmt.Println("用户已注册")
		} else {
			fmt.Println("用户注册成功")
			t, _ := template.ParseFiles("views/login.gtpl")
			log.Println(t.Execute(w, token))
		}
		// }
		fmt.Println("username:", r.Form["username"])
		fmt.Println("email:", r.Form["email"])
		fmt.Println("password:", r.Form["password"])
	}
}
func Upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20) //开一个内存空间出来
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./src/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

// 示例函数：根据文件名推断MIME类型
func getContentTypeFromFileName(fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	// 添加其他格式的支持...
	default:
		return "application/octet-stream"
	}
}
func ServeImage(w http.ResponseWriter, r *http.Request) {
	// 获取请求URL的路径
	urlPath := r.URL.Path

	// 提取URL路径的最后一段（图片文件名）
	fileName := filepath.Base(urlPath)

	// 构建图片文件的完整路径
	srcDir := "./src/images" // 调整为实际的src目录相对路径
	fullPath := filepath.Join(srcDir, fileName)

	// 读取图片文件
	imgData, err := ioutil.ReadFile(fullPath)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// 设置Content-Type和Content-Length响应头
	contentType := getContentTypeFromFileName(fileName) // 实现getContentTypeFromFileName函数以根据文件名推断MIME类型
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(imgData)))

	// 写入图片数据到响应体
	if _, err := w.Write(imgData); err != nil {
		http.Error(w, fmt.Sprintf("Failed to write image data: %v", err), http.StatusInternalServerError)
		return
	}
}

func Count(w http.ResponseWriter, r *http.Request) {
	sess := GlobalSessions.SessionStart(w, r)
	ct := sess.Get("countnum")
	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}
	t, _ := template.ParseFiles("views/count.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, sess.Get("countnum"))
}
