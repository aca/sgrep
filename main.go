package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var fs *flag.FlagSet

type filterFunc func(string) string

var argMap map[string]filterFunc

func init() {
	argMap = make(map[string]filterFunc)
	argMap["hostname"] = filterHostname
	argMap["ipv4"] = filterIPv4
	argMap["ip"] = filterIPv4
	argMap["email"] = filterEmail
	argMap["url"] = filterHTTP
	argMap["http"] = filterHTTP
	argMap["https"] = filterHTTP
}

func main() {
	fs = flag.NewFlagSet("root", flag.ExitOnError)
	fs.Parse(os.Args)

	var filter filterFunc

	filter, ok := argMap[strings.ToLower(fs.Arg(1))]
	if !ok {
		for k, _ := range argMap {
			fmt.Println(k)
		}

		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if s := filter(line); s != "" {
			fmt.Println(s)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func filterHostname(line string) string {
	match := hostnameRegexRFC952.FindAllString(line, -1)

	// filter single word hostname which is obviously not usual
	n := 0
	for _, v := range match {
		if strings.Contains(v, ".") {
			match[n] = v
			n++
		}
	}
	match = match[:n]
	return strings.Join(match, " ")
}

func filterEmail(line string) string {
	match := emailRegex.FindAllString(line, -1)
	return strings.Join(match, " ")
}

func filterIPv4(line string) string {
	match := ipv4Regex.FindAllString(line, -1)
	return strings.Join(match, " ")
}

func filterHTTP(line string) string {
	match := httpRegex.FindAllString(line, -1)
	return strings.Join(match, " ")
}
