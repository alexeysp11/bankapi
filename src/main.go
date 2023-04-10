package main 

import "fmt" 
import "strconv"
import "bufio"
import "os"

func main() {
	var conferenceName string = "Go conference"
	fmt.Println("Welcome to", conferenceName, "banking application")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter userName: ")
	userName, _ := reader.ReadString('\n')
	fmt.Println("Username:", userName)

	var num string = "235"
	i, err := strconv.Atoi(num)
    if err != nil {
        panic(err)
    }
    fmt.Println(num, i)
}
