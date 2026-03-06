package main

import "fmt"

type GoMap struct {
	arr [1000]*Entry
}
type Entry struct {
	Key   string
	Value string
	Next  *Entry
}

func (mapobj *GoMap) GetHash(key string) int32 {
	//iterate the array and get asci of the character

	var sum int32 = 5381
	for _, val := range key {
		//sum = sum + (sum << 5) + val
		sum = sum + val
	}

	index := sum % int32(len(mapobj.arr))
	if index < 0 {
		return -index
	}
	return index

}

func (mapobj *GoMap) Set(key string, value string) {
	entry := &Entry{Key: key, Value: value}
	index := mapobj.GetHash(key)

	entry.Next = mapobj.arr[index]
	mapobj.arr[index] = entry

	//fmt.Printf("Stored %s %s at %d ", key, value, index)
}

func GetValue(entry *Entry, key string) string {

	if entry.Key == key {
		return entry.Value
	}
	if entry.Next != nil {
		return GetValue(entry.Next, key)
	}
	return key + " key not found"

}
func (mapobj *GoMap) Get(key string) string {
	index := mapobj.GetHash(key)
	if index < 0 {
		return ""
	}
	entry := mapobj.arr[index]
	if entry == nil {
		return key + " key not found"
	}
	return GetValue(entry, key)

}
func main() {
	gomap := &GoMap{}
	gomap.Set("India", "UG")
	gomap.Set("Indai", "MBA")
	gomap.Set("Scotland", "Masters")
	fmt.Println(gomap)
	fmt.Println(gomap.Get("India"))
	fmt.Println(gomap.Get("Indai"))
	fmt.Println(gomap.Get("Scotland"))

}
