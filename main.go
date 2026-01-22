package main

import (
	"fmt"

	"github.com/antchfx/xmlquery"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/gogo/protobuf/proto"

	// - "github.com/gogs/gogs"
	_ "github.com/hashicorp/golang-lru"
	_ "github.com/owncast/owncast/logging"
)

func main() {
	pat := "glpat-Qr-GyZZiilXWYGuScic8gW86MQp1OmV0Z3ZsCw.01.120xrlxxx"
	fmt.Printf("I am randomly printing a PAT %s\n", pat)

	pat1 := "glpat-Qr-GyZZiilXWYGuScic8gW86MQp1OmV0Z3ZsCw.01.121xrlxxx"
	fmt.Printf("I am randomly printing a PAT %s\n", pat1)

	pat2 := "glpat-Qr-GyZZiilXWYGuScic8gW86MQp1OmV0Z3ZsCw.01.122xrlxxx"
	fmt.Printf("I am randomly printing a PAT %s\n", pat2)

	pat3 := "glpat-Qr-GyZZiilXWYGuScic8gW86MQp1OmV0Z3ZsCw.01.123xrlxxx"
	fmt.Printf("I am randomly printing a PAT %s\n", pat3)

	newsecret := "glpat-Qr-GyZZiilXWYGuScic8gW86MQp1OmV0Z3ZsCq.01.120yyyxxx"
	fmt.Printf("I am randomly printing a new secret %s\n", newsecret)

	// ==== 150 More Random GitLab PATs ====
	fmt.Println("glpat-AbcDefGhijkLmNoPqRsTuVwXyZ123456.01.001aaa111")
	fmt.Println("glpat-AbcDefGhijkLmNoPqRsTuVwXyZ123456.01.002bbb222")
	fmt.Println("glpat-AbcDefGhijkLmNoPqRsTuVwXyZ123456.01.003ccc333")
	fmt.Println("glpat-AbcDefGhijkLmNoPqRsTuVwXyZ123456.01.004ddd444")
	fmt.Println("glpat-AbcDefGhijkLmNoPqRsTuVwXyZ123456.01.005eee555")
	fmt.Println("glpat-AbcDefGhijkLmNoPqRsTuVwXyZ123456.01.006fff666")
	fmt.Println("glpat-AbcDefGhijkLmNoPqRsTuVwXyZ123456.01.007ggg777")
	fmt.Println("glpat-AbcDefGhijkLmNoPqRsTuVwXyZ123456.01.008hhh888")
	fmt.Println("glpat-AbcDefGhijkLmNoPqRsTuVwXyZ123456.01.009iii999")
	wadl, err := xmlquery.LoadURL("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	attr := xmlquery.FindOne(wadl, "//application/@xmlns")
	fmt.Println(attr.InnerText())
}
