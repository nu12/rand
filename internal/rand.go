package rand

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
)

func getRandomInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		fmt.Println("Error generating random number: ", err)
		os.Exit(1)
	}
	return int(n.Int64())
}
