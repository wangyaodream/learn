package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTimes   int
	printUsage bool
}

var errPosArgSpecified = errors.New("positional arguments specified")

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]
A greeter application which prints the name you entered <integer> number of times.
`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintln(w, usageString)
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("must specify a number greater than 0")
	}

	return nil
}

// func parseArgs(args []string) (config, error) {
// 	var numTimes int
// 	var err error
// 	c := config{}
// 	if len(args) != 1 {
// 		return c, errors.New("invalid number of arguments")
// 	}

// 	if args[0] == "-h" || args[0] == "--help" {
// 		c.printUsage = true
// 		return c, nil

// 	}

// 	numTimes, err = strconv.Atoi(args[0])
// 	if err != nil {
// 		return c, nil
// 	}

// 	c.numTimes = numTimes

// 	return c, nil
// }

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
    // 可以处理命令行应用程序接收的参数
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
    // 用于写入任何诊断或输出的消息
	fs.SetOutput(w)
	fs.IntVar(&c.numTimes, "n", 0, "Number of times to greet")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() != 0 {
		return c, errPosArgSpecified
	}

	return c, nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Fprintln(w, msg)

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("you didn't enter your name")
	}
	return name, nil
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
		fmt.Fprintln(w, msg)
	}
}

func main() {
	c, err := parseArgs(os.Stderr, os.Args[1:])
	if err != nil {
        // 使用自定义err
        if errors.Is(err, errPosArgSpecified) {
            fmt.Fprintln(os.Stdout, err)
        }
        os.Exit(1)
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

}
