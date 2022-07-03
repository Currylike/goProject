package main
import (
	"fmt"
	"github.com/golang-module/carbon"
)
func main() {
	carbon := carbon.NewCarbon()
	now := carbon.Now().ToString()
	fmt.Println("now---",now)
}
