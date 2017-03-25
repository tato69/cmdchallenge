package main

/*
 * parse challenges.json, print all keys
 * get output for keys
 * split string
 */
import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

const defaultFailedCode = 1

type expectedOutput struct {
	Lines []string `json:"lines"`
	Order bool     `json:"order"`
}

func readChallenges(cBlob []byte) (challenges, error) {
	var challengesRaw []json.RawMessage
	err := json.Unmarshal(cBlob, &challengesRaw)
	if err != nil {
		panic(err)
	}
	newChallenges := make([]challenge, 0, len(challengesRaw))
	for _, c := range challengesRaw {
		// Set json defaults
		ch := challenge{
			ExpectedOutput: expectedOutput{
				Order: true,
			},
		}
		if ch.unmarshal(c) {
			newChallenges = append(newChallenges, ch)
		}
	}
	return newChallenges, err
}

func runCombinedOutput(command string) (cmdout string, exitCode int) {
	cmd := exec.Command("bash", "-c", command)
	outb, err := cmd.CombinedOutput()
	cmdout = fmt.Sprintf("%s", outb)
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		}
	} else {
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	return cmdout, exitCode
}

func main() {

	type CmdResult struct {
		CmdOut      string
		CmdExitCode int
		Tests       bool
		TestOut     []string
		Rand        bool
		RandOut     []string
	}

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s <command>:\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	command := flag.Args()[0]

	cmdOut, cmdExitCode := runCombinedOutput(command)

	result := CmdResult{
		CmdOut:      cmdOut,
		CmdExitCode: cmdExitCode,
	}

	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(b)

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)
	fmt.Println("hi")
}
