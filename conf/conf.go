package conf

import (
	"os"
	"strings"
)

// 是否使用嵌入到可执行程序里的静态资源文件。默认false
func UseEmbedTpl() bool {
	val, ok := os.LookupEnv("USE_EMBED_TPL")
	if ok {
		if strings.EqualFold(val, "true") || val == "1" {
			return true
		}
	}
	return false
}
