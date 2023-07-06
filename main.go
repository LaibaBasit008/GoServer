package main
import(
	"log"
	"net/http"
	
)
type Book struct{
	Title string `json:"title"`
	Author string `json:"author"`

}
func Hello(w http.ResponseWriter,r* http.Request){
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("<h1 color=steelblue>HELLO</h1>"))
}

func GetBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	book:=Book{Title:"harry", Author:"JK"}
	json.NewEncoder(w).Encode(book)
}
func main(){
	http.HandleFunc("/listening",Hello)
	http.HandleFunc("/book",GetBook)
	log.Fatal(http.ListenAndServe(":5000",nil))
}