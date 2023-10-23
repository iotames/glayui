package web

import (
	"fmt"
	"net/http"
)

func NewServer(addr string) *http.Server {
	server := http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(httpHandler),
		// MaxHeaderBytes: 1 << 20, // 1048576
	}

	fmt.Printf(`
	欢迎使用 GlayUI
	当前版本: v1.0.1
	运行地址: %s
`, addr)
	return &server
}
