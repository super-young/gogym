Implements a command to write get an HTTP resource and write it to a file.

## Instruction

To make an HTTP Get request:

```
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
```

* [http.Response](https://golang.org/pkg/net/http/#Response)
  * Useful fields are
    * `Header map[string][]string` - The response HTTP headers.
    * `StatusCode int` - check the HTTP status code, e.g. 2XX, 3XX, 4XX.
* [http.Response.Body](https://golang.org/pkg/net/http/#Response)
  * This is an `io.Reader` to read the http response data.
  * Always remember to `Close` it.

Then write the response body to a file:

* [os.Create](https://golang.org/pkg/os/#Create)
  * This creates a file for writing.

## Example

```
$ go run wget.go https://golang.org golang.html
```

You should see the downloaded page in golang.html.