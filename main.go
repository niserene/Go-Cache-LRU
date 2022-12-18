package main

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Key   string
	Value string
}

type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

type Cache struct {
	Queue        Queue
	Hash         Hash
	MaxQueueSize int
}

func NewQueue() Queue {

	return Queue{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

func NewCache(params ...int) *Cache {

	var maxQueueSize = 10
	if len(params) > 0 {
		maxQueueSize = params[0]
	}
	return &Cache{
		Queue:        NewQueue(),
		Hash:         Hash{},
		MaxQueueSize: maxQueueSize,
	}
}

func (cache *Cache) Remove(node *Node) {

	if node.Left != nil {
		node.Left.Right = node.Right
	}
	if node.Right != nil {
		node.Right.Left = node.Left
	}
	cache.Hash[node.Key] = nil
	if cache.Queue.Tail.Key == node.Key {
		cache.Queue.Tail = node.Left
	}
	if cache.Queue.Head.Key == node.Key {
		cache.Queue.Head = node.Right
	}
	cache.Queue.Size -= 1
}

func (cache *Cache) Add(node *Node) {

	queue := &(cache.Queue)
	if cache.Queue.Size == 0 {
		queue.Head = node
		queue.Tail = node
	} else {
		node.Right = queue.Head
		queue.Head.Left = node
		queue.Head = node
	}
	cache.Hash[node.Key] = node
	queue.Size += 1
	if queue.Size > cache.MaxQueueSize {
		fmt.Println("max length achived")
		cache.Remove(cache.Queue.Tail)
	}
}

func (cache Cache) ContainsKey(key string) bool {

	if cache.Hash[key] != nil {
		return true
	}
	return false
}

func (cache Cache) Get(key string) *Node {

	if !cache.ContainsKey(key) {
		return nil // or throw error
	}
	return cache.Hash[key]
}

func (cache *Cache) Check(key string, value string) {

	var node *Node
	if cache.Hash[key] == nil {
		node = &Node{
			Key:   key,
			Value: value,
		}
	} else {
		node = cache.Hash[key]
		cache.Remove(node)
		node.Value = value
	}
	cache.Add(node)
}

func (cache *Cache) Display() {

	tempHead := cache.Queue.Head
	fmt.Printf("Size : %d\n", cache.Queue.Size)
	for {
		if tempHead == nil {
			break
		}
		fmt.Print(tempHead.Key + " : " + tempHead.Value)
		tempHead = tempHead.Right
	}
	fmt.Print("\n\n")
}

type Hash map[string]*Node

func main() {

	fmt.Println("Cache start")

	cache := NewCache(10)

	for {

		var key string
		var value string
		fmt.Println("enter key value pair")
		fmt.Scan(&key)
		if key == "1" {
			break
		}
		fmt.Scan(&value)
		cache.Check(key, value)
		cache.Display()
	}

	fmt.Println("Ended")
}
