package main

import "strings"

// "fmt"
// "strings"
// "unicode"
// "strings"

// func Reverse(str string) string {
// 	runes := []rune(str)
// 	slices.Reverse(runes)
// 	return string(runes)
// }

// func wordCount(text string) map[string]int {
// 	mapa := map[string]int{}
// 	words := strings.Fields(text)

// 	for _, v := range words {
// 		mapa[strings.ToLower(string(v))]++
// 	}

// 	return mapa
// }

// func camelToSnake(s string) string {
// 	var build strings.Builder
// 	for index, r := range s {
// 		if unicode.IsUpper(r) {
// 			if index != 0 {
// 				build.WriteString("_")
// 			}
// 			build.WriteRune(unicode.ToLower(r))
// 		} else {
// 			build.WriteRune(r)
// 		}
// 	}
// 	return build.String()
// }

func repeat(s string, n int) string {
    var build strings.Builder
    for i := 0; i < n; i++ {
        build.WriteString(s)
    }
    return build.String()
}

func main() {
	// fmt.Println(Reverse("hello"))  // "olleh"
	// fmt.Println(Reverse("Привет")) // "тевирП"
	// fmt.Println(Reverse("Go🚀"))   // "🚀oG"

	// text := "Go is great and GO is fast"
	// fmt.Println(wordCount(text))
	// map[and:1 fast:1 go:2 great:1 is:2]

	// fmt.Println(camelToSnake("HelloWorld"))  // "hello_world"
	// fmt.Println(camelToSnake("UserName"))    // "user_name"
	// fmt.Println(camelToSnake("maxRetries"))  // "max_retries"
}