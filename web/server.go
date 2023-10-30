package web

import (
	"fmt"
	"net/http"
)

type EasyServer struct {
	httpServer  *http.Server
	routingList []Routing
	middles     []MiddleHandle
}

func NewEasyServer(addr string) *EasyServer {
	fmt.Printf(`
	欢迎使用 GlayUI
	当前版本: v1.0.1
	运行地址: %s
`, addr)
	return &EasyServer{httpServer: newServer(addr)}
}

func (s *EasyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !middleware(w, r, s.middles) {
		return
	}
}

func (s *EasyServer) SetMiddleware(middles []MiddleHandle) {
	s.middles = middles
}

func (s *EasyServer) AddRouting(routing Routing) {
	s.routingList = append(s.routingList, routing)
}

func (s *EasyServer) ListenAndServe() error {
	if len(s.middles) == 0 {
		s.middles = GetDefaultMiddlewareList()
	}
	if len(s.routingList) == 0 {
		fmt.Printf("----routingList不能为空。已启用默认路由设置。请使用SetRouting()方法添加路由，以处理业务逻辑-----\n")
		s.routingList = GetDefaultRoutingList()
	}
	s.middles = append(s.middles, NewMiddleRouter(s.routingList))
	s.httpServer.Handler = s
	for i, m := range s.middles {
		fmt.Printf("-----[%d]启用中间件(%+v)--\n", i, m)
	}
	for i, r := range s.routingList {
		fmt.Printf("-----[%d]路由(%s)----请求方法(%+s)--\n", i, r.Path, r.Methods)
	}
	return s.httpServer.ListenAndServe()
}

func newServer(addr string) *http.Server {
	server := http.Server{
		Addr: addr,
		// Handler: http.HandlerFunc(httpHandler),
		// MaxHeaderBytes: 1 << 20, // 1048576
	}
	return &server
}
