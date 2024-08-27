package main

func main() {
	list := List{head: &Node{data: 1, nextNode: nil}}

	item2 := Node{data: 2, nextNode: nil}
	item3 := Node{data: 3, nextNode: nil}
	item4 := Node{data: 4, nextNode: nil}
	item5 := Node{data: 5, nextNode: nil}

	list.addNewNode(&item2)
	list.addNewNode(&item3)
	list.addNewNode(&item4)
	list.addNewNode(&item5)

	list.printList()
}
