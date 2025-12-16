/*package structs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Firstname string
	Lastname  string
	Age       int64
	Created   time.Time
}

func newUser(firstname string, lastname string, age int64) *User {
	created := time.Now()
	user := User{
		Firstname: firstname,
		Lastname:  lastname,
		Age:       age,
		Created:   created,
	}
	return &user
}

func (u *User) PrintString() {
	fmt.Printf("Firstname: %s, Lastname: %s, Age: %d, Created: %s", u.Firstname, u.Lastname, u.Age, u.Created.Format(time.RFC3339))
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	firstnametxt := Clear("Enter your name: ")
	lastnamenametxt := Clear("Enter your lastname: ")
	age := ClearNum("Enter your age: ")
	user := newUser(firstnametxt, lastnamenametxt, age)
	user.PrintString()
}

func Clear(str string) (enteredstring string) {
	fmt.Print(str)
	enteredstring, _ = reader.ReadString('\n')
	enteredstring = strings.TrimSpace(enteredstring)
	return
}

func ClearNum(str string) (num int64) {
	fmt.Print(str)
	enteredstring, _ := reader.ReadString('\n')
	enteredstring = strings.TrimSpace(enteredstring)
	num, _ = strconv.ParseInt(enteredstring, 10, 64)
	return
}
*/

package mian

import "fmt"

func main() {
	fmt.Println("structs")
}
