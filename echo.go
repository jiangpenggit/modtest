package modtest

import (
	"fmt"
)
func Echo(something string ) {
	fmt.Println("go mod test start...")
	defer fmt.Println("go mod test end!")
	fmt.Println("Echo:", something)
}
