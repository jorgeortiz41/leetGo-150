package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 380. Insert Delete GetRandom O(1) (Medium) https://leetcode.com/problems/insert-delete-getrandom-o1/
// Implement the RandomizedSet class:
// RandomizedSet() Initializes the RandomizedSet object.
// bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
// bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
// int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
//
// You must implement the functions of the class such that each function works in average O(1) time complexity.

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

// TODO: Create keys slice to track keys then choose a random one for getRandom
type RandomizedSet struct {
	set  map[int]int
	keys []int
	rand *rand.Rand
} // map[int]struct{} is a set

func Constructor() RandomizedSet {
	return RandomizedSet{set: make(map[int]int), keys: []int{}, rand: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.set[val]
	if ok {
		return false
	}
	this.keys = append(this.keys, val)
	this.set[val] = len(this.keys) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.set[val]
	if ok {
		if this.set[val] == len(this.keys)-1 {
			this.keys = this.keys[:len(this.keys)-1]
			delete(this.set, val)
			return true
		}
		lastVal := this.keys[len(this.keys)-1]
		this.keys[this.set[val]] = lastVal
		this.keys = append(this.keys[:len(this.keys)-1], this.keys[len(this.keys)-1:]...)
		this.set[lastVal] = this.set[val]
		delete(this.set, val)
		return true
	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.keys) == 0 {
		return 0
	}
	randomIdx := this.rand.Intn(len(this.keys))
	return this.keys[randomIdx]
}

func TestRandomizedSet() {
	fmt.Println("Running RandomizedSet tests")
	obj := Constructor()
	fmt.Println("RandomizedSet: ", obj.set, "keys:", obj.keys)
	fmt.Println("Insert 1:", obj.Insert(1), "Expected: ", "true")
	fmt.Println("RandomizedSet: ", obj)
	fmt.Println("Insert 1:", obj.Insert(1), "Expected: ", "false")

	fmt.Println("Remove 2:", obj.Remove(2), "Expected: ", "false")

	fmt.Println("Insert 2:", obj.Insert(2), "Expected: ", "true")
	fmt.Println("RandomizedSet: ", obj)

	fmt.Println("Remove 2:", obj.Remove(2), "Expected: ", "true")
	fmt.Println("RandomizedSet: ", obj)

	fmt.Println("Insert 3:", obj.Insert(3), "Expected: ", "true")
	fmt.Println("Insert -1:", obj.Insert(-1), "Expected: ", "true")
	fmt.Println("RandomizedSet: ", obj)

	fmt.Println("GetRandom:", obj.GetRandom())
	fmt.Println("GetRandom:", obj.GetRandom())
	fmt.Println("GetRandom:", obj.GetRandom())
	fmt.Println("GetRandom:", obj.GetRandom())
	fmt.Println("GetRandom:", obj.GetRandom())
	fmt.Println("Expected: ", "1, 3, or -1")
}
