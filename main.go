package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) []int {

	if m == 0 {
		for k := 0; k < n; k++ {
			nums1[k] = nums2[k]
		}
		return nums1
	}
	nums1 = moveToEnd(nums1, m)
	i := n
	j := 0
	for k := 0; k < m+n; k++ {
		if i < m+n && j < n {
			if nums1[i] < nums2[j] {
				nums1[k] = nums2[i]
				i++
			} else {
				nums1[k] = nums2[j]
				j++
			}
		} else if j < n {
			nums1[k] = nums2[j]
			j++
		}
	}

	return nums1
}
func moveToEnd(nums []int, m int) []int {
	lenNums := len(nums)
	k := lenNums
	var i int
	for i = m - 1; i >= 0; i-- {
		nums[k-1] = nums[i]
		k--
	}
	return nums

}

func main() {
	output := merge([]int{2, 3, 4, 0, 0}, 3, []int{1, 5}, 2)
	fmt.Println(output)

	output = merge([]int{4, 5, 0, 0, 0, 0}, 2, []int{1, 2, 3, 7}, 4)
	fmt.Println(output)
	fmt.Println("Hello Wrold")
}

/*



(
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Guri:Guri@test.tee7h.mongodb.net/?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)

}*/

/*func main() {
	fmt.Println(calculator.Divide(4.0, 2.0))
}*/

/*func Add(x,y int) int {
	return x+y

}
func main(){
	fmt.Println("Hello World")

}*/

/*const (
	concurrencyLevel = 4
)

var (
	requestIDs = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
)

func main() {
	queue := make(chan bool, concurrencyLevel)

	for _, _ID := range requestIDs {
		queue <- true
		go func(ID int) {
			defer func() { <-queue }()
			makeRequest(ID)

		}(_ID)

	}
	for i := 0; i < concurrencyLevel; i++ {
		queue <- true
	}
}

func makeRequest(ID int) {
	time.Sleep(time.Second * 3)
	fmt.Println(ID)
}









/*type Person struct {
	Name string `json:"Name"`
	Age  int    `json:"age,omitempty"`
}

//type server struct{}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	var user Person
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user)
	json.NewEncoder(w).Encode(user)

	//w.WriteHeader((http.StatusOK))
	//w.Write([]byte(`{"Name" :"customName"}`))
}
func main() {

	http.HandleFunc("/user", HandleUser)
	log.Fatal(http.ListenAndServe(":9000", nil))

}*/

/*type Person struct {
	Name       string `json: "customName"`
	Age        int    `json: "age,omitempty"`
	CreditCard string `json: "-"`
}

func main() {
	p := Person{Name: "Guri", CreditCard: "Secret"}
	pBytes, err := json.Marshal(p)

	log.Print(err)
	log.Print(string(pBytes))

	p2AsRawJson := json.RawMessage(`{"customName": "Guri"}`)
	var p2 Person
	err2 := json.Unmarshal(p2AsRawJson, &p2)
	log.Print(err2)
	log.Printf("%+v", p2)

}











































/*type book struct {
	Title  string `json: "title"`
	ID     int    `json: "id"`
	Price  int    `json: "price"`
	Author string `json: "author"`
}

func main() {

	fmt.Println("Good Morning")

	book := book{Title: "Go-Lang", ID: 1234, Price: 50, Author: "Guri"}

	fmt.Printf("%v\n", book)

}

/*func add(x int, y int) {
	fmt.Println(x + y)

}
func main() {
	add(5, 5)

}*/
/*func main() {
q := []int{1, 2, 3, 4, 5, 6}
fmt.Println(q)

r := []bool{true, false, true, false, true, true}
fmt.Println(r)

s := []struct {
	i int
	b bool
}{
	{1, true},
	{2, false},
	{3, true},
	{4, false},
	{5, true},
	{6, true},
}

fmt.Println(s)*/
