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

var funcMap map[string]filterFunc

func init() {
	funcMap = make(map[string]filterFunc)
	funcMap["hostname"] = filterHostname
	funcMap["ipv4"] = filterIPv4
	funcMap["ip"] = filterIPv4
	funcMap["email"] = filterEmail
	funcMap["url"] = filterHTTP
	funcMap["http"] = filterHTTP
	funcMap["https"] = filterHTTP
}

func main() {
	fs = flag.NewFlagSet("root", flag.ExitOnError)
	fs.Parse(os.Args)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		matched := []string{}

		for _, arg := range fs.Args() {
			filter, ok := funcMap[strings.ToLower(arg)]
			if !ok {
				continue
			}
			if s := filter(line); s != "" {
				matched = append(matched, s)
			}
		}
		if len(matched) != 0 {
			fmt.Println(strings.Join(matched, " "))
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
