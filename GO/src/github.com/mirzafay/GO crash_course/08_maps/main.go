package main

import "fmt"

func main() {
	/*declaring maps

	emails:= make(map[string]string)

	  //assigning kv
	 emails["awez"] = "awez@gmail.com"
	 emails["faisal"] = "fay@gmmail.com"


	*/
	//shorthand method
	emails := map[string]string{"mirza": "mirza@gmail.com", "zak": "zaki@gmail.com", "awz": "awez@gmail.com", "yus": "yusuf@gmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["awz"])
	fmt.Println(len(emails))
	//deleting a map value
	delete(emails, "awz")
	fmt.Println(emails)
}
