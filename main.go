package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// seperator string
var sep string
var help string

var fs *flag.FlagSet

type filterFunc func(string) []string

var funcMap map[string]filterFunc

func init() {
	funcMap = make(map[string]filterFunc)
	funcMap["hostname"] = filterHostname

	funcMap["ipv4"] = filterIPv4
	funcMap["ipv6"] = filterIPv6
	funcMap["ip"] = filterIPv4

	funcMap["email"] = filterEmail
	funcMap["url"] = filterHTTP
	funcMap["http"] = filterHTTP

	funcMap["number"] = filterNumber
	funcMap["num"] = filterNumber

	funcMap["alpha"] = filterAlpha
}

func main() {
	fs = flag.NewFlagSet("root", flag.ExitOnError)
	fs.StringVar(&sep, "s", " ", "seperator")
	fs.Parse(os.Args[1:])

	// validate arguments
	for _, arg := range fs.Args() {
		if arg == "help" {
			printHelp()
			os.Exit(0)
		}

		_, ok := funcMap[strings.ToLower(arg)]
		if !ok {
			fmt.Fprintf(os.Stderr, "invalid argument %s, check 'sgrep help'\n", arg)
			os.Exit(1)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matched := []string{}

		for _, arg := range fs.Args() {
			filter := funcMap[strings.ToLower(arg)]
			matched = append(matched, filter(line)...)
		}
		if len(matched) != 0 {
			fmt.Println(strings.Join(matched, sep))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func printHelp() {
	fmt.Fprintf(os.Stderr, `usage: sgrep [flags] pattern...

example: 
 cat txt | sgrep hostname ipv4

`)

	fmt.Fprintf(os.Stderr, "pattern:\n")

	validArg := []string{}
	for arg, _ := range funcMap {
		validArg = append(validArg, arg)
	}
	sort.Strings(validArg)
	for _, arg := range validArg {
		fmt.Fprintf(os.Stderr, " %s\n", arg)
	}

	fmt.Fprintln(os.Stderr, "\nflag:")
	fs.PrintDefaults()
}
