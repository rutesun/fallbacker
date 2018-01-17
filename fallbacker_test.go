package fallbacker

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mock struct {
	err error
}

func (m *mock) Do() error { return m.err }

func TestFallbacksAdd(t *testing.T) {
	fbs := new(Fallbacker)
	fbs.Add(&Fallback{
		Do: func() error {
			fmt.Println("Do it!")
			return nil
		},
	})
	assert.Equal(t, 1, len(fbs.list))
}

func TestFallbacksDo(t *testing.T) {
	fbs := new(Fallbacker)
	fb1 := &Fallback{}
	fb1.Do = func() error {
		t.Log("First Do it!")
		return errors.New("go to next do")
	}
	fb2 := &Fallback{}
	fb2.Do = func() error {
		t.Log("Second Do it!")
		return nil
	}
	fbs.Add(fb1).Add(fb2)
	if err := fbs.Do(); err != nil {
		t.Fatal(err)
	}
}

func TestFallbacksDoWithHooks(t *testing.T) {
	fbs := new(Fallbacker)
	fb1 := &Fallback{}
	fb1.Before = func() {
		t.Log("Before of First Do it")
	}
	fb1.Do = func() error {
		t.Log("First Do it!")
		return errors.New("go to next do")
	}
	fb1.After = func() {
		t.Log("After of First Do it")
	}
	fb2 := &Fallback{}
	fb2.Do = func() error {
		t.Log("Second Do it!")
		return nil
	}
	fb2.Before = func() {
		t.Log("Before of Second Do it")
	}
	fb2.After = func() {
		t.Log("After of Secondt Do it")
	}
	fbs.Add(fb1).Add(fb2)
	if err := fbs.Do(); err != nil {
		t.Fatal(err)
	}
}
