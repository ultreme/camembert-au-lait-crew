package kryptos

var (
	encryptMapping = map[string]string{
		"a": "¡",
		"b": "¢",
		"c": "£",
		"d": "¤",
		"e": "¥",
		"f": "¦",
		"g": "§",
		"h": "¨",
		"i": "ª",
		"j": "«",
		"k": "¬",
		"l": "®",
		"m": "¯",
		"n": "°",
		"o": "©",
		"p": "±",
		"q": "²",
		"r": "³",
		"s": "´",
		"t": "µ",
		"u": "¶",
		"v": "·",
		"w": "¸",
		"x": "¹",
		"y": "º",
		"z": "»",
		"A": "¼",
		"B": "½",
		"C": "¾",
		"D": "¿",
		"E": "À",
		"F": "Á",
		"G": "Â",
		"H": "Ã",
		"I": "Ä",
		"J": "Å",
		"K": "Æ",
		"L": "Ç",
		"M": "È",
		"N": "É",
		"O": "Ê",
		"P": "Ë",
		"Q": "Ì",
		"R": "Í",
		"S": "Î",
		"T": "Ï",
		"U": "Ð",
		"V": "Ñ",
		"W": "Ò",
		"X": "Ó",
		"Y": "Ô",
		"Z": "Õ",
		"0": "Ö",
		"1": "×",
		"2": "Ø",
		"3": "Ù",
		"4": "Ú",
		"5": "Û",
		"6": "Ü",
		"7": "Ý",
		"8": "Þ",
		"9": "ß",
	}

	decryptMapping = map[string]string{}
)

func init() {
	for k, v := range encryptMapping {
		decryptMapping[v] = k
	}
}

// Encrypt encrypts a readable string with the Kryptos algorithm
func Encrypt(input string) string {
	output := ""
	for _, char := range input {
		if val, ok := encryptMapping[string(char)]; ok {
			output += val
		} else {
			output += string(char)
		}
	}
	return output
}

// Decrypt decrypts a string that was encrypted with the Kryptos algorithm
func Decrypt(input string) string {
	output := ""
	for _, char := range input {
		if val, ok := decryptMapping[string(char)]; ok {
			output += val
		} else {
			output += string(char)
		}
	}
	return output
}
