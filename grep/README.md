Implement a command similar to `grep`.

```
$ grep -n type bufio.go scan.go
bufio.go:31:type Reader struct {
bufio.go:520:type Writer struct {
bufio.go:726:type ReadWriter struct {
scan.go:18:// defined by a split function of type SplitFunc; the default split
scan.go:30:type Scanner struct {
scan.go:60:type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
```

## Instruction

You'll use three new packages in this exercise:

* [flag](https://golang.org/pkg/flag/)
  * Parse command line options
* [bufio](https://golang.org/pkg/bufio/)
  * Buffered I/O to make Read/Write more efficient.
  * Also provides additional features (like processing lines)
* [regexp](https://golang.org/pkg/regexp/)
  * Regular expression

### Regexp Matching

```
func main() {
	re, _ := regexp.Compile("^a+b*$")

	fmt.Println(
		re.MatchString("ab"),
		re.MatchString("a"),
		re.MatchString("abbb"),
		re.MatchString("aaaaaabbb"),
		re.MatchString("ba"),
	)

	// true true true true false
}
```

You'll also need to support case-insensitive matching. See:

* [Case-Insensitive Matching](http://stackoverflow.com/questions/15326421/how-do-i-do-a-case-insensitive-regular-expression-in-go)

### Process Reader By Lines

See: https://golang.org/pkg/bufio/#example_Scanner_lines

### Command Line Parsing

Use [flag](https://golang.org/pkg/flag) to parse command line arguments.

Here's the skeleton of the `grep` command we want to write:

```
var invertMatch = flag.Bool("v", false, "Selected lines are those not matching any of the specified patterns.")
var caseInsensitive = flag.Bool("i", false, "Perform case insensitive matching.")
var showLineNumber = flag.Bool("n", false, "Each output line is preceded by its relative line number in the file, starting at line 1.")

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Println("inverMatch", *invertMatch)
	fmt.Println("caseInsensitive", *caseInsensitive)
	fmt.Println("showLineNumber", *showLineNumber)
	fmt.Println("args", args)
}
```

Try to run the command with default option values:

```
$ go run example/flag/flag.go
inverMatch false
caseInsensitive false
showLineNumber false
args []
```

Try running different variations:

```
$ go run example/flag/flag.go -n
$ go run example/flag/flag.go -n -v
$ go run example/flag/flag.go -n -v -i
$ go run example/flag/flag.go -v -n -i
$ go run example/flag/flag.go -v -n -i a b c
$ go run example/flag/flag.go a b c -v -n -i
```

The help text is generated automatically:

```
$ go run example/flag/flag.go --help
Usage of /var/folders/q0/z_z6bcsx2bb9_c176v_zpt4h0000gq/T/go-build057103080/command-line-arguments/_obj/exe/flag:
  -i    Perform case insensitive matching.
  -n    Each output line is preceded by its relative line number in the file, starting at line 1.
  -v    Selected lines are those not matching any of the specified patterns.
exit status 2
```

## Examples

```
$ cat scan.go | go run grep.go var
var (
var ErrFinalToken = errors.New("final token")
var errorRune = []byte(string(utf8.RuneError))
                var r rune
                var r rune
```

```
cat scan.go | go run grep.go '^func'
func NewScanner(r io.Reader) *Scanner {
func (s *Scanner) Err() error {
func (s *Scanner) Bytes() []byte {
func (s *Scanner) Text() string {
func (s *Scanner) Scan() bool {
func (s *Scanner) advance(n int) bool {
func (s *Scanner) setErr(err error) {
func (s *Scanner) Buffer(buf []byte, max int) {
func (s *Scanner) Split(split SplitFunc) {
func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
func dropCR(data []byte) []byte {
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
func isSpace(r rune) bool {
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
```

Show line numbers of matching lines:

```
cat scan.go | go run grep.go -n '^func'
81:func NewScanner(r io.Reader) *Scanner {
90:func (s *Scanner) Err() error {
100:func (s *Scanner) Bytes() []byte {
106:func (s *Scanner) Text() string {
128:func (s *Scanner) Scan() bool {
227:func (s *Scanner) advance(n int) bool {
241:func (s *Scanner) setErr(err error) {
256:func (s *Scanner) Buffer(buf []byte, max int) {
268:func (s *Scanner) Split(split SplitFunc) {
278:func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
293:func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
326:func dropCR(data []byte) []byte {
339:func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
358:func isSpace(r rune) bool {
384:func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
```

Case-insensitive regex matching:

```
$ cat scan.go | go run grep.go -n -i 'eof'
24:// Scanning stops unrecoverably at EOF, the first I/O error, or a token too
46:// data and a flag, atEOF, that reports whether the Reader has no more data
57:// The function is never called with an empty data slice unless atEOF
58:// is true. If atEOF is true, however, data may be non-empty and,
60:type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
89:// Err returns the first non-EOF error that was encountered by the Scanner.
91:     if s.err == io.EOF {
124:// occurred during scanning, except that if it was io.EOF, Err
157:                                    // Returning tokens not advancing input at EOF.
167:            // If we've already hit EOF or an I/O error, we are done.
242:    if s.err == nil || s.err == io.EOF {
278:func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error) {
279:    if atEOF && len(data) == 0 {
293:func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error) {
294:    if atEOF && len(data) == 0 {
314:    if !atEOF && !utf8.FullRune(data) {
339:func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
340:    if atEOF && len(data) == 0 {
347:    // If we're at EOF, we have a final, non-terminated line. Return it.
348:    if atEOF {
384:func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
402:    // If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
403:    if atEOF && len(data) > start {
```

Find all the empty lines:

```
$ cat scan.go | go run grep.go -n '^\s*$'
4:
6:
13:
43:
61:
68:
75:
78:
88:
...
```

Use -v to invert match, and find all the non-empty lines:

```
$ cat scan.go | go run grep.go -n -v '^\s*$'
...
395:    for width, i := 0, start; i < len(data); i += width {
396:            var r rune
397:            r, width = utf8.DecodeRune(data[i:])
398:            if isSpace(r) {
399:                    return i + width, data[start:i], nil
400:            }
401:    }
402:    // If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
403:    if atEOF && len(data) > start {
404:            return len(data), data[start:], nil
405:    }
406:    // Request more data.
407:    return start, nil, nil
408:}
```

### Examples: Scanning Multiple Files

Instead of matching lines from stdin, grep also can search files.

Print the filename and line number of matching lines:

```
$ go run grep.go -n type scan.go bufio.go
scan.go:18:// defined by a split function of type SplitFunc; the default split
scan.go:30:type Scanner struct {
scan.go:60:type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)
bufio.go:31:type Reader struct {
bufio.go:520:type Writer struct {
bufio.go:726:type ReadWriter struct {
```