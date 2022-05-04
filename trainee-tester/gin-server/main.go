package main

func main() {
	r := setEndpoints()
	r.Run() // listen and serve on 0.0.0.0:8080
}
