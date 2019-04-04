//Package golangexamples is the in-class assignment
package golangexamples

import (
	"github.com/ehteshamz/greetings"
)

//ConcatSlice turns a byte slice into its equivalent string representation
func ConcatSlice(sliceToConcat []byte) string {
	stringRep := ""

	for index := 0; index < len(sliceToConcat); index++ {
		stringRep = stringRep + string(sliceToConcat[index])
	}

	return stringRep
}

//Encrypt encrypts a given byte slice using ceaserCount (given as parameter)
func Encrypt(sliceToEncrypt []byte, ceaserCount int) []byte {
	EncryptedSlice := make([]byte, len(sliceToEncrypt))

	ceaserCount = ceaserCount % 256 //make ceaserCount between -255 to 255

	for index := 0; index < len(sliceToEncrypt); index++ {
		var Limit byte

		if ceaserCount < 0 {
			Limit = sliceToEncrypt[index]

			if byte(-1*ceaserCount) <= Limit {
				//No wraparound

				EncryptedSlice[index] = sliceToEncrypt[index] - byte(-1*ceaserCount)
			} else {
				//There is wraparound

				wrapOverSize := (byte(-1*ceaserCount) - Limit - byte(1)) //-1 is to accommodate the change from 255 to 0

				EncryptedSlice[index] = byte(255) - wrapOverSize
			}
		} else {
			Limit = byte(255) - sliceToEncrypt[index]

			if byte(ceaserCount) <= Limit {
				//No wraparound

				EncryptedSlice[index] = sliceToEncrypt[index] + byte(ceaserCount)
			} else {
				//There is wraparound

				wrapOverSize := (byte(ceaserCount) - Limit - 1) //-1 is to accommodate the change from 255 to 0

				EncryptedSlice[index] = wrapOverSize
			}
		}
	}

	return EncryptedSlice
}

//EZGreetings returns the greeting string using the greetings package
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}
