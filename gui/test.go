package gui

import (
	"math/rand"
	"strings"
)

func generateString(length int) string {
    c := []rune("qwertyuiop")
    var r strings.Builder
    
    for i := 0; i < length; i++ {
        r.WriteRune(c[rand.Intn(len(c))])
    }

    return r.String()
}
