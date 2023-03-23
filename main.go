package main
import(
	"fmt"
	"net/http"
	"log"
)
func hihandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path!="/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w,"Hello!")
}
func fHandler(w http.ResponseWriter,r *http.Request){
	// if r.URL.Path!="/form"{
	// 	http.Error(w,"404 not found",http.StatusNotFound)
	// 	return
	// }
	// if r.Method!="POST"{
	// 	http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	// 	return
	// }
	if err:=r.ParseForm();err!=nil{
		fmt.Fprintf(w,"ParseForm() err: %v",err)
		return
	}
    fmt.Fprintf(w,"Yeah! You are in the form page\n")
	n:=r.FormValue("name")
	e:=r.FormValue("email")
	fmt.Fprintf(w,"Name: %s\n",n)
	fmt.Fprintf(w,"Email: %s\n",e)
}
func main(){
	fs:=http.FileServer(http.Dir("./static"))
	http.Handle("/",fs)
	http.HandleFunc("/hello",hihandler)
	http.HandleFunc("/form",fHandler)
	fmt.Println("Server is running at port 8080")
	if err:=http.ListenAndServe(":8080",nil);err!=nil{
		log.Fatal(err)
	}
}