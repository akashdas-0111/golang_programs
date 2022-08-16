package main
import (
	"fmt"
	"go-tutorails/internal/math"
)
func main() {
	fmt.Println("Enter the number")
	var ele int
	fmt.Scanf("%d\n",& ele)
	result:=math.ReverseNumber(ele)
	fmt.Println(result)
}
