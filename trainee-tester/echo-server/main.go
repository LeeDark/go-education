package main

func main() {
	e := setEndpoints()
	e.Logger.Fatal(e.Start(":1323"))
}
