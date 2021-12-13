package main

import (
    "os"
    "io/ioutil"
)

func main() {
    for _, i := range os.Args[1:] {
        data, err := ioutil.ReadFile(i)
        if err != nil {
            os.Stderr.WriteString(i + ": may be a directory.\n")
            os.Exit(1)
        }
        os.Stdout.Write(data)
    }

    os.Exit(0)
}
