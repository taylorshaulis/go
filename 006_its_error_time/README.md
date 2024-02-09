# go
Now its error time
https://go.dev/doc/tutorial/handle-errors

The first change is importing more than one module

We were doing this...

`import "fmt"`

but now we are doing this

`
import (
    "errors"
    "fmt"
)
`

So instead of just a single string we are giving the import function a list of strings. Is list right? Is this a go array? idk will figure that out later.

Next we are adding on to the functions return type declaration an 'error'

>func Hello(name string) (string, error) {

So now the Hello func returns an object thats (string, error)

So the if name == "" is obvious. If the name string is blank then that if block fires and returns "" which technically is a tring and an error with the message.

Whats real important is the return at the end of the function is now returning message which is a string and nil. So nil takes the place of error.


The code is now returning an error.