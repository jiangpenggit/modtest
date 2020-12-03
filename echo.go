package modtest

import (
	"fmt"
)
func Echo(something string ) {
	fmt.Println("Mod test start...")
	defer fmt.Println("Mod test end...")
	fmt.Println("Echo:", something)
}
