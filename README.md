
# go-pbo
[![Go Report Card](https://goreportcard.com/badge/github.com/TheMysteriousVincent/go-pbo)](https://goreportcard.com/report/github.com/TheMysteriousVincent/go-pbo)

*An easy-to-use package for building ArmA 3 .pbo-files in Go.*

## How to use

go-pbo is a easy-to-use package for creating .pbo-files for ArmA 3.
You can create a file by just call these function(s):
```go
pbo := pbo.New() //get a new PBO object
pbo.From = "test/testPbo" //set the mod directory destination
pbo.To = "test/test.pbo" //set the mod file target
pbo.Prefix = "testPbo" //set an optional mod prefix
pbo.Generate() //generate the buffer output
pbo.Save() //save the buffer output
```

Thats it!

**Have fun.**
