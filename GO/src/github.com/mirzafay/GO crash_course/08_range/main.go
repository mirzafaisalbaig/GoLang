package main

import "fmt"

func main() {
	/*
		ids := []int{4, 5, 6, 12, 54, 6} //declaring and assign array

		for i, id := range ids {
			fmt.Printf("%d - ID: %d\n", i, id)
		}

		//without Index
		for _, id := range ids {
			fmt.Printf("ID: %d\n", id)
		}
	*/

	emails := map[string]string{"mirza": "mirza@gmail.com", "zak": "zaki@gmail.com", "awz": "awez@gmail.com", "yus": "yusuf@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("name :" + k)
	}

}
