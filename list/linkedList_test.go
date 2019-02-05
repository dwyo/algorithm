package linkedList

import (
	"testing"
	//"github.com/davecgh/go-spew/spew"
	"fmt"
)

func TestList_InitList(t *testing.T) {
	l := new(List)

	if l.len != 0 {
		t.Error("error")
	}
}

func TestList_Append(t *testing.T) {
	l := new(List)

	for i := 0; i < 10; i++ {
		l.Append(&Element{value: i})
	}
	if l.len != 10 {
		t.Errorf("error")
	}

	l.Each()
	//l.Reversion()
	//l.Del(2)
	fmt.Println("---------------------------------")
	l.AppendNode(&Element{value: "hellp"}, 3)
	l.Each()
	fmt.Println("---------------------------------")
	l.Reversion()

}
