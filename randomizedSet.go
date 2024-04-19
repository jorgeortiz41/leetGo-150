package main

import "fmt"

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

type RandomizedSet struct {
}

func Constructor() RandomizedSet {
	return RandomizedSet{}
}

func (this *RandomizedSet) Insert(val int) bool {
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return 0
}

func TestRandomizedSet() {
	fmt.Println("Running RandomizedSet tests")
	obj := Constructor()
	fmt.Println("Insert 1:", obj.Insert(1))
	fmt.Println("Insert 2:", obj.Insert(2))
	fmt.Println("Insert 3:", obj.Insert(3))
	fmt.Println("Remove 2:", obj.Remove(2))
	fmt.Println("GetRandom:", obj.GetRandom())
	fmt.Println("GetRandom:", obj.GetRandom())
}
