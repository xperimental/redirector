# redirector

A small HTTP service that serves redirects.

## Usage

If you have a working Go installation you can get the binary by doing:

```bash
$ go get github.com/xperimental/redirector
```

Then you can run redirector:

```
$ redirector -addr <listenaddress> -destination <location>
```

You can also use environment parameters to configure the `PORT` and `DESTINATION` (for use with something like Heroku).

## Destination formatting

The destination location supports formatting. While generating the response it is passed the [Request struct](https://golang.org/pkg/net/http/#Request) that was used for the request, so all properties of that struct is available, so you can, for example, persist paths across redirects:

```bash
$ redirector -destination "https://some.example.com{{.URL}}"
```
