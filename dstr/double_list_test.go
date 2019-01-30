package dstr

import(
	"testing"
)

func TestDoubleList_Init(t *testing.T)  {
	list := new(DoubleList)
	list.Init()
	if list.Size == 0{
		t.Log("double list init success")
	} else {
		t.Error("double list init success")
	}
}

func TestDoubleList_Append(t *testing.T){
	list := new(DoubleList)
	list.Init()

	flag := list.Append(&DoubleNode{Data: 1})
	if flag && list.Size == 1{
		t.Log("double list append 1 success")
	} else {
		t.Error("double list append 1 success")
	}

	list.Display()

	flag = list.Append(&DoubleNode{Data: 2})
	if flag && list.Size == 2{
		t.Log("double list append 2 success")
	} else {
		t.Error("double list append 2 success")
	}

	list.Display()
}

func TestDoubleList_Get(t *testing.T){
	list := new(DoubleList)
	list.Init()

	list.Append(&DoubleNode{Data: 1})
	list.Append(&DoubleNode{Data: 2})

	node := list.Get(0)
	if 1 == node.Data.(int){
		t.Log("double list get 1 success")
	} else {
		t.Error("double list get 1 error")
	}
	list.Display()

	node = list.Get(1)
	if 2 == node.Data.(int){
		t.Log("double list get 2 success")
	} else {
		t.Error("double list get 2 error")
	}
}

func TestDoubleList_Insert(t *testing.T){
	list := new(DoubleList)
	list.Init()
	list.Insert(0, &DoubleNode{Data: 1})
	node := list.Get(0)
	if 1 == node.Data.(int){
		t.Log("double insert get 1 success")
	} else {
		t.Error("double insert get 1 error")
	}

	list.Insert(1, &DoubleNode{Data: 2})
	node = list.Get(1)
	if 2 == node.Data.(int){
		t.Log("double insert get 2 success")
	} else {
		t.Error("double insert get 2 error")
	}

	list.Insert(1, &DoubleNode{Data: 3})
	node = list.Get(1)
	if 3 == node.Data.(int){
		t.Log("double insert get 3 success")
	} else {
		t.Error("double insert get 3 error")
	}

	list.Display()
	list.Reverse()
}

func TestDoubleList_Delete(t *testing.T){
	list := new(DoubleList)
	list.Init()

	list.Append(&DoubleNode{Data: 1})
	list.Append(&DoubleNode{Data: 2})
	list.Append(&DoubleNode{Data: 3})
	list.Append(&DoubleNode{Data: 4})
	list.Append(&DoubleNode{Data: 5})

	// delete head
	flag := list.Delete(0)
	if flag && list.Size == 4{
		t.Log("double list delete 1 success")
	} else {
		t.Error("double list delete 1 error")
	}
	
	// delete tail 
	flag = list.Delete(3)
	if flag && list.Size == 3{
		t.Log("double list delete last success")
	} else {
		t.Error("double list delete last error")
	}

	// delete middle
	flag = list.Delete(1)
	if flag && list.Size == 2{
		t.Log("double list delete middle success")
	} else {
		t.Error("double list delete middle error")
	}

	list.Display()
}