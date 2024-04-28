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

### Flags

`-l` - Start listener (default false)  
`-u` - Use UDP (default false - uses TCP)  
`--port`, `-p` - Port to start listener on (default 8080) 
