package main
	
import (
    "fmt"   
)


func main(){

	i := 1

	for i <=3 {
		fmt.Println(i)
		i = i+1
	}
	fmt.Println("**********************")
	for j := 7; j <= 9; j++ {
        fmt.Println(j)
	}

	fmt.Println("**********************")
	
	for {
        fmt.Println("loop")
        break
	}
	
	fmt.Println("**********************")

	
	for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
	}
	
	fmt.Println("**********************")

	if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

}