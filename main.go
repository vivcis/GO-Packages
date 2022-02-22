package main

import (
	"fmt"
	"github.com/elliotchance/orderedmap"
)

func main() {

	pairs := orderedmap.NewOrderedMap()

	//The set function is used to replace a value for a key in the map
	pairs.Set("foo", "bar")
	pairs.Set("qux", 1.23)
	pairs.Set(123, true)
	pairs.Set("good", "0.55")
	pairs.Set("This is very interesting,", "oh yeah...")
	pairs.Set(1, 2)
	pairs.Set("go", "come")
	pairs.Set("yes", "no")
	pairs.Set("mean", "bad")
	pairs.Set("foo", "bar")
	pairs.Set("qux", 1.23)
	pairs.Set(123, true)
	pairs.Set("good", "0.55")
	pairs.Set("This is very interesting,", "oh yeah...")
	pairs.Set(1, 2)
	pairs.Set("go", "come")
	pairs.Set("yes", "no")
	pairs.Set("good", "bad")
	pairs.Set("sweet", "bitter")
	pairs.Set("Knowledge", "very valuable")
	pairs.Set("water", "thirsty")
	pairs.Set(567, 890)
	pairs.Set(1, "good")
	pairs.Set(2, "pass")
	pairs.Set(3, "excellent")
	pairs.Set(4, "average")
	pairs.Set(5, "poor")
	pairs.Set(6, "WONDERFUL")
	pairs.Set(7, "AWESOME")
	pairs.Set(8, "AMAZING")
	pairs.Set(9, "LOVELY")
	pairs.Set(10, "BAD")
	pairs.Set("name", "Sammy")
	pairs.Set("animal", "shark")
	pairs.Set("color", "blue")
	pairs.Set("location", "ocean")

	//the Delete function is used to delete/remove a key from the map.
	//It returns a boolean. If the key does exist and was deleted it returns true and if not it returns false
	//toDelete := pairs.Delete(234)
	//toDelete := pairs.Delete(678)

	//Get is used to access the value of a key. If it does not exist it returns nil and second parameter false
	//and if it does, it returns the value
	toGet, _ := pairs.Get("345")
	//toGet, _ := m.Get("good")

	//The len function is used to get the length of the map.
	//This returns only the number of unique keys
	toLen := pairs.Len()

	//This iteration is only suitable for a small number of items
	//for _, key := range pairs.Keys() {
	//	value, _ := pairs.Get(key)
	//	fmt.Println(key, value)
	//}
	//fmt.Println(toDelete)
	fmt.Println(toGet)
	fmt.Println(toLen)

	//To iterate through larger maps, you use Front() to print all
	//elements from oldest to newest

	for element := pairs.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Key, element.Value)
	}

	//You use Back() to iterate through all elements from newest to oldest
	//This is the reverse of using Front()

	//for element := pairs.Back(); element != nil; element = element.Prev() {
	//	fmt.Println(element.Key, element.Value)
	//}

}
