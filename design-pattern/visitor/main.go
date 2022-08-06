package main

// データ構造とデータを処理するクラスを分離する

type User struct{}

func (u *User) Accept(v Visitor) {
	v.VisitUser(u)
}

type Group struct{}

func (g *Group) Accept(v Visitor) {
	v.VisitGroup(g)
}

type Visitor interface {
	VisitUser(*User)
	VisitGroup(*Group)
}

type VisitorImplement struct{}

func (v *VisitorImplement) VisitUser(u *User) {
	println("visit user")
}

func (v *VisitorImplement) VisitGroup(g *Group) {
	println("visit group")
}

func main() {
	u := &User{}
	g := &Group{}
	v := &VisitorImplement{}
	u.Accept(v)
	g.Accept(v)
}
