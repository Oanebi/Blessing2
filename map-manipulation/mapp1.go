package main
import "fmt"
func Longest(words []string)( string,string){
	highest := words[0]
	lowest := words[0]
	for _, v := range words{
		if len(v) > len(highest){
			highest = v
		}
		if len(v) < len(lowest){
		lowest = v
		}
	}
	return lowest, highest
}
func main(){
	words :=  []string{"good", "tommorrow", "production"}
	low, high := Longest(words)
	
	fmt.Println("low:", low)
		fmt.Println("high:", high)

	

	}
