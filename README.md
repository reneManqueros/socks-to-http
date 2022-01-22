# SOCKS5 to HTTP proxy
#### Tool to convert a SOCKS5 proxy to HTTP
 
## Installation
Download:
```shell
git clone https://github.com/reneManqueros/socks-to-http
```

Compile:
```shell
cd socks-to-http && go build .
````

## Usage

Basic usage:
```shell
./sockstohttp run
```

## Parameters 
Custom socks address - Default ":8081"
```shell
./sockstohttp run socks=127.0.0.1:8080
```

Custom listen address - Default ":8083":
```shell
./sockstohttp run listen=127.0.0.1:9090
```

Custom timeout (seconds) - Default "10":
```shell
./sockstohttp run timeout=20
```

Parameters can be combined:
```shell
./sockstohttp run timeout=20 listen=127.0.0.1:9090 socks=127.0.0.1:8080
```

 

