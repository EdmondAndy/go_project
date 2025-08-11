package main

import "regexp"

func normalize(phone string) string {
	re := regexp.MustCompile(`[^0-9]`)
	return re.ReplaceAllString(phone, "")
}

// func normalize(phone string) string {
// 	// Normalize the phone number by removing non-numeric characters
// 	// we want these - 0123456789
// 	var buff bytes.Buffer
// 	for _, ch := range phone {
// 		if ch >= '0' && ch <= '9' {
// 			buff.WriteRune(ch)
// 		}
// 	}

// 	return buff.String()
// }

func main() {

}