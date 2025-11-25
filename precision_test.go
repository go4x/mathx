package mathx

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestPrec1(t *testing.T) {
	// 假设这是一个高精度的金额，比如比特币交易量或者极小的汇率
	// 这里有 19 位有效数字
	strVal := "1234567890123456.789"

	// 1. 使用 float64 传入 (模拟你的 Div 函数入参)
	fVal := 1234567890123456.789
	dFromFloat := decimal.NewFromFloat(fVal)

	// 2. 使用 String 传入 (正确做法)
	dFromString, _ := decimal.NewFromString(strVal)

	fmt.Println("原始数值: ", strVal)                  // 1234567890123456.789
	fmt.Println("float传入:", dFromFloat.String())   // 1234567890123456.8, 截断了一部分
	fmt.Println("string传入:", dFromString.String()) // 1234567890123456.789
}

func TestPrec2(t *testing.T) {
	// 场景：上一层业务逻辑用了 float 做减法
	// 理论上 1.0 - 0.9 应该等于 0.1
	a := 1.0 - 0.9

	// 此时 a 在 float64 里已经是 0.09999999999999998 了

	// 放入 decimal
	d := decimal.NewFromFloat(a)

	fmt.Println("a (float64):", a)         // 0.1
	fmt.Println("Decimal转换后:", d.String()) // 0.1
}

func TestPrec3(t *testing.T) {
	// 场景：两个 float64 相加
	v1 := 0.1
	v2 := 0.2

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)

	// 错误的做法：使用 NewFromFloat
	d1 := decimal.NewFromFloat(v1)
	d2 := decimal.NewFromFloat(v2)

	// 看看 decimal 到底存了什么
	fmt.Println("d1:", d1.String())
	fmt.Println("d2:", d2.String())

	// 计算加法
	sum := d1.Add(d2)
	fmt.Println("结果:", sum.String())
}
