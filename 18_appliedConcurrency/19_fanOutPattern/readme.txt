what is fan out

multiple func reading from the same channel until that channel is closed

fanIn
a func can read from multiple input (probably channel)
and proceed until all channels are closed by multiplexing
the input channels onto a SINGLE channel
that is closed
when all the inputs are closed

Pattern
there is a pattern to our pipeline func
- stages close their outbound channel when all send operations are done

- stages keep receiving values from inbound channels until those channels are closed
