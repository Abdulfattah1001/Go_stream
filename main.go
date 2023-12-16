package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Helllo World")
}

func main(){

	PORT:= os.Getenv("PORT")


	http.HandleFunc("/", handler)

	fmt.Println("Server is listening on Port number")
	var err error

	//err:=http.ListenAndServe(":8080", nil)
	if PORT!=""{
		err = http.ListenAndServe(PORT,nil)
	}else{

		PORT = ":8080"
		err = http.ListenAndServe(PORT, nil)
	}

	if err!=nil{
		fmt.Println("Error", err)
	}

	/**url:="https://stream-xhaz.onrender.com"

	response, err:= http.Get(url)

	if err!=nil{
		fmt.Println("ERROR: ", err)
		return
	}

	defer response.Body.Close()

	body, err:= io.ReadAll(response.Body)
	if err!=nil{
		fmt.Println("Error reading response: ", err)
		return
	}

	fmt.Println("Response body: ", string(body))*/
}