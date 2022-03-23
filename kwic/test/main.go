package main

import (
	"gobasic/test/kwic"
)

func main() {
	fileName := "C:\\Users\\薛文朝\\Desktop\\GolangCode\\test\\chooseAndSort.txt"
	kwic.Input(fileName)
	kwic.Shift()
	kwic.Sort()
	kwic.Output("result.txt")

	// kwic.Test()
	// kwic.TestStrings()
}
