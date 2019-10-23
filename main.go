package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	//exec.Command("ssh", "godfrey@192.168.56.101", "`echo hello from go`").Run()
	cmd := exec.Command("ssh", "user1@192.168.56.102")
	reader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(reader)
	//while scanner has something to read
	go func() {
		for scanner.Scan() {
			fmt.Println("user1> ", scanner.Text())
		}
	}()
	cmd.Start()
	cmd.Wait()

}
