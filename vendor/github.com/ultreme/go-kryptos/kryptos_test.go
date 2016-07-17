package kryptos

import "fmt"

func Example() {
	fmt.Println(Decrypt(Encrypt("Merci nature d'être là !")))
	// Output: Merci nature d'être là !
}

func ExampleEncrypt() {
	fmt.Println(Encrypt("Merci nature d'être là !"))
	// Output: È¥³£ª °¡µ¶³¥ ¤'êµ³¥ ®à !
}

func ExampleDecrypt() {
	fmt.Println(Decrypt("È¥³£ª °¡µ¶³¥ ¤'êµ³¥ ®à !"))
	// Output: Merci nature d'être là !
}
