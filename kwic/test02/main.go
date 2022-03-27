package main

import (
	"gobasic/test02/input"
	"gobasic/test02/output"
	"gobasic/test02/shift"
	"gobasic/test02/sortAlph"
)

func main() {
	// 实例化input,读取txt数据加入lineTxt
	var input = new(input.Input)
	input.Input("C:\\Users\\薛文朝\\Desktop\\GolangCode\\test02\\chooseAndSort.txt")

	// 实例化shift类,将input类中的lineTxt传入shift类中并对lineTxt进行循环移位处理
	var shift = new(shift.Shift)
	shift.SetLineTxt(input.GetLineTxt())
	shift.Kwic()

	// 实例化sortAlph sort alphabet
	var sortAlph = new(sortAlph.SortAlph)
	sortAlph.SetKwicList(shift.GetLineTxt())
	sortAlph.Sort()

	// output in a new file
	var output = new(output.Output)
	output.SetKwicList(sortAlph.GetKwicList())
	output.SetNewFileName("result02.txt")
	output.Write()
}
