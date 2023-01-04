package queue

type Queue []int
func (q *Queue)Push(v int){
	*q=append(*q,v)   //改变原有指针
}

func (q *Queue)Pop() int{
	head:= (*q)[0]
	*q =(*q)[1:]
	return head
}
func (q *Queue)ISEmpty()bool{
	return len(*q) == 0
}



func main() {
	
}
