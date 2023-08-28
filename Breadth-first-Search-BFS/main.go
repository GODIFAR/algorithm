package main

import "fmt"

type GrapgMap map[string][]string

func main() {

	var graprhmap GrapgMap = make(GrapgMap, 0)

	graprhmap["me"] = []string{"a", "b", "c"}
	graprhmap["a"] = []string{"ccc,ddd"}
	graprhmap["b"] = []string{"ddd", "fff"}
	graprhmap["c"] = []string{"fff", "hhh"}
	graprhmap["ccc"] = []string{}
	graprhmap["ddd"] = []string{}
	graprhmap["fff"] = []string{}
	graprhmap["hhh"] = []string{}

	search_quene := graprhmap["me"]

	for {
		var person string
		if len(search_quene) > 0 {
			person, search_quene = search_quene[0], search_quene[1:]
			if Isccc(person) {
				fmt.Printf("%s is the man", person)
				break
			} else {
				//fmt.Println(search_quene)
				search_quene = append(search_quene, graprhmap[person]...)

			}
		} else {
			fmt.Println("Not Found")
			break
		}
	}

}
func Isccc(person string) bool {
	return person == "fff"

}
