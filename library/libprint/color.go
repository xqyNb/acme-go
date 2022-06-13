package libprint

import "fmt"

// 文本颜色
const (
	TextDefault = 0
	TextBlack   = iota + 29
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

// 背景颜色
const (
	BgDefault = 0
	BgBlack   = iota + 39
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// HintWarning 定义Hint标志
const (
	HintWarning = "【 警告 】"
	HintPanic   = "【 异常 】"
)

// SetColor 设置颜色
func SetColor(text string, textColor int, colors ...int) string {
	bgColor := BgDefault
	if len(colors) >= 1 {
		bgColor = colors[0]
	}

	return fmt.Sprintf("%c[%d;%dm %s%c[0m", 0x1B, bgColor, textColor, text, 0x1B)
}

// PrintHint 打印提示文本
func PrintHint(hintFlag, text string) {
	PrintColorln(fmt.Sprintf("%s %s", hintFlag, text), TextYellow)
}

// PrintErrorHint 打印错误文本
func PrintErrorHint(hintFlag, text string) {
	PrintColorln(fmt.Sprintf("%s %s", hintFlag, text), TextRed)
}

// PrintColorln 输出颜色
func PrintColorln(text string, textColor int, colors ...int) {
	fmt.Println(SetColor(text, textColor, colors...))
}
