package examples

import (
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

// Copied from here:
// https://stackoverflow.com/questions/20437336/how-to-execute-system-command-in-golang-with-unknown-arguments
func commandExample2(cmd string, wg *sync.WaitGroup) {
	fmt.Println("Command Is:", cmd)
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("%s", out)
	wg.Done()
}

func commandExample1() {
	fmt.Println("Executing...")
	command := exec.Command("echo", "'Hello, World!'")
	output, _ := command.Output()
	fmt.Println("Output:", output)
}
