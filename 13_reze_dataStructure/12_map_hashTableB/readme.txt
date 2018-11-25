hash table allow us to do search very very fast

wiki hash table

map keys to values
constant search time
hash table more efficent than search tress or other table lookup structure

e.g.
dictionary, millions of word, not ordered
we have to search a word that we want the meaning
the search efficiency can be improved by sorting order
but still not fast enough

another method
break up into category
word start with A (buckets for A words)
word start with S (buckets of S words )
word start with W (etc.. etc)

problem, is some letters have more words than other
eg. S bucket a lot of words
Z bucket not a lot of words
uneven distribution of words
this uneven is not good for searching

to get even distribution, we create algorithm that take a "word" and produce a "value"
the "value" will determine which bucket the "word" belong too

if the hash algorithm is good, we can then evenly distribute our words across the buckets
and the SEARCH TIME will be constant
