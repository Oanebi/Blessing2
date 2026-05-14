package main
import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	if len(os.Args) != 3{
		fmt.Println("go run . input.txt output.txt")
	}
	inputpath := os.Args[1]
	outputpath := os.Args[2]
	inputfile,err := os.Open(inputpath)
	if err != nil{
		fmt.Println("Error 1", err)
		return
	}
	outputfile, err := os.Create(outputpath)
	if err != nil {
		fmt.Println("Error 2", err)
		return 
	}
	scanner := bufio.NewScanner(inputfile)
	writer := bufio.NewWriter(outputfile)
	for scanner.Scan(){
	line := scanner.Text()
	result := processor.go(line)
	writer := fmt.Println(result + "\n")
	}
	writer.Flush()
}
