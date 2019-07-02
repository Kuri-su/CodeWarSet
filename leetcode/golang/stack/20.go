package stack

import "container/list"

func isValid_20(s string) bool {
	list := list.New()
	cMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}
	for _, value := range s {
		v := string(value)
		if list.Back() == nil {
			list.PushBack(v)
			continue
		}
		if v == cMap[list.Back().Value.(string)] {
			list.Remove(list.Back())
		} else {
			list.PushBack(v)
		}
	}

	return list.Len() == 0
}

//TODO 更好的方法
