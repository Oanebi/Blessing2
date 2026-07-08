
package main
import (
	"fmt"
	"net/http"
	"io"
	
)
func TextCounter(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet {
		fmt.Fprintf(w,"Send a POST request with text to count words")
		return
	}
	if r.Method == http.MethodPost{
		data, err := io.ReadAll(r.Body)
		if err != nil{
			fmt.Fprintln(w, "error reading file")
			return
		}
		result := string(data)
			fmt.Fprintf(w,"%d\n",len(result))
}
	}
