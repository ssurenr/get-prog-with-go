package main

import "fmt"

func main() {
	dwarfs1 := []string{"Cerus", "Pluto", "Hauma", "Makemake", "Eris"}
	dwarfs2 := append(dwarfs1, "Orcus")
	dwarfs3 :=  append(dwarfs2, "Salacia", "Quaror", "Sedna")

	fmt.Println(dwarfs1)
	fmt.Println(dwarfs2)
	fmt.Println(dwarfs3)

	dwarfs3[1] = "Pluto!"

	fmt.Println(dwarfs1)
	fmt.Println(dwarfs2)
	fmt.Println(dwarfs3)
}
