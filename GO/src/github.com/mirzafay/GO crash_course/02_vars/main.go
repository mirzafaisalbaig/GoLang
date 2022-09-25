package main
import "fmt"

func main(){
	 //Datatypes in golang
/*
String
int
bool
int8 int15 int 32 int64  
uint8 uint16 uint32 uint64 uintptr
byte - alias for int8
rune - alias for int32
float32 float64
complex64 complex128
*/	

//using var 

/*
var name = "faisal"

    // without var
age := 22      //scope is local not global, if u write it outside funtion it will throw an error
fmt.Println(name,email)    //uppercase 'P' not small p. for printing

var isCool = true //boolean
isCool = false //changes to false becuz using var instead const
const isCool  = true // cannot change this becuz const

 var marks = 2.5 //bydafault it is float64 but u can change it to 32
 var marks float32 = 4.2  
*/

//shorthand use of vars
name,email := "mirza","mirza@gmail.com"
fmt.Println(name,age,email)

//typeof data
//fmt.Printf("%T",name,email) // "%T" is used for knowing data type 

}