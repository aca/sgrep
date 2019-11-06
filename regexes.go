package main

import "regexp"

const (
	ipv4RegexString = `([0-9]{0,3}\.){3}[0-9]{0,3}`

	// https://stackoverflow.com/questions/53497/regular-expression-that-matches-valid-ipv6-addresses
	ipv6RegexString = `(fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|2[0-4][0-9]|1{0,1}[0-9]{0,1}[0-9])\.{3,3})(25[0-5]|2[0-4][0-9]|1{0,1}[0-9]{0,1}[0-9])|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|2[0-4][0-9]|1{0,1}[0-9]{0,1}[0-9])\.{3,3})(25[0-5]|2[0-4][0-9]|1{0,1}[0-9]{0,1}[0-9])|:((:[0-9a-fA-F]{1,4}){1,7}|:))`

	// https://stackoverflow.com/questions/3809401/what-is-a-good-regular-expression-to-match-a-url
	httpRegexString = `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

	// git commit (40 length SHA1 sum)
	commitRegexString = `\b([a-f0-9]){40}\b`

	// macAddressRegexString = `(?:[[:xdigit:]]{2}([-:]))(?:[[:xdigit:]]{2}\\1){4}[[:xdigit:]]{2}`
	macAddressRegexString = "([[:xdigit:]]{2}[:-]){5}([[:xdigit:]]{2})"
	// macAddressRegexString = `([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})`

	// based on https://github.com/go-playground/validator/blob/v9/regexes.go
	alphaRegexString                 = "[a-zA-Z]+"
	alphaNumericRegexString          = "[a-zA-Z0-9]+"
	alphaUnicodeRegexString          = "[\\p{L}]+"
	alphaUnicodeNumericRegexString   = "[\\p{L}\\p{N}]+"
	numericRegexString               = "[-+]?[0-9]+(?:\\.[0-9]+)?"
	numberRegexString                = "[0-9]+"
	hexadecimalRegexString           = "[0-9a-fA-F]+"
	hexcolorRegexString              = "#(?:[0-9a-fA-F]{3}|[0-9a-fA-F]{6})"
	rgbRegexString                   = "rgb\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*\\)"
	rgbaRegexString                  = "rgba\\(\\s*(?:(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])|(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%\\s*,\\s*(?:0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)"
	hslRegexString                   = "hsl\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*\\)"
	hslaRegexString                  = "hsla\\(\\s*(?:0|[1-9]\\d?|[12]\\d\\d|3[0-5]\\d|360)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0|[1-9]\\d?|100)%)\\s*,\\s*(?:(?:0.[1-9]*)|[01])\\s*\\)"
	emailRegexString                 = "(?:(?:(?:(?:[a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(?:\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|(?:(?:\\x22)(?:(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(?:\\x20|\\x09)+)?(?:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(?:(?:(?:\\x20|\\x09)*(?:\\x0d\\x0a))?(\\x20|\\x09)+)?(?:\\x22))))@(?:(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(?:(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])(?:[a-zA-Z]|\\d|-|\\.|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*(?:[a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?"
	base64RegexString                = "(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})"
	base64URLRegexString             = "(?:[A-Za-z0-9-_]{4})*(?:[A-Za-z0-9-_]{2}==|[A-Za-z0-9-_]{3}=|[A-Za-z0-9-_]{4})"
	iSBN10RegexString                = "(?:[0-9]{9}X|[0-9]{10})"
	iSBN13RegexString                = "(?:(?:97(?:8|9))[0-9]{10})"
	uUID3RegexString                 = "[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}"
	uUID4RegexString                 = "[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}"
	uUID5RegexString                 = "[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}"
	uUIDRegexString                  = "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"
	uUID3RFC4122RegexString          = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
	uUID4RFC4122RegexString          = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}"
	uUID5RFC4122RegexString          = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}"
	uUIDRFC4122RegexString           = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
	aSCIIRegexString                 = "[\x00-\x7F]*"
	printableASCIIRegexString        = "[\x20-\x7E]*"
	dataURIRegexString               = "data:.+\\/(.+);base64"
	latitudeRegexString              = "[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)"
	longitudeRegexString             = "[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)"
	sSNRegexString                   = `[0-9]{3}[ -]?(0[1-9]|[1-9][0-9])[ -]?([1-9][0-9]{3}|[0-9][1-9][0-9]{2}|[0-9]{2}[1-9][0-9]|[0-9]{3}[1-9])`
	hostnameRegexStringRFC952        = `[a-zA-Z][a-zA-Z0-9\-\.]+[a-zA-Z0-9]`    // https://tools.ietf.org/html/rfc952
	hostnameRegexStringRFC1123       = `[a-zA-Z0-9][a-zA-Z0-9\-\.]+[a-zA-Z0-9]` // accepts hostname starting with a digit https://tools.ietf.org/html/rfc1123
	btcAddressRegexString            = `[13][a-km-zA-HJ-NP-Z1-9]{25,34}`        // bitcoin address
	btcAddressUpperRegexStringBech32 = `BC1[02-9AC-HJ-NP-Z]{7,76}`              // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	btcAddressLowerRegexStringBech32 = `bc1[02-9ac-hj-np-z]{7,76}`              // bitcoin bech32 address https://en.bitcoin.it/wiki/Bech32
	ethAddressRegexString            = `0x[0-9a-fA-F]{40}`
	ethAddressUpperRegexString       = `0x[0-9A-F]{40}`
	ethAddressLowerRegexString       = `0x[0-9a-f]{40}`
)

var (
	ipv4Regex = regexp.MustCompile(ipv4RegexString)
	ipv6Regex = regexp.MustCompile(ipv6RegexString)
	httpRegex = regexp.MustCompile(httpRegexString)
	commitRegex = regexp.MustCompile(commitRegexString)
	macAddressRegex = regexp.MustCompile(macAddressRegexString)

	alphaRegex                 = regexp.MustCompile(alphaRegexString)
	alphaNumericRegex          = regexp.MustCompile(alphaNumericRegexString)
	alphaUnicodeRegex          = regexp.MustCompile(alphaUnicodeRegexString)
	alphaUnicodeNumericRegex   = regexp.MustCompile(alphaUnicodeNumericRegexString)
	numericRegex               = regexp.MustCompile(numericRegexString)
	numberRegex                = regexp.MustCompile(numberRegexString)
	hexadecimalRegex           = regexp.MustCompile(hexadecimalRegexString)
	hexcolorRegex              = regexp.MustCompile(hexcolorRegexString)
	rgbRegex                   = regexp.MustCompile(rgbRegexString)
	rgbaRegex                  = regexp.MustCompile(rgbaRegexString)
	hslRegex                   = regexp.MustCompile(hslRegexString)
	hslaRegex                  = regexp.MustCompile(hslaRegexString)
	emailRegex                 = regexp.MustCompile(emailRegexString)
	base64Regex                = regexp.MustCompile(base64RegexString)
	base64URLRegex             = regexp.MustCompile(base64URLRegexString)
	iSBN10Regex                = regexp.MustCompile(iSBN10RegexString)
	iSBN13Regex                = regexp.MustCompile(iSBN13RegexString)
	uUID3Regex                 = regexp.MustCompile(uUID3RegexString)
	uUID4Regex                 = regexp.MustCompile(uUID4RegexString)
	uUID5Regex                 = regexp.MustCompile(uUID5RegexString)
	uUIDRegex                  = regexp.MustCompile(uUIDRegexString)
	uUID3RFC4122Regex          = regexp.MustCompile(uUID3RFC4122RegexString)
	uUID4RFC4122Regex          = regexp.MustCompile(uUID4RFC4122RegexString)
	uUID5RFC4122Regex          = regexp.MustCompile(uUID5RFC4122RegexString)
	uUIDRFC4122Regex           = regexp.MustCompile(uUIDRFC4122RegexString)
	aSCIIRegex                 = regexp.MustCompile(aSCIIRegexString)
	printableASCIIRegex        = regexp.MustCompile(printableASCIIRegexString)
	dataURIRegex               = regexp.MustCompile(dataURIRegexString)
	latitudeRegex              = regexp.MustCompile(latitudeRegexString)
	longitudeRegex             = regexp.MustCompile(longitudeRegexString)
	sSNRegex                   = regexp.MustCompile(sSNRegexString)
	hostnameRegexRFC952        = regexp.MustCompile(hostnameRegexStringRFC952)
	hostnameRegexRFC1123       = regexp.MustCompile(hostnameRegexStringRFC1123)
	btcAddressRegex            = regexp.MustCompile(btcAddressRegexString)
	btcUpperAddressRegexBech32 = regexp.MustCompile(btcAddressUpperRegexStringBech32)
	btcLowerAddressRegexBech32 = regexp.MustCompile(btcAddressLowerRegexStringBech32)
	ethAddressRegex            = regexp.MustCompile(ethAddressRegexString)
	ethaddressRegexUpper       = regexp.MustCompile(ethAddressUpperRegexString)
	ethAddressRegexLower       = regexp.MustCompile(ethAddressLowerRegexString)
	
)
