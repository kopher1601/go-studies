package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]

A greeter application which prints the name you entered <integer> number
of times.
`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0")
	}

	return nil
}

// os.Stdin, os.Stdout 보다 io.Reader, io.Writer 인터페이스를 사용하는 것이 좋다.
// 이렇게 하면 테스트 코드를 작성하기 쉽다.
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done. \n"
	fmt.Fprintf(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You didn't enter a name")
	}

	return name, nil
}

type config struct {
	numTimes   int
	printUsage bool
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)

	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() != 0 {
		return c, errors.New("Positional arguments provided")
	}

	return c, nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	if err := validateArgs(c); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	if err := runCmd(os.Stdin, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
