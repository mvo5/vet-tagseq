package happy

type Happy struct {
	A string `json:"a" yaml:"a"`
	B int    `json:"b" yaml:"b"`
}

type NoTagsFine struct {
	A string
	B int
}
