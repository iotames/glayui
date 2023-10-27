package main

import (
	"fmt"

	"github.com/iotames/glayui/web"
)

func main() {
	s := web.NewServer(":1598")
	s.ListenAndServe()
}

func HelloWordd() {
	fmt.Println("HELLO GLAYUI")
}
