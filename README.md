# go-rpmdb

rpm database information

[godoc.org](https://godoc.org/github.com/vbatts/go-rpmdb)

## Example

```golang
pkgs, err := NVRs()
if err != nil { 
	// do the thing
}
for _, pkg := range pkgs {
	info, err := Info(pkg)
	if err != nil { 
		// do the thing
	}
	fmt.Printf("%s-%s-%s :: %s\n", info.Name(), info.Version(), info.Release(), info.Sourcerpm())
	//fmt.Printf("%#v\n", info)
}
```
