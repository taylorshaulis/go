# go
Multiple inputs and outputs
https://go.dev/doc/tutorial/greetings-multiple-people

First cleaned up the extra code from my test in the 007

Oh boy now we are getting into loops

Actually before that... maps? make?

Gosh I miss PHP where this stuff is just done super easy

Ok maps
https://go.dev/blog/maps

So maps are arrays in go? Good old hash tables.

Ok next cool thing... `_`

https://go.dev/doc/effective_go#blank

It's a bit like writing to the Unix /dev/null file: it represents a write-only value to be used as a place-holder where a variable is needed but the actual value is irrelevant. 

I also need to talk about slice
https://go.dev/tour/moretypes/7

I guess in go arrays are a fixed size but a slice can be dynamic.


Anyways

>for _, name := range names

So range explodes the names slice into index and key. We are gonna dev null the index and assign the current key to the new variable name.

Then we do the Hello function with the name. Test if there was an error.... go loves err checking. And assuming there was no error add the message to the messages map with the name being the index key.

Pass back the map of names with their messages and print them out.

This tutorial just prints out the names as one big map... that looks not cool. Lets fix that

Added this bit at the end of the hello.go to loop threw the messages.

`
    for name, message := range messages {
        fmt.Println(name, " - ", message)
    }
`