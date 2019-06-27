package php

//map排序，按key排序
type KeySorter []ArrayItem

func Ksort(m map[string]interface{}) KeySorter {
	ms := make(KeySorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, ArrayItem{Key: k, Val: v})
	}
	return ms
}

type ArrayItem struct {
	Key string
	Val interface{}
}

func (ms KeySorter) Len() int {
	return len(ms)
}

func (ms KeySorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

//按键排序
func (ms KeySorter) Less(i, j int) bool {
	return ms[i].Key < ms[j].Key
}
