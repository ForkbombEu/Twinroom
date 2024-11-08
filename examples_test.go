package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

// Example of using the list command in a folder with slang files.
func Example_listCmd() {

	// Prepare the command
	cmd := exec.Command("go", "run", "main.go", "list", "contracts")

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	// Check for errors
	if err != nil {
		fmt.Println("Command execution failed:", err)
		return
	}

	// Output the results
	output := out.String()
	fmt.Print(output)

	// Output:
	// Listing slangroom files in folder: contracts
	// Found file: env.slang (Path: contracts/test/env.slang)
	// Found file: execute_zencode.slang (Path: contracts/test/execute_zencode.slang)
	// Found file: hello.slang (Path: contracts/test/hello.slang)
	// Found file: param.slang (Path: contracts/test/param.slang)
	// Found file: test.slang (Path: contracts/test/test.slang)
}

// Example of using the run command to execute a specific slangroom file.
func Example_runCmd() {

	// Prepare the command to run the slang file
	cmd := exec.Command("go", "run", "main.go", "test", "hello")

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	// Check for errors
	if err != nil {
		fmt.Println("Command execution failed:", err)
		return
	}

	// Output the results
	fmt.Print(out.String())

	// Output:
	// {"output":["Hello_from_embedded!"]}
}
func Example_runCmdWithExtraData() {

	// Prepare the command to run the slang file
	cmd := exec.Command("go", "run", "main.go", "test", "execute_zencode")

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	// Check for errors
	if err != nil {
		fmt.Println("Command execution failed:", err)
		return
	}

	// Output the results
	fmt.Print(out.String())

	// Output:
	// {"result":{"ecdh_public_key":"BLOYXryyAI7rPuyNbI0/1CfLFd7H/NbX+osqyQHjPR9BPK1lYSPOixZQWvFK+rkkJ+aJbYp6kii2Y3+fZ5vl2MA="}}
}

func Example_runCmdWithParam() {

	// Prepare the command to run the slang file
	cmd := exec.Command("go", "run", "main.go", "test", "param", "username", "-n", "testname", "-D", "small")

	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	// Check for errors
	if err != nil {
		fmt.Println("Command execution failed:", err)
		return
	}

	// Output the results
	fmt.Print(out.String())

	// Output:
	// {"drink":"small","name":"testname","timeout":60,"username":"username"}
}

func Example_runCmdWithEnvVariable() {

	// Prepare the command to run the slang file
	cmd := exec.Command("go", "run", "main.go", "test", "env")
	os.Setenv("FILE_CONTENT", "The enviroment variable is set correctly")
	// Capture the output
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	// Check for errors
	if err != nil {
		fmt.Println("Command execution failed:", err)
		return
	}

	// Output the results
	fmt.Print(out.String())

	// Output:
	//{"content":"The enviroment variable is set correctly"}
}
