package main // 10/55 505
import "fmt"

func main() {
	var yuan int = 10
	for i := 0; i < 3; i++ {
		yuan = (yuan-5)*10 + 5
		fmt.Println(yuan)
	}

}
