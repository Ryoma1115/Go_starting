func main() {
	defer func() {
		fmt.Println("defer")
	}()

	os.Exit(0)
}