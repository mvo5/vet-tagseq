package sad

type Sad struct { // want `inconsistent struct tags found in "json:\\"a\\" yaml:\\"b\\""`
	A string `json:"a" yaml:"b"`
}
