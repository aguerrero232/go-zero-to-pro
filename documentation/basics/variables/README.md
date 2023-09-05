# <img src="../../../favicon.ico" width="30px"> **GO** - ***Variables*** 🔡

## **Description** 👀

GO is a statically typed language, which means that the type of a variable (whether it is an integer, a floating point number, a character or a string, etc.) cannot change. Once a variable has been declared to be of a certain type, it can only store values of that type and any attempt to store any other type of value in it will result in an error. This is in contrast to dynamically typed languages like Python or JavaScript, where the type of a variable can change at runtime.

## Variables Naming Rules

1. Variable names must start with a letter (lower or uppercase) or an underscore.
2. Variables must be unicode letter characters, so you can use any language.
3. Variable names cannot start with a number.
4. Variable names cannot contain a punctuation character.
5. Variable names cannot be a Go keyword.
    - i.e. `var` is a keyword, so you cannot use it as a variable name.

## **Examples** 🧩

- declaring a variable

    ```go
    var x int 
    ```

  - **var** short for variable
  - **x** is the name of the variable, also known as the `identifier`
  - **int** is the type of the variable, once a variable is declared with a type, it cannot be changed.

- running the program

    ```go
    // documentation/basics/variables/main.go
    package main
    import "fmt"

    func main() {
        var x int
        fmt.Println(x)
    }
    ```

    ```bash
    go run main.go
    ```

    `output`: **0**

  - the output is **0** because the default value of an integer is **0**.

[↩️](../README.md)
