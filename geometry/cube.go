package geometry

import (
	"errors"
	"fmt"
)

func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("this is error")
	}
}

func Hi() {
	fmt.Println("Hi in geometry")
}
