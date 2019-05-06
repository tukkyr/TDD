package my_interface

import "testing"

func TestMyInterface(t *testing.T) {
	taro := New(Male, "Taro", "Yamada")
	male_fullName := FullName(taro)
	t.Logf("%v\n", male_fullName)
}
