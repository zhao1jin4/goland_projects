package main

//go test github.com/spf13/pflag
import (
	"fmt"

	flag "github.com/spf13/pflag"
)

var flagvar int

func init() {
	flag.IntVar(&flagvar, "files", 1234, "help message for flagname")
}

func main() {
	// go run  .\src\k8s\k8s_pflag.go  --flagname 11 --files 22
	//混合用 go run  .\src\k8s\k8s_pflag.go  --flagname=11 --files 22 -p 888 --auth
	var ip *int = flag.Int("flagname", 1234, "help message for flagname")

	var port *int = flag.IntP("port", "p", 8081, "port number")
	var authVar bool
	flag.BoolVarP(&authVar, "auth", "a", false, "is set password")

	flag.Parse()

	fmt.Println("ip has value ", *ip)
	fmt.Println("flagvar has value ", flagvar)

	fmt.Println("port ", *port)
	fmt.Println("authVar ", authVar)

}
