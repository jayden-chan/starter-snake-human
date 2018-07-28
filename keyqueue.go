package main

type KeyQueue struct {
	Keys []rune
}

func (k *KeyQueue) Enqueue(key rune) {
	k.Keys = append(k.Keys, key)
}

func (k *KeyQueue) Dequeue() (key rune) {
	key = k.Keys[0]
	k.Keys = k.Keys[1:]
	return
}

func (k *KeyQueue) Peek() (key rune) {
	return k.Keys[0]
}

func (k *KeyQueue) Clear() {
	k.Keys = []rune{}
}

func (k *KeyQueue) IsEmpty() (b bool) {
	return len(k.Keys) == 0
}

func (k *KeyQueue) Size() (s int) {
	return len(k.Keys)
}
