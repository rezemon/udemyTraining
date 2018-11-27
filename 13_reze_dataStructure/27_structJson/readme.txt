
marshal is from STRUCT to JSON in form of string
encode is from STRUCT to JSON in form of stream
above should be correct, i think

==
definition of JSON

encoding
we talk about marshal, unmarshal
marshal (assemble and arrange (a group of people, especially troops) in order.)
make JSON(struct) to a string is called marshal
make string to a JSON(struct) is called unmarshal

encode and decode is for stream
stream of data,
decode is from stream to JSON (some API send me some JSON over the web , and i want to convert it back to stream )

Marshal
 - basic
 - exported fields
  - tags

primary data format used for asynchronous browser communication

in
https://godoc.org/encoding/json
i saw that there is
encode type
decode type
marshal type
unmarshal type ??
also see
func marshal
func unmarsal
==
28nov2018 0058am
we talk about JSON
because in golang
we represent JSON using struct
how cool is that

the struct explained is vvv simple
and that is what golang is about
keeping it simple and powerful !!!
