package util

import (
  "math/rand"
  "time"
  "strings"
  "fmt"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
  rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomString(n int) string {
  var sb strings.Builder
  k := len(alphabet)
  for i := 0; i < n; i++ {
    c := alphabet[rand.Intn(k)]
    sb.WriteByte(c)
  }
  return sb.String()
}

func RandomName() string {
  return RandomString(6)
}

// numeric in postgres will be stored as string, float will not work with postgres numeric
func RandomNumeric() string {
  intPart := rand.Intn(1000000)
  decimalPart := rand.Intn(100)
  return fmt.Sprintf("%d.%02d", intPart, decimalPart)
}

