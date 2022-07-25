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

// here i've defined all the global variables for the project, these will serve as the parameters for what can be used to generate a password
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

	// here the password is checked to ensure the password length matches the criteria

	total_char_length_without_lower_char := min_upper_char + min_special_char + min_number_char

	if total_char_length_without_lower_char >= password_length {
		fmt.Println("Please provide valid password length")
		os.Exit(1)
	}

	// Get the user input - target folder needs to be organized
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("How many passwords you want to generate? - ")
	scanner.Scan()

	number_of_passwords, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Please provide correct value for number of passwords")
		os.Exit(1)
	}

	// it generate random number
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
		//fmt.Println(specialCharSet[random])
		//fmt.Printf("%v and %T \n", random, specialCharSet[random])
		password = password + string(special_char_set[random])
	}

	// generate random upper character based on minUpperChar
	for i := 0; i < min_upper_char; i++ {
		random := rand.Intn(len(upper_char_set))
		password = password + string(upper_char_set[random])
	}

	// generate random upper character based on minNumberChar
	for i := 0; i < min_number_char; i++ {
		random := rand.Intn(len(number_char_set))
		password = password + string(number_char_set[random])
	}

	// find remaining lowerChar
	totalCharLenWithoutLowerChar := min_upper_char + min_special_char + min_number_char

	remainingCharLen := password_length - totalCharLenWithoutLowerChar

	// generate random lower character based on remainingCharLen
	for i := 0; i < remainingCharLen; i++ {
		random := rand.Intn(len(lower_char_set))
		password = password + string(lower_char_set[random])
	}

	// shuffle the password string

	passwordRune := []rune(password)
	rand.Shuffle(len(passwordRune), func(i, j int) {
		passwordRune[i], passwordRune[j] = passwordRune[j], passwordRune[i]
	})

	password = string(passwordRune)
	return password
}
