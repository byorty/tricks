package main

import "fmt"

type Getter interface {
	GetId() int
}

type Setter interface {
	SetId(int)
}

type Identifier interface {
	Getter
	Setter
}

type Slice interface {
	Len() int
	Get(int) interface{}
}

type List []Identifier

func (l List) Len() int {
	return len(l)
}

func (l List) Get(i int) Identifier {
	return l[i]
}

type User struct {
	id int
}

func (u User) GetId() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

type Post struct {
	id int
}

func (p Post) GetId() int {
	return p.id
}

func (p *Post) SetId(id int) {
	p.id = id
}

func main() {
	list := List{new(User), new(User), new(Post), new(Post), new(Post)}

	for i, item := range list {
		assign(item, i)
		print(item)
	}
}

func assign(setter Setter, i int) {
	setter.SetId(i + 1)
}

func print(getter Getter) {
	fmt.Printf("%T %d\n", getter, getter.GetId())
}
