package resource

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

//go:embed *
var resourceFs embed.FS

func GetResourceFile(rescpath string) (fullpath string, fss fs.FS, err error) {
	if strings.Index(rescpath, "/") == 0 {
		rescpath = rescpath[1:]
	}
	fullpath = "resource/" + rescpath
	var finfo fs.FileInfo
	finfo, err = os.Stat(fullpath)
	if err != nil {
		return
	}
	if finfo.IsDir() {
		err = fmt.Errorf("path(%s) is dir", fullpath)
		return
	}
	fss = resourceFs
	return
}
