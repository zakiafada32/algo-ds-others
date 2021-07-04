package main

import "fmt"

type Notebook struct {
	poems map[string][]byte
}

func NewNotebook() *Notebook {
	return &Notebook{
		poems: map[string][]byte{},
	}
}

func (n *Notebook) Save(name string, content []byte) {
	n.poems[name] = content
}

func (n *Notebook) Load(name string) []byte {
	return n.poems[name]
}

func (n *Notebook) Types() string {
	return "Notebook"
}

type Napkin struct {
	poem []byte
}

func NewNapkin() *Napkin {
	return &Napkin{
		poem: []byte{},
	}
}

func (n *Napkin) Save(name string, content []byte) {
	n.poem = content
}

func (n *Napkin) Load(name string) []byte {
	return n.poem
}

func (n *Napkin) Types() string {
	return "Napkin"
}

func main() {
	notebook := NewNotebook()
	napkin := NewNapkin()

	poem := NewPoem(notebook)
	poem.Save("My first poem")

	poem = NewPoem(notebook)
	poem.Load("My first poem")
	fmt.Println(poem)

	poem = NewPoem(napkin)
	poem.Save("My second poem")

	poem = NewPoem(napkin)
	poem.Load("My second poem")
	fmt.Println(poem)

}
