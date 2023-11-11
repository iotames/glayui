package resource

import (
	"embed"
	// "fmt"
	// "io/fs"
	// "os"
)

const RESOURCE_DIR = "resource"

//go:embed *
var ResourceFs embed.FS

// func GetResourceFile(fullpath string) (fss fs.FS, err error) {
// 	// if strings.Index(rescpath, "/") == 0 {
// 	// 	rescpath = rescpath[1:]
// 	// }
// 	// fullpath = resourceDir + "/" + rescpath
// 	// fullpath = filepath.Join(resourceDir, rescpath)
// 	fmt.Printf("----GetResourceFile--fullpath(%s)---\n", fullpath)
// 	var finfo fs.FileInfo
// 	finfo, err = os.Stat(fullpath)
// 	pwd, _ := os.Getwd()
// 	if err != nil {
// 		fmt.Printf("-----GetResourceFile--os.Stat-err(%v)-InPath(%s)\n", err, pwd)
// 		return
// 	}
// 	if finfo.IsDir() {
// 		err = fmt.Errorf("path(%s) is dir", fullpath)
// 		return
// 	}
// 	fss = ResourceFs
// 	return
// }
