package snatch_test

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"

	"github.com/podhmo/snatch"
)

type S struct {
	cont func(indent int) error
}

func (s *S) Do() error {
	// do something
	return s.cont(1)
}

func New() *S {
	return &S{
		cont: func(n int) error {
			return nil
		},
	}
}
func Test(t *testing.T) {
	// {
	// 	s := New()
	// 	fmt.Println(s.Do())
	// }

	s := New()
	snatch.Field(s, "cont", func(ptr unsafe.Pointer) {
		realPtr := (*func(int) error)(ptr)
		cont := func(n int) error {
			return fmt.Errorf("%s%s", strings.Repeat("@", n*4), "<black magic>")
		}
		*realPtr = cont
	})

	err := s.Do()
	if err == nil {
		t.Fatalf("must be replaced")
	}

	want := `@@@@<black magic>`
	if want != err.Error() {
		t.Errorf("want %q, but %q", want, err)
	}
}
