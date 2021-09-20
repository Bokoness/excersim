package main

func main() {
	fmt.Println(Handshake(19))
}
func Handshake(n uint) []string {
	fmt.Printf("%b \n", n)
	b := fmt.Sprintf("%b", n)
	for i, v := range b {
		fmt.Println(i, v)
	}
	return []string{"hey"}
}
