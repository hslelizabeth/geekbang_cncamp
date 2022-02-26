package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {

	//03. 读取当前系统的环境变量中的 version 配置，并写入 response header
	os.Setenv("VERSION", "v1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os version: %s \n", version)

	//02. request 中带的 header 写入 response header
	for k, v := range r.Header {
		//fmt.Println(k, v)
		for _, vv := range v {
			w.Header().Set(k, vv)
		}
	}

	//04. Server 端记录访问日志，包括客户端 IP，HTTP 返回码
	clientIP := getCurrentIP(r)
	fmt.Println(r.RemoteAddr)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! clientip: %s", clientIP)
}

func getCurrentIP(r *http.Request) string {
	// 这里也可以通过X-Forwarded-For请求头的第一个值作为用户的ip
	// 但是要注意的是这两个请求头代表的ip都有可能是伪造的
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		// 当请求头不存在即不存在代理时直接获取ip
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func clientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

// 05. 检查健康的路由
func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("start http server failedj, error: %s\n", err.Error())
	}
}
