package main

import (
	"ecommerce/config"
	"fmt"
)

func main() {
	cnf := config.GetEnv()
	fmt.Print(cnf)

}
