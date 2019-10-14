package bst

// type Node struct {
// 	Key   int
// 	Left  *Node
// 	Right *Node
// }

// func (n *Node) Search(key int) bool {
// 	if n == nil {
// 		return false
// 	}

// 	if n.Key < key {
// 		return n.Right.Search(key)
// 	} else if n.Key > key {
// 		return n.Left.Search(key)
// 	}
// 	//n.Key == key
// 	return true
// }

// func (n *Node) Insert(key int) {
// 	if n.Key < key {
// 		if n.Right == nil {
// 			n.Right = &Node{Key: key}
// 		} else {
// 			n.Right.Insert(key)
// 		}
// 	} else if n.Key > key {
// 		if n.Left == nil {
// 			n.Left = &Node{Key: key}
// 		} else {
// 			n.Left.Insert(key)
// 		}
// 	}
// }

// func (n *Node) Delete(key int) {
// 	if n.Key < key {
// 		n.Right = n.Right.Delete(key)
// 	} else if n.Key > key {
// 		n.Left = n.Left.Delete(key)
// 	} else { //n.Key == key

// 	}
// }
