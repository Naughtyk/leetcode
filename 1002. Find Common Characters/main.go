package main

import (
	"fmt"
)

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

func commonChars(words []string) []string {
    mapa := make(map[rune]int)
    for _, ch := range words[0]{
        _, ok := mapa[ch]
        if ok{
            mapa[ch] += 1
        }else{
            mapa[ch] = 1
        }
    }
    for _, word := range words[1:]{
        mapa2 := make(map[rune]int)
        for _, ch := range word{
            _, ok := mapa2[ch]
            if ok{
                mapa2[ch] += 1
            }else{
                mapa2[ch] = 1
            }
        }
        for k, v := range mapa{
            mapa[k] = min(v, mapa2[k])
        }
    }
    ret := []string{}
    for k, v := range mapa{
        for _ = range(v){
            ret = append(ret, string(k))
        }
    }
    return ret
}

func min(v1, v2 int) int{
    if v1 < v2{
        return v1
    }
    return v2
}