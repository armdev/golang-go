package main
	
import (
    "fmt"
    "math"
)

const server_path string = "http://axele.com"


func main (){

	fmt.Println(server_path)

	const n = 500000000
	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println("Int 64: ", int64(d))
	
	fmt.Println(math.Sin(n))

}