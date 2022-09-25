package main
import "fmt"

func main(){
	//Arrays declare
	var fruitArr[2]string 
	fruitArr[0] = "Apple"
	fruitArr[1] = "Grapes"

	//Decclaring and assigning

   //fruitArray := [3]string{"Lemon","Banana","Guava"}
	//fmt.Println(fruitArray)

// Arrays without specifying lenght

fruitSlice := []string{"Lemon","Banana","Guava"}
fmt.Println(fruitSlice[1:2]) //Slicing the array
fmt.Println(len(fruitSlice))// it returns the lenght of array


}