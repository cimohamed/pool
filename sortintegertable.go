package main
import("fmt")

func SortIntegerTable(table []int) {
	for i:=0;i<len(table);i++{
		for j:=0;j<len(table)-i-1;j++{
			if table[j]>table[j+1]{
				table[j],table[j+1]=table[j+1],table[j]
			}
		}
	}

}
func main() {
	s := []int{5,4,3,2,1,0}
	SortIntegerTable(s)
	fmt.Println(s)
}
// $ go run .
// [0 1 2 3 4 5]
// $

