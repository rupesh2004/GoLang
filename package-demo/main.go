// main.go
package main

import (
    "fmt"
    "package-demo/myPackage"
)

func main() {
    fmt.Println("Package demo")
    myPackage.AddVal(4, 5)
}
