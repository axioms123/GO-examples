package main

import "fmt"
type Node struct{
    Value int
}

type Queue struct{
    nodes[] *Node
    count int
}
func NewQueue() *Queue{
    return &Queue{}
}
  func (s *Queue)  Push(n *Node){
         s.nodes=append(s.nodes[:s.count],n)
         
         s.count++//pont to next position
    }
func (s *Queue)  Pop()  *Node{
         if s.count==0{ return nil}
         s.count--
        return  s.nodes[s.count]
    }
func (n  *Node) String()string{
    return fmt.Sprint(n.Value)
}
func main(){
   s:=NewStack()
   s.Push( &Node{10} )
   s.Push( &Node{12} )
   s.Push( &Node{14} )
   s.Push( &Node{16} )

    fmt.Println(s.Pop()  ,s.Pop(),s.Pop(),s.Pop())

}
