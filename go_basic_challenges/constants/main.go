package main

func main() {
	// //Exercise #1
	// const daysWeek = 7
	// const lightSpeed = 299792458
	// const pi = 3.14159

	// // Exercise #2
	// const(
	// 	constdaysWeek = 7
	// 	lightSpeed    = 299792458
	// 	pi            = 3.14159
	// )

	// //Exercise #3
	// const secPerDay = 24 * 60 * 60
	// const daysYear = 365
	// println(daysYear * secPerDay)

	// // Exercise #4
	// const x int = 10

	// // declaring a constant of type slice int ([]int)
	// var m = []int{1: 3, 4: 5, 6: 8}
	// _ = m

	//Exercise #5
	const a = 7
	const b float64 = 5.6
	const c = a * b

	const x = 8
	const xc int = x

	//const noIPv6 = math.Pow(2, 128)
	const (
		Jun = 6 + iota
		Jul
		Aug
	)
	println(Jun, Jul, Aug)
}
