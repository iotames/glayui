package main

import (
	"github.com/iotames/glayui/web"
)

func main() {
	s := web.NewServer(":1598")
	s.ListenAndServe()
}
