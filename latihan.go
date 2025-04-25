package main
import "fmt"

func main(){
	var n, hasil int
	fmt.Scan(&n)
	hasil =	fibo(n)
	fmt.Print(hasil)
}
func fibo(n int)int{
	if n == 1{
		return 1
	}else if n == 2{
		return 1
	}else{
		return fibo(n-1)+fibo(n-2)
	}
}