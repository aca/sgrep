# sgrep

Simple grep. Extract url, ip address, email ... from stdin.

Install
---
```
go get -u github.com/aca/sgrep
```

Usage
---
```
» sgrep help
usage: sgrep [flags] pattern...

example:
  cat txt | sgrep hostname ipv4

pattern:
  hostname, host
  ipv4, ip
  ipv6
  email
  url, http
  num, number
  alpha
  commit
  mac, macaddress
  uuid

flag:
  -f int
        field selector, replaces "awk '{print $3}'" as "-f3"
  -s string
        seperator (default " ")
```
