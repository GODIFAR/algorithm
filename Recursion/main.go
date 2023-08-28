package main

// type Item struct {
// 	ID    int
// 	Type  string
// 	Child *Item
// }

// // type ItemClassifier interface {
// // 	IsDoll() bool
// // }

// func IsDoll(it *Item) bool {
// 	if it.Type == "doll" {
// 		return true
// 	}
// 	return false
// }

// func findDiamond(item Item) Item {
// 	if IsDoll(&item) {
// 		return findDiamond(*item.Child)
// 	} else {
// 		return item
// 	}

// }

func main() {
	// doll := Item{
	// 	ID:   1,
	// 	Type: "doll",
	// 	Child: &Item{
	// 		ID:   2,
	// 		Type: "doll",
	// 		Child: &Item{
	// 			ID:   3,
	// 			Type: "doll",
	// 			Child: &Item{
	// 				ID:    4,
	// 				Type:  "diamond",
	// 				Child: nil,
	// 			},
	// 		},
	// 	},
	// }
	// diamond := findDiamond(doll)
	// fmt.Printf("Item %d is diamond\n", diamond.ID)
}
