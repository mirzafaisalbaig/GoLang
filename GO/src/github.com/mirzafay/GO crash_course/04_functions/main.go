package main
import "fmt"

func greetings(name string) string{
return "hello " + name
}
func getSum(num1 , num2 int)int{ //as num1 and num2 have same type we dont haue to write the type of each seperately
	return num1 + num2
}


func main(){

	fmt.Println( getSum(3,4))
}