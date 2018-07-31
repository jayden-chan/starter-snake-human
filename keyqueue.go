package main

// KeyQueue is a Queue data structure for keys (runes)
type KeyQueue struct {
	Keys []rune
}

// Enqueue adds a key to the back of the queue
func (k *KeyQueue) Enqueue(key rune) {
	k.Keys = append(k.Keys, key)
}

// Dequeue removes and returns a key from the front of the queue
func (k *KeyQueue) Dequeue() (key rune) {
	key = k.Keys[0]
	k.Keys = k.Keys[1:]
	return
}

// Peek returns the key at the front of the queue
func (k *KeyQueue) Peek() (key rune) {
	return k.Keys[0]
}

// Clear removes all elements from the queue
func (k *KeyQueue) Clear() {
	k.Keys = []rune{}
}

// IsEmpty returns whether there are any elements in the queue
func (k *KeyQueue) IsEmpty() (b bool) {
	return len(k.Keys) == 0
}

// Size returns the number of elements in the queue
func (k *KeyQueue) Size() (s int) {
	return len(k.Keys)
}
