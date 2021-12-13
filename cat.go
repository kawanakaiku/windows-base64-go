package main

import (
    "os"
    "io/ioutil"
)

func main() {
    var code int

    code = 0
    for _, i := range os.Args[1:] {
        data, err := ioutil.ReadFile(i)
        if err != nil {
            os.Stderr.WriteString(i + ": no such file or directory.\n")
            code = 1
        }
        os.Stdout.Write(data)
    }

    os.Exit(code)
}
