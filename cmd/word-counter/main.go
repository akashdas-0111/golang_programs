package main
import (
	"fmt"
	"go-tutorails/internal/math"
)
func main() {
	var st string
	fmt.Println("Enter a word to count number of letters")
	fmt.Scanf("%s", &st)
	res := math.Wordc(st)
	fmt.Printf("The number of letters present  is %d\n", res)
}
