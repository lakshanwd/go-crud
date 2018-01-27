package handler

import "container/list"

func convertListToArray(given *list.List) *[]interface{} {
	docs := make([]interface{}, given.Len())
	index := 0
	for e := given.Front(); e != nil; e = e.Next() {
		docs[index] = e.Value
		index++
	}
	return &docs
}
