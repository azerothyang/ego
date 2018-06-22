package test

import (
	"testing"
	"util"
	"fmt"
)

func TestOnlyCols(t *testing.T)  {
	form := map[string]string{
		"code":"",
		"msg":"",
		"phone":"1832085ddow",
	}
	cols := []string{
		"phone",
		"msg",
	}
	util.OnlyCols(cols, &form)
	fmt.Println(form)
}

func BenchmarkOnlyCols(b *testing.B)  {
	form := map[string]string{
		"code":"",
		"msg":"",
		"phone":"1832085ddow",
	}
	cols := []string{
		"phone",
		"msg",
	}
	util.OnlyCols(cols, &form)
	fmt.Println(form)
}
