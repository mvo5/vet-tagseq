package sad

type Dupe struct { // want `duplicate struct tags found in "yaml:\\"a\\" toml:\\"a\\""`
	A string `yaml:"a" toml:"a"`
	B string `yaml:"a" toml:"a"`
}
