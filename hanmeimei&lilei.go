package main

import (
	"fmt"
)

type Student interface {
	learningEnglish(string)
}

type Human struct {
	Name string
	Sex  string
}

func (human Human) learningEnglish(learning string) {
	fmt.Printf("%s: %s\n", human.Name, learning)
}

type Male struct {
	Human "嵌入字段"
}

type Female Male

func (this Human) Pee(how string) {
	fmt.Printf("%s %s %s 撒尿\n", this.Name, this.Sex, how)
}

func doLearning(student Student, learning string) {
	student.learningEnglish(learning)
}

func doPee(human interface{}) {
	switch sex := human.(type) {
	case Male:
		sex.Pee("站着")
	case Female:
		sex.Pee("坐着")
	}
}

func main() {
	lilei := Male{}
	lilei.Name = "李雷"
	lilei.Sex = "男"
	fmt.Printf("%s %s 出场\n", lilei.Name, lilei.Sex)

	hanmeimei := Female{}
	hanmeimei.Name = "韩梅梅"
	hanmeimei.Sex = "女"
	fmt.Printf("%s %s 出场\n", hanmeimei.Name, hanmeimei.Sex)
	doPee(lilei)
	doPee(hanmeimei)

	doLearning(lilei, "How are you?")
	doLearning(hanmeimei, "I'm fine, thank you !")
}
