type MinStack struct {
	Items []int
	MinItems []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.Items = append(this.Items, val)
	
	if len(this.MinItems) == 0 || val <= this.MinItems[len(this.MinItems)-1] {
		this.MinItems = append(this.MinItems, val)
	}
}

func (this *MinStack) Pop() {
	topItem := this.Items[len(this.Items)-1]
	this.Items = this.Items[:len(this.Items)-1]
	if len(this.MinItems) > 0 && topItem == this.MinItems[len(this.MinItems)-1]{
		this.MinItems = this.MinItems[:len(this.MinItems)-1]
	}
}

func (this *MinStack) Top() int {
	return this.Items[len(this.Items)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinItems[len(this.MinItems)-1]
}
