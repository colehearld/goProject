package main

import (
	"fmt"
	t "goProject/PacketState"
)
func main() {
	packets()
	var ImageName string
	fmt.Println("Enter a docker image: ")
	fmt.Scanln(&ImageName)
	t.RunContainer(ImageName)

}
