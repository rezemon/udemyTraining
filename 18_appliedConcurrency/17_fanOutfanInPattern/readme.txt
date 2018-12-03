fan out
fan in

fan out is
some channel
multiple funcs reading from that channel until it is closed

fan in is
some channel
multiple Channels writing to the same channel
MUTIPLE channels writing to the SAME channel
multiple channel with different result , writing to 1 channel


have data to process,
i distribute it to goRoutine to process it
that is call fan out
fan out the work

e.g.
channel of a bunch of numbers
we fan out to a group of workers, goRoutine
to calculate Factorial of the number
maybe SqRoot, maybe DoubleTheNumber, if we can distribute them
all done in parallel , on different CPU
provided we have multiple CPU

seems like could be use to do Quick sort
26 sets of database, A to Z
not sorted yet,
26 func sort A, B, C ...Z
26 func write data to 26 channels

then 26 channels CONVERT to a single channel,
WHOLE DICTIONARY IS SORTED
