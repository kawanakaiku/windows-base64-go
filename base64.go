package main

import (
    "os"
    "bufio"
    b64 "encoding/base64"
    "fmt"
    "errors"
    "io/ioutil"
)

func main() {
    var encode bool
    var file string
    var data []byte
    var err error

    encode = true
    file = ""

    for _, i := range os.Args[1:] {
        if (i == "-d") {
            encode = false
        } else {
            _, err := os.Stat(i)

            if (err == nil) {
                if (file == "") {
                    file = i
                } else {
                    os.Stderr.WriteString(i + ": more than two files given.\n")
                os.Exit(1)
                }
            } else if errors.Is(err, os.ErrNotExist) {
                os.Stderr.WriteString(i + ": not exists.\n")
                os.Exit(1)
            } else {
                os.Stderr.WriteString(i + ": error.\n")
                os.Exit(1)
            }
        }

    }

    if (file != "") {
        data, err = ioutil.ReadFile(file)
        if err != nil {
            os.Stderr.WriteString(file + ": may be a directory.\n")
            os.Exit(1)
        }
    } else {
        stdin := bufio.NewScanner(os.Stdin)
        stdin.Scan()
        data = stdin.Bytes()
    }

    fmt.Println(string(data))

    if encode {
        out := b64.StdEncoding.EncodeToString(data)
        fmt.Println(out)
    } else {
        out, _ := b64.StdEncoding.DecodeString(string(data))
        fmt.Println(string(out))
    }
    os.Exit(0)
}
