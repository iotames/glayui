package web

import (
	"fmt"
	"net/http"

	"github.com/iotames/glayui/component"
)

// type WebComponent interface {
// 	Handler(w http.ResponseWriter, r *http.Request)
// }

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c Context) NewTable() *component.Table {
	return component.NewTable("")
}

type EasyServer struct {
	httpServer  *http.Server
	routingList []Routing
	middles     []MiddleHandle
}

// NewEasyServer addr like: ":1598", "127.0.0.1:1598"
// You Can SET ENV: USE_EMBED_FILE=true To UseEmbedFile
func NewEasyServer(addr string) *EasyServer {
	fmt.Printf(`
	欢迎使用 GlayUI
	当前版本: v1.0.1
	运行地址: %s
`, addr)
	return &EasyServer{httpServer: newServer(addr)}
}

func (s *EasyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 按顺序依次执行中间件。业务处理逻辑包含在路由中间件里
	for _, m := range s.middles {
		if !m.Handler(w, r) {
			break
		}
	}
}

func (s *EasyServer) SetMiddleware(middles []MiddleHandle) {
	s.middles = middles
}

func (s *EasyServer) AddRouting(routing Routing) {
	s.routingList = append(s.routingList, routing)
}

func (s *EasyServer) AddHandler(method, urlpath string, ctxfunc func(ctx Context)) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		hctx := Context{Writer: w, Request: r}
		ctxfunc(hctx)
	}
	routing := Routing{Methods: []string{method}, Path: urlpath, handler: handler}
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
	for i, m := range s.middles {
		fmt.Printf("---[%d]--EnableMiddleware(%#v)--\n", i, m)
	}
	for i, r := range s.routingList {
		fmt.Printf("---[%d]--RoutePath(%s)---Methods(%+s)--\n", i, r.Path, r.Methods)
	}
	s.middles = append(s.middles, NewMiddleRouter(s.routingList))
	s.httpServer.Handler = s
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
