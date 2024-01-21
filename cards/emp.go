package main

type getNameObj interface {
	getName() string
}
type dept struct {
	deptId int
	dName  string
}
type employee struct {
	name string
	age  int
	dept
}

func getGenericName(obj getNameObj) string {
	return obj.getName()
}
func (d dept) getName() string {
	return d.dName
}
func (e employee) getName() string {
	return e.name
}

func (ptr *employee) updateName(_name string) {
	(*ptr).name = _name
}
func (e employee) updateName2(_name string) {
	e.name = _name
}
