package main

import (
	"fmt"
	"message-broker/sdk"
)

func main() {
	for i := 0; i < 10; i++ {
		sdk.Publish("orders", fmt.Sprintf("message-%d", i))
	}

	fmt.Println("GROUP A first poll:")
	fmt.Println(sdk.Poll("orders", "A"))

	sdk.Poll("orders", "A")

	fmt.Println("GROUP A second poll:")
	fmt.Println(sdk.Poll("orders", "A"))

	fmt.Println("GROUP B:")
	fmt.Println(sdk.Poll("orders", "B"))
}
