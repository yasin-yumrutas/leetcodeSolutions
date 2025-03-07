package main

import "fmt"

func romanToInt(s string) int {
    romanValues := map[byte]int{
        'I': 1, 'V': 5, 'X': 10, 'L': 50,
        'C': 100, 'D': 500, 'M': 1000,
    }

    result := 0
    length := len(s)

    for i := 0; i < length; i++ {
        if i < length-1 && romanValues[s[i]] < romanValues[s[i+1]] {
            result -= romanValues[s[i]] 
        } else {
            result += romanValues[s[i]]
        }
    }

    return result
}

func main() {
    fmt.Println(romanToInt("III"))      
    fmt.Println(romanToInt("LVIII"))    
    fmt.Println(romanToInt("MCMXCIV"))  
}
