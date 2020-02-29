package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// seperator string
var sep string
var field int
var help string

var fs *flag.FlagSet

type filterFunc func(string) []string

var funcMap map[string]filterFunc
var cmdList [][]string

func addCommand(f filterFunc, aliases ...string) {
	for _, alias := range aliases {
		funcMap[alias] = f
	}
	cmdList = append(cmdList, aliases)
}

func init() {
	funcMap = make(map[string]filterFunc)
	addCommand(filterHostname, "hostname", "host")
	addCommand(filterIPv4, "ipv4", "ip")
	addCommand(filterIPv6, "ipv6")
	addCommand(filterEmail, "email")
	addCommand(filterHTTP, "url", "http")
	addCommand(filterNumber, "num", "number")
	addCommand(filterAlpha, "alpha")
	addCommand(filterCommit, "commit")
	addCommand(filterMacAddress, "mac", "macaddress")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fs = flag.NewFlagSet("sgrep", flag.ExitOnError)
	fs.StringVar(&sep, "s", " ", "seperator")
	fs.IntVar(&field, "f", 0, `field selector, replaces "awk '{print $3}'" as "-f3"`)
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

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(buf), "\n")
	for _, line := range lines {
		matched := []string{}

		for _, arg := range fs.Args() {
			filter := funcMap[strings.ToLower(arg)]
			matched = append(matched, filter(line)...)
		}

		if field != 0 {
			if len(matched) >= field {
				matched = append([]string{}, matched[field-1])
			} else {
				matched = []string{}
			}
		}

		if len(matched) != 0 {
			fmt.Println(strings.Join(matched, sep))
		}

	}
}

func printHelp() {
	fmt.Fprintf(os.Stderr, `usage: sgrep [flags] pattern...

example: 
  cat txt | sgrep hostname ipv4

`)

	fmt.Fprintf(os.Stderr, "pattern:\n")

	for _, cmd := range cmdList {
		s := strings.Join(cmd, ", ")

		fmt.Fprintf(os.Stderr, "  %v\n", s)
	}

	fmt.Fprintln(os.Stderr, "\nflag:")
	fs.PrintDefaults()
}
