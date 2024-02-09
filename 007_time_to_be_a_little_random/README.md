# go
Random time
https://go.dev/doc/tutorial/random-greeting

We add a function to greetings that returns a random string. It has no input so I am guessing there is no need for error checking.

The first thing i wanna look at is how a new array of type string is declared

`
formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
`

So it does the whole ':=' operator to declare a new variable
That variable is set as an array [] and string. So its an array of strings.  The elements of the string are set within the brackets and separated by a comma. you can have an extra comma at the end.

The next interesting part is

>return formats[rand.Intn(len(formats))]

It appears we can access the elements of the array in the standard numerical order.
formats[0] would give us "Hi, %v. Welcome!"
formats[1] would give us "Great to see you, %v!"
and formats[2] would give us "Hail, %v! Well met!"

Now looking into the intn function in the rand package.
https://pkg.go.dev/math/rand#Intn

A non negative number between 0 and X.  

len is a built in function
https://pkg.go.dev/builtin#len

I would think this is the count of elements in the formats array.

Now I am curious. Does the array element count start at 0? Does the len result in 3 given there is 3 elements? If that is the case rand would resuilt in 0,1,2,3 but the array only has 3 elements not 4.

Time to explore this and find the answer.

Ok I added the math rand to hello and was able to confirm this.
The len does give the count of the array which is 3.
Also the rand is giving values of 0,1,2

Now that I am thinking of it it makes sense. The parameter value of 3 means the rand.Intn returns a value randomly from a list of 3 starting with 0 so one element of 0,1,2

I think I have been out of the programming game too long but its coming back to me.

I commented out the testing code here but will have it removed from further steps.