/*
 * @Author: iRorikon
 * @Date: 2023-04-14 16:43:30
 * @FilePath: \api-service\main.go
 */
package main

import (
	"fmt"

	"github.com/irorikon/api-service/command"
)

func main() {
	fmt.Println("Hello World!")
	command.Execute()
}
