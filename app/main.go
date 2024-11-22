package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	scripts, err := arrScript("script.txt")
	if err != nil {
		log.Fatalf("arr script error %v \n", err)
	}
	cmdExecute(scripts)
}
func arrScript(fileName string) ([]string, error) {
	var arrCmd []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arrCmd = append(arrCmd, scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, fmt.Errorf("scanner Error %v", scanner.Err())
	}
	return arrCmd, nil
}

// 6313 skf
func cmdExecute(arrCmd []string) {
	for postion, value := range arrCmd {
		if len(value) < 1 {
			oneCommand := exec.Command(value)
			// fmt.Printf("executing the command number %v , %v \n", postion+1, value)
			if oneCommand.Err != nil {
				fmt.Printf("error during command execution %v", oneCommand.Err)
			}
			output, err := oneCommand.Output()
			if err != nil {
				fmt.Printf("output error %v", err)
			}
			fmt.Printf("executing the command number %v , %v \n the result %v \n", postion+1, value, string(output))

		} else {
			// for command with aruments
			cmd := strings.Split(value, " ")
			argsCmd := cmd[:0]
			oneCommand := exec.Command(cmd[0], argsCmd...)
			if oneCommand.Err != nil {
				fmt.Printf("error during command execution %v", oneCommand.Err)
			}
			output, err := oneCommand.Output()
			if err != nil {
				fmt.Printf("output error %v", err)
			}
			fmt.Printf("executing the command number %v , %v \n the result %v \n", postion+1, value, string(output))
		}
	}
}
