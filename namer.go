package main

import "strconv"
import "os"
import "fmt"
import "time"
import "math/rand"

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
    letterIdxBits = 6                    // 6 bits to represent a letter index
    letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
    letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytesMaskImprSrc(n int) string {
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return string(b)
}

func usage() {
    fmt.Println("usage: namer <bit length> <number of names>")
}

func main() {
    args := os.Args[1:]
    b, err := strconv.Atoi(args[0])
    if err != nil {
        usage()
        os.Exit(-1)
    }
    n, err := strconv.Atoi(args[1])
    if err != nil {
        usage()
        os.Exit(-1)
    }

    names := make([]string, n)
    for i := 0; i < n; i++ {
        names[i] = RandStringBytesMaskImprSrc(b)
    }
    
    //fmt.Println(names);
}
