package myfuzz
import "github.com/stellar/go/strkey"

func Fuzz(data []byte) int {
	strkey.Decode(strkey.VersionByteAccountID, string(data))
	return 0
}
