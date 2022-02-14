# gNmap (go Network Mapper)

A tiny Network mapper for go learning purposes and fun!

- It's obviously inspired in Nmap, but it's not going to be 100% compatible
- The project is aimed for golang learning purposes
- My initial idea is to support a couple of host discovery/port scanning techniques (See [TODO](#TODO))

## Requirements

- Go >= 1.17

## Tests

```bash
make test
go test ./...
?       github.io/pabloxio/gnmap        [no test files]
?       github.io/pabloxio/gnmap/cmd    [no test files]
ok      github.io/pabloxio/gnmap/pkg/mapper     (cached)
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
