package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type post struct{
	Userid int `json:"userid"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

func main(){
	url:="https://jsonplaceholder.typicode.com/posts/1"

	response, err:=http.Get(url)
	if err!=nil{
		fmt.Printf("HTTP Get request failed: %s\n",err)
		return
	}
   defer response.Body.Close()

   body, err:=ioutil.ReadAll(response.Body)
   if err!=nil{
	fmt.Printf("Failed to read respopnse body: %s\n",err)
	return
   }
   var post post
   err=json.Unmarshal(body,&post)
   if err!=nil{
	fmt.Printf("Failed to parse json: %s\n",err)
	return
   }
    fmt.Printf("User id: %d\n",post.Userid)
    fmt.Printf("id: %d\n",post.Id)
    fmt.Printf("Title: %d\n",post.Title)
    fmt.Printf("Body: %d\n",post.Body)
   



}