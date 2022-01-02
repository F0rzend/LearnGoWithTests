package maps

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
