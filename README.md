# H2C (HTTP/2 over Cleartext) PoC

## Resources

- [HTTP2 Cleartext H2C Client Via GoLang | Mailgun](https://www.mailgun.com/blog/http-2-cleartext-h2c-client-example-go/)
- [support h2c with prior knowledge by taoso · Pull Request #1398 · gin-gonic/gin](https://github.com/gin-gonic/gin/pull/1398)

## Summary

- GIN support H2C since 1.8, which is still not released (2021-04-12).
- Golang `http2` standard library does not support H2C directly, but workaround exists.

## Usage

```
go build -o client ./cmd/client
go build -o server ./cmd/server

./server
./client
```

Checks Wireshark for the HTTP 2 traffic.
