//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//fmt.Println(rand.Intn(7))
	rand.New(rand.NewSource(time.Now().UnixNano()))

	diceSides := rand.Intn(12)
	diceTimes := rand.Intn(6)
	diceNum := rand.Intn(4)
	var sum int = 0

	for i := 1; i <= diceTimes; i++ {
		sum = 0

		for j := 1; j <= diceNum; j++ {
			rolled := rand.Intn(diceSides) + 1
			sum += rolled
			fmt.Println("Roll: ", i, "Die: ", j, "Rolled: ", rolled)
		}

		fmt.Printf("The total Roll is %d.\nThe number of dice is %d\nThe number of times to roll is %d\nThe rolls are %d sided\n", sum, diceNum, diceTimes, diceSides)
		switch roll := sum; {
		case roll == 2 && diceNum == 2:
			fmt.Println(`"Shake eyes"`)
		case roll == 7:
			fmt.Println(`"Lucky seven"`)
		case roll%2 == 0:
			fmt.Println(`"Even"`)
		case roll%2 != 0:
			fmt.Println(`"Odd"`)
		}

	}
	//fmt.Println(diceSides)

}
