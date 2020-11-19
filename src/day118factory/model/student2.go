package model

type student2 struct {
	name  string
	score float64
}

func NewStudent(n string, s float64) *student2 {
	return &student2{
		name:  n,
		score: s,
	}
}

func (stu *student2) GetName() string {
	return stu.name
}

func (stu *student2) SetName(name string) string {
	stu.name = name
	return stu.name
}
