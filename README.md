# Ramen
Ramen is a programming language.

## Install
```
go get -u -v github.com/akanoren/ramen
```

## Sample code
Clone this repository to somewhere, and run the below command in the repository.

```
ramen examples/hello.ramen
```

## Syntax
A sequence of '🍜' represents a command of the language.
Other characters are delimiters or comments.

## Commands
There are only eight commands for the language(all commands are the same as Brainf*ck's one).

1. 🍜
    Increment the value to which the program pointer points

2. 🍜🍜
    Decrement the value to which the program pointer points

3. 🍜🍜🍜
    Increment the program pointer

4. 🍜🍜🍜🍜
    Decrement the program pointer

5. 🍜🍜🍜🍜🍜
    Write the value to which the program pointer points to Stdout

6. 🍜🍜🍜🍜🍜🍜
    Read a unicode character from Stdin and set the value to which the program pointer points to it

7. 🍜🍜🍜🍜🍜🍜🍜
    If the value to which the program pointer points is zero, jump to the next '🍜🍜🍜🍜🍜🍜🍜🍜'

8. 🍜🍜🍜🍜🍜🍜🍜🍜
    Jump to '🍜🍜🍜🍜🍜🍜🍜' which most recently appeared
