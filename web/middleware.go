package web

import (
	"io"
	"net/http"
	"os"
	"strings"
)

// GetMiddlewareList 获取中间件列表。按顺序依次执行中间件
func GetMiddlewareList() []func(w http.ResponseWriter, r *http.Request) (subNext bool, mainNext bool) {
	return []func(w http.ResponseWriter, r *http.Request) (subNext bool, mainNext bool){
		middleCORS,
		middleStatic,
		MiddleRouter,
	}
}

// middleware 中间件的处理逻辑。按顺序依次执行中间件
func middleware(w http.ResponseWriter, r *http.Request) bool {
	middles := GetMiddlewareList()
	var next bool
	for _, h := range middles {
		var next1 bool
		next1, next = h(w, r)
		if !next1 {
			break
		}
	}
	return next
}

// middleCORS 跨域设置
func middleCORS(w http.ResponseWriter, r *http.Request) (subNext bool, mainNext bool) {
	w.Header().Add("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
	w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	w.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Length, Content-Type, Accept, Token, Auth-Token, X-Requested-With")
	w.Header().Add("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.Method == "OPTIONS" {
		return false, false
	}
	return true, true
}

func getStaticUrlPath() []string {
	return []string{
		"/static/",
	}
}

// middleStatic 定义静态资源
func middleStatic(w http.ResponseWriter, r *http.Request) (subNext bool, mainNext bool) {
	rpath := r.URL.Path
	fpath := "resource" + rpath
	for _, v := range getStaticUrlPath() {
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

func errWrite(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	w.Write([]byte(msg))
}
