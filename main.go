package main

import (
	"fmt"

	gosayhello "github.com/faisalfajarns/Go_Modules_Faisal/v2"
)

func main(){
testPackage := gosayhello.SayHello("Faisal")

fmt.Println(testPackage)
}