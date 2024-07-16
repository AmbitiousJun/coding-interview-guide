// 解法1
// 在猫和狗两个类的基础上, 包上一层封装类
// 封装类除了携带动物对象之外, 还会携带一个加入队列的时间戳
// 只要在入队时维护好时间戳变量, 在出队的时候就能够依据时间戳正确出队
package main

import "fmt"

func main() {
	dcq := NewDogCatQueue()
	dcq.Add(new(Cat))
	dcq.Add(new(Dog))
	dcq.Add(new(Cat))
	dcq.Add(new(Dog))
	dcq.Add(new(Cat))
	dcq.Add(new(Dog))
	dcq.Add(new(Dog))
	fmt.Println(dcq.PollAll().GetPetType())
	fmt.Println(dcq.PollCat().GetPetType())
	fmt.Println(dcq.PollAll().GetPetType())
	fmt.Println(dcq.PollAll().GetPetType())
	fmt.Println(dcq.PollDog().GetPetType())
	fmt.Println(dcq.PollDog().GetPetType())
	fmt.Println(dcq.PollCat().GetPetType())
}

type Pet interface {
	GetPetType() string
}

type Dog struct{}

func (d *Dog) GetPetType() string { return "dog" }

type Cat struct{}

func (c *Cat) GetPetType() string { return "cat" }

type PetWrapper struct {
	Pet
	Cnt int
}

func NewPetWrapper(p Pet, cnt int) *PetWrapper {
	return &PetWrapper{
		Pet: p,
		Cnt: cnt,
	}
}

type DogCatQueue struct {
	DogQ []*PetWrapper
	CatQ []*PetWrapper
	Cnt  int
}

func NewDogCatQueue() *DogCatQueue {
	return &DogCatQueue{
		DogQ: make([]*PetWrapper, 0),
		CatQ: make([]*PetWrapper, 0),
		Cnt:  0,
	}
}

func (dcq *DogCatQueue) Add(p Pet) {
	if p.GetPetType() == "dog" {
		dcq.DogQ = append(dcq.DogQ, NewPetWrapper(p, dcq.Cnt))
		dcq.Cnt++
	} else if p.GetPetType() == "cat" {
		dcq.CatQ = append(dcq.CatQ, NewPetWrapper(p, dcq.Cnt))
		dcq.Cnt++
	} else {
		panic("not dog or cat")
	}
}

func (dcq *DogCatQueue) PollAll() Pet {
	if !dcq.IsCatEmpty() && !dcq.IsDogEmpty() {
		// 两条队列都不为空, 选一个时间戳小的弹出
		cFirst, dFirst := dcq.CatQ[0], dcq.DogQ[0]
		if cFirst.Cnt < dFirst.Cnt {
			dcq.CatQ = dcq.CatQ[1:]
			return cFirst.Pet
		}
		dcq.DogQ = dcq.DogQ[1:]
		return dFirst.Pet
	}
	if !dcq.IsCatEmpty() {
		// 只有 cat 队列有数据
		p := dcq.CatQ[0]
		dcq.CatQ = dcq.CatQ[1:]
		return p.Pet
	}
	if !dcq.IsDogEmpty() {
		// 只有 dog 队列有数据
		p := dcq.DogQ[0]
		dcq.DogQ = dcq.DogQ[1:]
		return p.Pet
	}
	panic("empty queue!")
}

func (dcq *DogCatQueue) PollDog() *Dog {
	if dcq.IsDogEmpty() {
		panic("empty dog queue")
	}
	p := dcq.DogQ[0]
	dcq.DogQ = dcq.DogQ[1:]
	return p.Pet.(*Dog)
}

func (dcq *DogCatQueue) PollCat() *Cat {
	if dcq.IsCatEmpty() {
		panic("empty cat queue")
	}
	p := dcq.CatQ[0]
	dcq.CatQ = dcq.CatQ[1:]
	return p.Pet.(*Cat)
}

func (dcq *DogCatQueue) IsEmpty() bool {
	return dcq.IsDogEmpty() && dcq.IsCatEmpty()
}

func (dcq *DogCatQueue) IsDogEmpty() bool {
	return len(dcq.DogQ) == 0
}

func (dcq *DogCatQueue) IsCatEmpty() bool {
	return len(dcq.CatQ) == 0
}
