# gNmap (go Network Mapper)

A tiny Network mapper for go learning purposes and fun! It's obviously inspired by [Nmap](https://nmap.org/). My original goal is only to support a couple of host discovery and port scanning techniques (See [TODO](#TODO))

## Requirements

- Go >= 1.17

## Build and Usage

Running `make build` will generate the binary `bin/gnmap`

```bash
bin/gnmap --help
go Network Mapper

Usage:
   [flags]

Flags:
  -h, --help          help for this command
      --ips ipSlice   Target IPs (default [127.0.0.1])
      --ports ints    Target ports (default [21,22,23,25,80,110,139,443,445,3389])
      --scan string   Scan type (tcp or udp) (default "tcp")
```

Running a basic scan:

```bash

bin/gnmap --ips 127.0.0.1,192.168.0.1 --ports 21,22
gNmap (pablox.io)

scan report for 127.0.0.1

PORT -- STATE
21/tcp -- closed
22/tcp -- closed

scan report for 192.168.0.1

PORT -- STATE
21/tcp -- closed
22/tcp -- open

gNmap done: 2 IP addresses scanned in 9.1145ms
```

## Tests

```bash
make test
go test -cover ./...
?       github.io/pabloxio/gnmap        [no test files]
?       github.io/pabloxio/gnmap/cmd    [no test files]
ok      github.io/pabloxio/gnmap/pkg/mapper     0.275s  coverage: 91.4% of statements
```

## TODO

[Host discovery](https://nmap.org/book/man-host-discovery.html) support
  - [ ] TCP SYN/ACK,
  - [ ] UDP,
  - [ ] ICMP probes
  - [ ] ARP

[Scanning techniques](https://nmap.org/book/man-port-scanning-techniques.html)
  - [x] TCP
  - [ ] UDP

Etc
- [ ] Multiple targets support
- [ ] JSON Output
- [ ] Concurrency support
