package main

func main() {
	startDb()
	defer pool.Close()

	query("admin")
}
