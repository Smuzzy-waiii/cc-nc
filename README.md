# ccnc

Netcat - which is usually abbreviated to `nc` is a command line networking utility for reading and writing to network connections with TCP or UDP.  

This is a clone of netcat following this coding challenge - https://codingchallenges.substack.com/p/coding-challenge-59-netcat

## Usage

### Examples

Build the binary
```shell
go build .
```

Start a TCP listener on port 8888
```shell
./ccnc -l -p 8888
```

Start a UDP listener on port 8888
```shell
./ccnc -l -u -p 8888
```

Try connecting to a single port
```shell
./ccnc -z localhost 8080
```

Try connecting to a range of ports
```shell
./ccnc -z localhost 8000-9000
```

### Flags

`-l` - Start listener (default false)  
`-u` - Use UDP (default false - uses TCP)  
`--port`, `-p` - Port to start listener on (default 8080)  
`-z` - Try establishing connecting with a port or range of ports (boolean flag) (Check Examples for usage)