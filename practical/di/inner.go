package main

// inner layer -> abstraction
type Poem struct {
	content []byte
	storage PoemStorage
}

type PoemStorage interface {
	Types() string       // return a string describing the storage type
	Load(string) []byte  // load a poem by name
	Save(string, []byte) // save a poem by name
}

func NewPoem(ps PoemStorage) *Poem {
	return &Poem{
		content: []byte("I am a poem from a " + ps.Types() + "."),
		storage: ps,
	}
}

func (p *Poem) Save(name string) {
	p.storage.Save(name, p.content)
}

func (p *Poem) Load(name string) string {
	return string(p.content)
}
