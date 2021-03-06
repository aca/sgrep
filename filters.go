package main

import (
	"net"
	"strings"
)

func filterHostname(line string) []string {
	match := hostnameRegexRFC952.FindAllString(line, -1)

	// filter single word hostname which is obviously not usual
	n := 0
	for _, v := range match {
		if strings.Contains(v, ".") {
			match[n] = v
			n++
		}
	}
	return match[:n]
}

func filterCommit(line string) []string {
	return commitRegex.FindAllString(line, -1)
}

func filterEmail(line string) []string {
	return emailRegex.FindAllString(line, -1)
}

func filterIPv4(line string) []string {
	match := ipv4Regex.FindAllString(line, -1)
	n := 0
	for _, v := range match {
		if ip := net.ParseIP(v); ip != nil {
			match[n] = v
			n++
		}
	}
	return match[0:n]
}

func filterIPv6(line string) []string {
	return ipv6Regex.FindAllString(line, -1)
}

func filterHTTP(line string) []string {
	return httpRegex.FindAllString(line, -1)
}

func filterNumber(line string) []string {
	return numericRegex.FindAllString(line, -1)
}

func filterAlpha(line string) []string {
	return alphaRegex.FindAllString(line, -1)
}

func filterMacAddress(line string) []string {
	return macAddressRegex.FindAllString(line, -1)
}

func filterUUIDRegex(line string) []string {
	return uUIDRegex.FindAllString(line, -1)
}
