package fallbacker

type fallbacks []*Fallback

func (ss fallbacks) Iterator() *iterable {
	return &iterable{
		head: 0,
		data: fallbacks(ss),
	}
}

type iterable struct {
	head int
	data fallbacks
}

func (iter *iterable) next() *Fallback {
	var cur *Fallback

	if iter.hasNext() {
		cur = iter.data[iter.head]
		iter.head++
		return cur
	}
	return cur
}

func (iter *iterable) hasNext() bool {
	l := len(iter.data)
	return l > 0 && iter.head < l
}
