package main

import (
	"fmt"
	"husithink/models"
	"log"
	"net/http"
)

// notFoundMiddleware 是一个中间件，用于处理未定义的URL（404 Not Found）。
func notFoundMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 先尝试执行下一个处理器（可能是路由处理器或其它中间件）
		next.ServeHTTP(w, r)

		// 如果响应状态码为200（表示已成功处理请求），则无需做任何事
		if w.Header().Get("Status") != "200" {
			return
		}

		// 如果响应状态码不是200，则说明请求未被成功处理
		// 此时，我们重置状态码为404，并返回“Not Found”响应
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"error")
		http.Error(w, "The requested URL was not found on this server.", http.StatusNotFound)
	})
}

func main() {
	mux := http.NewServeMux()

	// 添加路由处理器
	mux.HandleFunc("/", models.SayhelloName)
	mux.HandleFunc("/login", models.Login)
	mux.HandleFunc("/submit", models.Submit)
	mux.HandleFunc("/enroll", models.Enroll)
	mux.HandleFunc("/upload", models.Upload)
	mux.HandleFunc("/src/*", models.ServeImage)

	// 设置默认的404处理器
	handler := notFoundMiddleware(mux)

	// 启动服务器
	server := &http.Server{
		Addr:    ":9999",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
