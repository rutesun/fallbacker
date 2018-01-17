package fallbacker

import "fmt"

type HookFunc func()
type HandlerFunc func() error

type Fallback struct {
	Before HookFunc
	Do     HandlerFunc
	After  HookFunc
	Retry  int
	Delay  int
}

type Fallbacker struct {
	list []*Fallback
}

func (f *Fallbacker) Add(fb *Fallback) *Fallbacker {
	f.list = append(f.list, fb)
	return f
}

func (f *Fallbacker) Count() int {
	return len(f.list)
}

func (f *Fallbacker) Do() error {
	iter := fallbacks(f.list).Iterator()
	var (
		cur *Fallback
		err error
	)
	for iter.hasNext() {
		try := -1
		cur = iter.next()
		if cur.Before != nil {
			cur.Before()
		}

		for try < cur.Retry {
			if err = cur.Do(); err != nil {
				fmt.Println(err.Error())
				try++
			} else {
				break
			}
		}

		if cur.After != nil {
			cur.After()
		}
	}
	return err
}
