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
  ifconfig wlp2s0 | sgrep mac

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

flag:
  -s string
    	seperator (default " ")
```
