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

func (s *RandomizedSet) Insert(val int) bool {
	_, ok := s.set[val]
	if ok {
		return false
	}
	s.keys = append(s.keys, val)
	s.set[val] = len(s.keys) - 1

	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	_, ok := s.set[val]
	if ok {
		if s.set[val] == len(s.keys)-1 {
			s.keys = s.keys[:len(s.keys)-1]
			delete(s.set, val)
			return true
		}
		lastVal := s.keys[len(s.keys)-1]
		s.keys[s.set[val]] = lastVal
		s.keys = append(s.keys[:len(s.keys)-1], s.keys[len(s.keys)-1:]...)
		s.set[lastVal] = s.set[val]
		delete(s.set, val)
		return true
	}
	return false
}

func (s *RandomizedSet) GetRandom() int {
	if len(s.keys) == 0 {
		return 0
	}
	randomIdx := s.rand.Intn(len(s.keys))
	return s.keys[randomIdx]
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
