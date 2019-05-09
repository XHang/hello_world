package mode

import "fmt"

type FileWrite struct {
}

func (write FileWrite) WriteByte(b byte) {
	fmt.Println(b)
}
func (write FileWrite) Fresh() {

}
