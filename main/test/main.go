package main

import (
	"fmt"
	"go_template/hello"
	m "go_template/meet"
	"net/http"

	"github.com/go-sql-driver/mysql"

	"github.com/karlseguin/ccache"
)

func main() {
	a := mysql.Config{}
	fmt.Println(a.Addr)
	fmt.Println(hello.SpeekHello())
	fmt.Println(m.Te)
	http.NewRequest("POST", "sdsds", nil)
	fmt.Println(ccache.Cache{})
}
