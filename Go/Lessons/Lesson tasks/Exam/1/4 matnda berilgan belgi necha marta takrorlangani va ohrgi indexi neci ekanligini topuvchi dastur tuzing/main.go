package main

import "fmt"

func lastIndexAndCount(s string, ch rune) (int, int) {
    count := 0
    lastIndex := -1

    for i, char := range s {
        if char == ch {
            count++
            lastIndex = i
        }
    }

    return lastIndex, count
}

func main() {
    inputString := "mirjalol"
    targetChar := 'a'
    lastIndex, count := lastIndexAndCount(inputString, targetChar)

    if lastIndex == -1 {
        fmt.Printf("Belgi '%c' topilmadi matnda.\n", targetChar)
    } else {
        fmt.Printf("Belgi '%c' matnda %d marta takrorlangan.\n", targetChar, count)
        fmt.Printf("Oxirgi marta %c belgisi %d-indexda\n", targetChar, lastIndex)
    }
}
