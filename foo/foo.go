package foo

// 大文字：外部参照可能（変数、関数共通）
const Max = 100

// 小文字：外部参照不可（変数、関数共通）
const min = 1

func ReturnMin() int {
	return min
}
