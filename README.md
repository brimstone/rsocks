rsocks
======

Tiny little reverse socks5 client & server.

Usage:
------
```
rsocks - reverse socks5 server/client
https://github.com/brimstone/rsocks

Usage:
1) Start rsocks -listen :8080 -socks 127.0.0.1:1080 on the client.
2) Start rsocks -connect client:8080 on the server.
3) Connect to 127.0.0.1:1080 on the client with any socks5 client.
4) Enjoy. :]
```

This binary functions as both the client and the server. It does not encrypt
communications.
