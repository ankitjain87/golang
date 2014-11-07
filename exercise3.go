//Exercise: Maps
package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    st := strings.Split(s, " ")
    m := make(map[string]int)
    for i := 0; i < len(st); i++ {
        if v, ok := m[st[i]]; ok {
            m[st[i]] = v + 1
        } else {
            m[st[i]] = 1
        }
    }
    return m
}

func main() {
    wc.Test(WordCount)
}
