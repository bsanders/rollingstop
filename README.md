# rollingstop
A RESTful API writen in Go to generate a random dice roll.

*rollingstop* allows you to generate arbitrary dice rolls, returned as JSON.  The API is simple, it only comes with a single endpoint, but it allows you to specify the number of sides on the die as well as a modifier and a number of times to do the roll.  The route looks like this:

/d/\<dice pattern\>

Where \<dice pattern\> is of the form \<number of sides\>\<?count=positive_number\>\<?modifier=number\>.

For example:

/d/20?modifier=2 translates to 1d20+2

## Why?

To learn Go, mostly.  But also to provide a common die roller for a few other projects (such as an IRC bot and a Reddit bot).
