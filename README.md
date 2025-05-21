Custom vet for struct tags
--------------------------

This repository implements a custom go vet that checks that all
struct tags are consitent.

## Run

Run with:
```console
$ cd ~/your/source/tree
$ go run github.com/mvo5/vet-tagseq/cmd/tagseq@latest ./...
```

## Example output

```go
package sad

type Sad struct {
	A string `json:"a" yaml:"b"`
}
```

When run:
```console
$ tagseq ./tagseq/testdata/src/sad/
.../tagseq/testdata/src/sad/sad.go:3:6: inconsistent struct tags found in "json:\"a\" yaml:\"b\""
```
