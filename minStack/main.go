package main

func main() {

}

type MinStack struct {
	Val []int
	Min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var minStack MinStack
	return minStack
}

func (this *MinStack) Push(x int) {
	if len(this.Val) == 0 {
		this.Min = append(this.Min, x)
	}
	if x <= this.Min[len(this.Min)-1] {
		this.Min = append(this.Min, x)
	}
	this.Val = append(this.Val, x)
}

func (this *MinStack) Pop() {
	if this.Val[len(this.Val)-1] == this.Min[len(this.Min)-1] {
		this.Min = this.Min[:len(this.Min)-1]
	}
	this.Val = this.Val[:len(this.Val)-1]

}

func (this *MinStack) Top() int {
	return this.Val[len(this.Val)-1]
}

func (this *MinStack) GetMin() int {
	// var min int
	// for i, n:=range this.Val{
	//     if i==0{
	//         min=n
	//     }
	//     if min>=n{
	//         min=n
	//     }
	// }
	if len(this.Min) > 0 {
		return this.Min[len(this.Min)-1]
	}
	return 0

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
