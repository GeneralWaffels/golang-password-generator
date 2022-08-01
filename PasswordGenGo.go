package main

// i have used the following packages for this project, through out the project i tried to minimize the amount of packages to ensure fast compilation
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// i have defined all the global variables for the project, these will serve as the parameters for what can be used to generate a password
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

	// I get the user input here as a response to a question asked by the program, this will tell us how many passwords to generate. We are also verifying the input is a number
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("How many passwords you want to generate? - ")
	scanner.Scan()

	number_of_passwords, error := strconv.Atoi(scanner.Text())

	if error != nil {
		fmt.Println("Please provide correct value for number of passwords")
		os.Exit(1)
	}

	// it generate random
	rand.Seed(time.Now().Unix())

	for i := 0; i < number_of_passwords; i++ {
		password := generate_password()
		fmt.Printf("Password %v is %v \n", i+1, password)
	}

}

func generate_password() string {

	// declare empty password variable
	password := ""

	// generate random special character based on min_special_char
	for i := 0; i < min_special_char; i++ {
		random := rand.Intn(len(special_char_set))
		password = password + string(special_char_set[random])
	}

	// generate random upper character based on min_upper_char
	for i := 0; i < min_upper_char; i++ {
		random := rand.Intn(len(upper_char_set))
		password = password + string(upper_char_set[random])
	}

	// generate random upper character based on min_number_char
	for i := 0; i < min_number_char; i++ {
		random := rand.Intn(len(number_char_set))
		password = password + string(number_char_set[random])
	}

	// find remaining lowerChar
	total_char_length_without_lower_char := min_upper_char + min_special_char + min_number_char

	remaining_char_length := password_length - total_char_length_without_lower_char

	// generate random lower character based on remaining_char_length
	for i := 0; i < remaining_char_length; i++ {
		random := rand.Intn(len(lower_char_set))
		password = password + string(lower_char_set[random])
	}

	// shuffle the password string
	password_rune := []rune(password)
	rand.Shuffle(len(password_rune), func(i, j int) {
		password_rune[i], password_rune[j] = password_rune[j], password_rune[i]
	})

	password = string(password_rune)
	return password
}
