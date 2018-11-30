
how to check
browse to
https://godoc.org/bufio
search for io.Reader
this is to check those func that can accept io.Reader
then you can see all the func that can take a io.Reader
for example
func NewReader
func NewReaderSize
func Reset
func NewScanner
func  ReadFrom

https://godoc.org/io/ioutil
ctrl-F search for io.Reader
func that accept io.Reader
func NopCloser
func ReadAll


===
power of interface
code substitutability
Bill Kennedy

to send (write) TO internet
from Server to client's web browser
we need a io.Writer

to receive (read) FROM internet
from e.g (external API) to my (e.g. reze) Server
we need a io.Reader

==
video 116 code substitutability
hashtable english alphabet revisited
Readall(r io.Reader)
   = ioutil.ReadAll(response.Body)
reze: above probably read from a EXTERNAL website body
NewScanner (r io.Reader)
   = bufio.NewScanner(response.Body)

so above, what is ioutil
what is scanner

NOW TO EXPLAIN
why we can pass in response.Body
is because both ask for a io.Reader
this io.Reader is an interface
this io.Reader is an interface
this io.Reader is an interface
so ANYTHING that implements an io.Reader
can be passed to ioutil.ReadAll
can be passed to bufio.NewScanner

from above
Readall(r io.Reader)
   = ioutil.ReadAll(response.Body)
reze: above probably read from a EXTERNAL website body
NewScanner (r io.Reader)
   = bufio.NewScanner(response.Body)

response.Body DEFINITELY implements io.Reader
otherwise CANNOT be passed in

so now check whether response.Body implements io.Reader or not

Readall(r io.Reader)
   = ioutil.ReadAll(response.Body)
Readall works for ANYTHING that implements io.Reader

NewScanner (r io.Reader)
   = bufio.NewScanner(response.Body)
NewScanner works for ANYTHING that implements io.Reader
