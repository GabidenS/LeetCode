package main

func main() {

}

type MyHashMap struct {
	key   interface{}
	value interface{}
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	m := make(MyHashMap{key})
	return m
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {

}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {

}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {

}
