package main

// i have used the following packages for this project, through out the project i tried to minimize the amount of packages to ensure fast compilation and execution time.
// bufio - used to read user input
// strconv - used to convert string to integer
// rand - used to generate random numbers
// time - used to generate random numbers based on the current time
// fmt - used to print to the console
// os - used to exit the program when the user does not provide a valid input
// strings - used to shuffle the password string
// unicode - used to shuffle the password string
// math - used to generate random numbers based on the current time

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// i have defined all the global variables for the project, these will serve as the parameters for what can be used to generate
var (
	lower_char_set          = "abcdefghijklmnopqrstuvwxyz"
	upper_char_set          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_char_set string = "!@#$%^&*()-"
	number_char_set         = "123567890"
	min_special_char        = 1
	min_upper_char          = 1
	min_number_char         = 1
	password_length         = 20
)

func main() {

	// I get the user input here as a response to a question asked by the program, this will tell us how many passwords to generate. We are also verifying the input is a number and not a string or something else that is not a number and if it is not a number we will exit the program and ask the user to input a number
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("How many passwords do you want to generate? - ")
	scanner.Scan()

	number_of_passwords, error := strconv.Atoi(scanner.Text())

	if error != nil {
		fmt.Println("Please provide correct value for number of passwords")
		os.Exit(1)
	}

	// Here I generate a random number this is used later in the program to generate a random number between 1 and the number of passwords to generate and this will be used to generate the passwords and print them to the console
	rand.Seed(time.Now().Unix())

	for i := 0; i < number_of_passwords; i++ {
		password := generate_password()
		fmt.Printf("Password %v is %v \n", i+1, password)
	}

}

func generate_password() string {

	// Here I declare password as an empty variable and I will use this to store the password that is generated by the program and return this password to the main function
	password := ""

	// Here I generate a random special character based on the specification in the min_special_char variable above and add it to the password variable
	for i := 0; i < min_special_char; i++ {
		random := rand.Intn(len(special_char_set))
		password = password + string(special_char_set[random])
	}

	// Here I generate a random upper case character based on min_upper_char variable above and add it to the password variable
	for i := 0; i < min_upper_char; i++ {
		random := rand.Intn(len(upper_char_set))
		password = password + string(upper_char_set[random])
	}

	// Here i generate a random number character based on min_number_char variable above and add it to the password variable
	for i := 0; i < min_number_char; i++ {
		random := rand.Intn(len(number_char_set))
		password = password + string(number_char_set[random])
	}

	// Here i find the remaining lower_char set length and then generate a random lower case character based on this length and add it to the password variable
	total_char_length_without_lower_char := min_upper_char + min_special_char + min_number_char

	remaining_char_length := password_length - total_char_length_without_lower_char

	// Here I generate a random lower case character based on remaining_char_length variable above and add it to the password variable and then I shuffle the password string using the shuffle_password function below and return the password to the main function
	for i := 0; i < remaining_char_length; i++ {
		random := rand.Intn(len(lower_char_set))
		password = password + string(lower_char_set[random])
	}

	password_rune := []rune(password)
	rand.Shuffle(len(password_rune), func(i, j int) {
		password_rune[i], password_rune[j] = password_rune[j], password_rune[i]
	})

	password = string(password_rune)
	return password
}