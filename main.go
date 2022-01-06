package main
import(
  "github.com/Shyun304/go_study/banking"
  "fmt"
)
  


func main(){
  account := banking.Account{Owner: "nico", Balance: 1000}
  fmt.Println(account)
}