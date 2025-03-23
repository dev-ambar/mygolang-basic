package main

import frmt "fmt"

//  toy can declare the variable in the following way
// var <variable_name> <variable_type> = <value>

const pai float32 = 3.14

const log string = "LOGGER"

// if you declare  variable start from Capital letter then it will be public
// if you declare  variable start from small letter then it will be private

var LoginToken string = "123456"

func main() {

	var a int
	frmt.Println(a)
	frmt.Printf("type of the varibale is %T\n", a)

	var userName string = "Ambar Gupta"

	frmt.Println(userName)
	frmt.Printf("type of the varibale is %T\n", userName)

	var isTrue bool = true
	frmt.Println(isTrue)
	frmt.Printf("type of the varibale is %T\n", isTrue)

	var c uint = 258 // range of unit is 0 to 65535
	frmt.Println(c)
	frmt.Printf("type of the varibale is %T\n", c)

	var d uint8 = 255 // range of unit8 is 0 to 255
	frmt.Println(d)
	frmt.Printf("type of the varibale is %T\n", d)

	var e float32 = 3.14562374566666 // range of float32 is 1.18e-38 to 3.4e38
	frmt.Println(e)
	frmt.Printf("type of the varibale is %T\n", e)

	var f float64 = 3.14562374566666 // range of float64 is 2.23e-308 to 1.80e308
	frmt.Println(f)
	frmt.Printf("type of the varibale is %T\n", f)

	var g complex64 = 3.14562374566666 + 3.14562374566666i // range of complex64 is 3.4e38
	frmt.Println(g)
	frmt.Printf("type of the varibale is %T\n", g)

	// default values of the variables
	var defaultInt int
	frmt.Println(defaultInt)
	frmt.Printf("type of the varibale is %T\n", defaultInt)

	var defaultString string
	frmt.Println("value of default  of String " + defaultString)
	frmt.Printf("type of the varibale is %T\n", defaultString)

	var defaultBool bool
	frmt.Println(defaultBool)
	frmt.Printf("type of the varibale is %T\n", defaultBool)

	var defaultfloat32 float32
	frmt.Println(defaultfloat32)
	frmt.Printf("type of the varibale is %T\n", defaultfloat32)

	// defining constant values

	const constValue = 45
	frmt.Println(constValue)
	frmt.Printf("type of the varibale is %T\n", constValue)

	// constant values does not allow to change the value

	// constValue = 50 // this will give error

	// define  no type variable

	var noTypeVar = 45
	frmt.Println(noTypeVar)
	frmt.Printf("type of the varibale is %T\n", noTypeVar)

	// but we can not change the type of the variable once it assigned

	// noTypeVar = "Ambar" // this will give error

	// define novar no notype variable
	// notypevar can not be used outside the function
	// this is only for the function level scope

	novar := 45.00

	frmt.Println(novar)
	frmt.Printf("type of the varibale is %T\n", novar)

	// calling  constant value

	frmt.Println(pai)
	frmt.Println(log)

	var div int = 45
	var div1 int = 7
	frmt.Println(div/div1, div%div1)

	var div2 int = 45
	var div3 int = 0
	var err = div2 / div3
	frmt.Println(err)

}
