package web

import (
	"io"
	"net/http"
	"os"
	"strings"
)

type MiddleHandle interface {
	Handler(w http.ResponseWriter, r *http.Request) (subNext bool, mainNext bool)
}

// middleStatic 静态资源中间件
type middleStatic struct {
	wwwrootDir    string
	staticUrlPath []string
}

// middleCORS CORS跨域设置中间件
type middleCORS struct {
	allowOrigin string
}

// middleRouter 路由中间件。处理业务逻辑
type middleRouter struct {
	routingList []Routing
}

func NewMiddleRouter(routingList []Routing) *middleRouter {
	return &middleRouter{routingList: routingList}
}

// Handler 路由中间件。处理业务逻辑
func (m middleRouter) Handler(w http.ResponseWriter, r *http.Request) (subNext bool, mainNext bool) {
	routings := m.routingList
	isMatch := false
	rpath := r.URL.Path
	rmethod := r.Method

	for _, rt := range routings {
		if rt.Path == rpath {
			if len(rt.Methods) == 0 {
				isMatch = true
			} else {
				for _, m := range rt.Methods {
					// strings.ToUpper(m) == strings.ToUpper(rmethod)
					if strings.EqualFold(m, rmethod) {
						isMatch = true
					}
				}
			}
			if isMatch {
				rt.handler(w, r)
				break
			}
		}
	}

	if !isMatch {
		ResponseNotFound(w, r)
	}

	return true, true
}

// 静态资源中间件 NewMiddleStatic
// wwwroot 网站根目录。字符串末尾不用添加斜杠/。默认值为 "resource"
// urlPathBegin 启用静态资源的URL路径。必须以正斜杠/开头和结尾。如 []string{"/static/"}
func NewMiddleStatic(wwwroot string, urlPathBegin []string) *middleStatic {
	if wwwroot == "" {
		wwwroot = "resource"
	}
	return &middleStatic{wwwrootDir: wwwroot, staticUrlPath: urlPathBegin}
}

// middleStatic 定义静态资源
func (m middleStatic) Handler(w http.ResponseWriter, r *http.Request) (subNext bool, mainNext bool) {
	rpath := r.URL.Path
	fpath := m.wwwrootDir + rpath
	for _, v := range m.staticUrlPath {
		if strings.Index(rpath, v) == 0 {
			// 匹配命中URL静态资源
			finfo, err := os.Stat(fpath)
			if err != nil {
				if os.IsNotExist(err) {
					// 文件不存在
					// errWrite(w, "file IsNotExist ", 400)
					return true, true
				} else {
					// 其他错误
					errWrite(w, err.Error(), 500)
				}
				return false, false
			}

			if finfo.IsDir() {
				errWrite(w, "not allow visit dir path", 400)
				return false, false
			}

			f, err := os.Open(fpath)
			if err != nil {
				errWrite(w, err.Error(), 500)
				return false, false
			}
			b, err := io.ReadAll(f)
			if err != nil {
				errWrite(w, err.Error(), 500)
				return false, false
			}
			w.Write(b)
			return false, false
		}
	}
	return true, true
}

// NewMiddleCORS CORS中间件: 跨域设置
// allowOrigin: 允许跨域的站点。默认值为 "*"
func NewMiddleCORS(allowOrigin string) *middleCORS {
	if allowOrigin == "" {
		allowOrigin = "*"
	}
	return &middleCORS{allowOrigin: allowOrigin}
}

func (m middleCORS) Handler(w http.ResponseWriter, r *http.Request) (subNext bool, mainNext bool) {
	w.Header().Add("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Length, Content-Type, Accept, Token, Auth-Token, X-Requested-With")
	w.Header().Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", m.allowOrigin)
	if r.Method == "OPTIONS" {
		return false, false
	}
	return true, true
}

// GetDefaultMiddlewareList 获取中间件列表。按数组列表的顺序依次执行中间件
func GetDefaultMiddlewareList() []MiddleHandle {
	return []MiddleHandle{
		NewMiddleCORS("*"),
		NewMiddleStatic("", []string{"/static/"}),
	}
}

// middleware 中间件的处理逻辑。按顺序依次执行中间件
func middleware(w http.ResponseWriter, r *http.Request, middles []MiddleHandle) bool {
	var next bool
	for _, m := range middles {
		var next1 bool
		next1, next = m.Handler(w, r)
		if !next1 {
			break
		}
	}
	return next
}

func errWrite(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}
