package main

import "fmt"

/*import (
	"bufio"
	"os"
)*/

func main() {
	/*for i := 0; i < 2; i++ {
		var name string
		fmt.Print("Please Input Name: ")
		// n, err := fmt.Scanf("%s\n", &name)
		n, err := fmt.Scanln(&name)
		fmt.Println(n, err, name)
	}*/

	r := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("Enter command-> ")

		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		fmt.Println(line)
	}
}
