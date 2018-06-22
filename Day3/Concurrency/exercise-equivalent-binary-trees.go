package main

import "golang.org/x/tour/tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int)
func Walk(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        Walk(t.Left, ch)
    }
    ch <- t.Value
    if t.Right != nil {
        Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(tree1, tree2 *tree.Tree) bool
func Same(tree1, tree2 *tree.Tree) bool {
    c1 := make(chan int, 10)
    c2 := make(chan int, 10)
    go Walk(tree1, c1)
    go Walk(tree2, c2)
    for i := 0; i < 10; i++ {
        x, y := <-c1, <-c2
        fmt.Printf("x: %v, y: %v\n", x, y)
        if x != y {
            return false
        }
    }
    return true
}


func main() {
}
