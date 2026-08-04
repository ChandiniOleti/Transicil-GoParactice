package main
import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
)
func main() {
  for _, p := range []string{"Admin@123","Bhargavi@2007","Merchant@123","Ramesh@123"} {
    h, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
    if err != nil { panic(err) }
    fmt.Printf("%s|%s\n", p, string(h))
  }
}
