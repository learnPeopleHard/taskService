package test

import (
	"fmt"
	"loginService/src/go/common"
	"testing"
)

func TestMd5(t *testing.T)  {
	fmt.Println(common.Md5V("2"))
}
