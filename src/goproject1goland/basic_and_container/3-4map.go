package main

import "fmt"

func nonrepeating(s string) int {
	lastoccired := make(map[byte]int)
	start :=0
	maxlength :=0
	for i,ch :=range []byte(s){
		if last,ok:=lastoccired[ch];ok && last >=start{
			start = last+1
		}
		if i- start+1>maxlength{
			maxlength=i-start+1
		}
		lastoccired[ch]=i
	}
	return maxlength
}

func main() {
	m :=map[string] string{
		"name":"lisi",
		"age":"18",
		"city":"nanjing",
		"pet":"dog",
	}
	fmt.Println(m)

	m2 := make(map[string]int)//m2==empty map
	var m3 map[string]int//m3==nil
	fmt.Println(m2,m3)

	fmt.Println("=================travaling map")//无序
	for k,v :=range m{
		fmt.Println(k,v)
	}

	fmt.Println("=================getting value")//无序
	name := m["name"]
	fmt.Println(name)

	if sex,ok := m["sex"];ok{//判断存不存在
		fmt.Println(sex)
	}else {
		println("key is not exist")
	}

	fmt.Println("=================deleting key")//无序
	name,ok :=m["name"]
	fmt.Println(name,ok)

	delete(m,"name")
	name,ok =m["name"]  //不存在
	fmt.Println(name,ok)

	fmt.Println(
		nonrepeating("abcbabcbcbabcab"),
		nonrepeating("omnjkk"),
		nonrepeating(""),
		nonrepeating("a"),
		nonrepeating("aaaa"),
		nonrepeating("abcdefgh"),
		)
}
