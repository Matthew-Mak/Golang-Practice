package main

var taskDone bool //fine in global scope

func main() {
	var mv int        //max value
	var max_value int //not OK
	_, _ = mv, max_value

	var packetsReceived int //NOT OK, name too long
	var b int               //OK
	_, _ = packetsReceived, b

	const MAX_VALUE = 100 //NOT OK
	const N = 100         //Better

	var Max = 100 //can be used in other packages
	var max = 100 //local?
	_, _ = max, Max

	maxValue := 1000   //recommended
	max_Values := 1000 //Not recommended
	_, _ = maxValue, max_Values

	writeToDB := true //ok, idiomatic
	writeToDb := true //not ok
	_, _ = writeToDB, writeToDb

	// 	// use short, concise names especially in shorter scopes
	// // common names for common types:
	// var s string   //string
	// var i int      //index
	// var num int    //number
	// var msg string //message
	// var v string   //value
	// var err error  //error value
	// var done bool  //bool, has been done?

	// // use mixedCase a.k.a camelCase instead of snake_case (variables and  functions)
	// var maxValue = 100  // recommended (camelCase)
	// var max_value = 100 // not recommended (snake_case)

}
