package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if 4 != AddXXX(1, 3) {
		t.Error(" 1 + 3  != 4 ")
	}

}

func TestAddFlag(t *testing.T) {
	if -1 != AddXXX(-1, 3) {
		t.Error("不相同")
	}
}
