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
	scriptsListArr, err := arrScript("script.txt")
	if err != nil {
		log.Fatalf("arr script error %v \n", err)
	}
	outputChan := cmdExecute(scriptsListArr)
	command(outputChan, len(scriptsListArr))
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
func cmdExecute(arrCmd []string) chan string {
	cmdChannels := make(chan string, len(arrCmd))
	for postion, value := range arrCmd {
		if len(value) < 1 {
			oneCommand := exec.Command(value)
			// fmt.Printf("executing the command number %v , %v \n", postion+1, value)
			if oneCommand.Err != nil {
				fmt.Printf("error during command execution %v", oneCommand.Err)
			}
			output, err := oneCommand.CombinedOutput()
			if err != nil {
				fmt.Printf("output error %v 1 \n", err)
			}
			cmdChannels <- fmt.Sprintf("executing the command number %v , %v \n the result %v \n", postion+1, value, string(output))

		} else {
			// for command with aruments
			cmd := strings.Split(value, " ")
			argsCmd := cmd[:0]
			multipleCommand := exec.Command(cmd[0], argsCmd...)
			if multipleCommand.Err != nil {
				fmt.Printf("error during command execution %v", multipleCommand.Err)
			}
			output, err := multipleCommand.CombinedOutput()
			if err != nil {
				fmt.Printf("output error %v 2 \n", err)
			}
			cmdChannels <- fmt.Sprintf("executing the command number %v , %v \n the result %v \n", postion+1, value, string(output))
		}
	}
	return cmdChannels
}
func command(outputsChan chan string, scriptSize int) {
	for i := 0; i < scriptSize; i++ {
		output := <-outputsChan
		fmt.Println(output)
	}
}
