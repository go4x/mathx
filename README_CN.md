# MathX - 高精度数学计算库

MathX 是一个基于 `shopspring/decimal` 库的高精度数学计算包，专为 Go 语言设计。它提供了精确的十进制运算、链式 API 调用和丰富的数学函数，解决了 Go 原生浮点数运算精度不足的问题。

## 特性

- 🎯 **高精度计算** - 基于 `shopspring/decimal` 库，避免浮点数精度问题
- 🔗 **链式 API** - 支持流畅的方法链调用
- 🛡️ **Safe 函数** - 提供直接使用 `decimal.Decimal` 的高精度函数，避免精度丢失
- 📊 **统计函数** - 提供平均值、中位数、标准差等统计计算
- 💰 **货币格式化** - 支持千位分隔符和货币格式化
- 🧮 **数学函数** - 包含幂运算、开方、取整等常用数学函数
- 🔢 **整数运算** - 提供 int64 相关的精确运算函数
- 🛡️ **类型安全** - 使用 Go 泛型确保类型安全
- 📝 **完整文档** - 提供详细的 API 文档和使用示例

## 安装

```bash
go get github.com/go4x/mathx
```

## 快速开始

### 基本运算

```go
package main

import (
    "fmt"
    "github.com/go4x/mathx"
    "github.com/shopspring/decimal"
)

func main() {
    // 精确的加法运算
    result := mathx.Add(0.1, 0.2)
    fmt.Printf("0.1 + 0.2 = %s\n", result.ToString())
    // 输出: 0.1 + 0.2 = 0.3
    
    // 链式操作
    chainResult := mathx.Add(0.1, 0.2).
        Add(decimal.NewFromFloat(10)).
        Div(decimal.NewFromFloat(3), 2).
        Round(2).
        ToStringFixed(2)
    fmt.Printf("链式结果: %s\n", chainResult)
    // 输出: 链式结果: 1.00
}
```

### 高精度运算（Safe 函数）

```go
// 使用 Safe 函数避免精度丢失
a, _ := decimal.NewFromString("0.1")
b, _ := decimal.NewFromString("0.2")
result := mathx.AddSafe(a, b)
fmt.Printf("高精度结果: %s\n", result.String())
// 输出: 高精度结果: 0.3

// Safe 函数链式操作
values := []decimal.Decimal{
    decimal.RequireFromString("0.1"),
    decimal.RequireFromString("0.2"),
    decimal.RequireFromString("0.3"),
}
sum := mathx.SumSafe(values...)
fmt.Printf("高精度求和: %s\n", sum.String())
// 输出: 高精度求和: 0.6
```

### 货币格式化

```go
// 货币格式化
price := mathx.Mul(99.99, 1.15).
    Round(2).
    FormatMoney(2)
fmt.Printf("价格: $%s\n", price)
// 输出: 价格: $114.99

// 大数字格式化
amount := mathx.Mul(1000, 12.5).
    FormatMoney(2)
fmt.Printf("金额: $%s\n", amount)
// 输出: 金额: $12,500.00

// 整数格式化
intAmount := mathx.FormatMoneyInt(1234567, 2)
fmt.Printf("整数金额: $%s\n", intAmount)
// 输出: 整数金额: $1,234,567.00
```

### 统计计算

```go
// 统计计算
scores := []float64{85, 92, 78, 96, 88}

avg := mathx.Average(scores...)
std := mathx.StandardDeviation(scores...)

fmt.Printf("平均分: %.2f\n", avg)
fmt.Printf("标准差: %.2f\n", std)
// 输出:
// 平均分: 87.80
// 标准差: 6.87
```

## API 参考

### 基本运算

#### 基础函数

```go
// 基本运算
func Add(a, b float64) Result
func Sub(a, b float64) Result
func Mul(a, b float64) Result
func Div(a, b float64, precision int32) Result
func DivTrunc(a, b float64, precision int32) Result
```

#### Safe 函数（高精度）

Safe 函数直接接受 `decimal.Decimal` 类型，避免 float64 转换时的精度丢失：

```go
// 高精度运算
func AddSafe(a, b decimal.Decimal) Result
func SubSafe(a, b decimal.Decimal) Result
func MulSafe(a, b decimal.Decimal) Result
func DivSafe(a, b decimal.Decimal, precision int32) Result
func DivTruncSafe(a, b decimal.Decimal, precision int32) Result
```

**使用示例**：
```go
a, _ := decimal.NewFromString("0.1")
b, _ := decimal.NewFromString("0.2")
result := mathx.AddSafe(a, b) // 完全精确的 0.3
```

### 整数运算

```go
// 整数除法
func Int64Div(dividend, divisor int64, precision int32) float64
func Int64DivTrunc(dividend, divisor int64, precision int32) float64

// 整数与浮点数乘法
func Int64MulFloat64(multiplicand int64, multiplier float64) float64
```

**使用示例**：
```go
// 整数除法
result := mathx.Int64Div(1, 3, 4) // 0.3333

// 整数截断除法
result := mathx.Int64DivTrunc(10, 3, 2) // 3.33

// 整数与浮点数乘法
result := mathx.Int64MulFloat64(100, 0.1) // 10.0
```

### 数学函数

```go
// 基础数学函数
func Round(value float64, precision int32) Result
func Truncate(value float64, precision int32) Result
func Abs(value float64) float64
func Ceil(value float64) float64
func Floor(value float64) float64
func Pow(base, exponent float64) float64
func Sqrt(value float64) float64
func IsEqual(a, b float64, precision int32) bool

// Safe 版本的数学函数
func RoundSafe(a decimal.Decimal, precision int32) Result
func TruncateSafe(a decimal.Decimal, precision int32) Result
func AbsSafe(a decimal.Decimal) decimal.Decimal
func CeilSafe(a decimal.Decimal) decimal.Decimal
func FloorSafe(a decimal.Decimal) decimal.Decimal
func PowSafe(a decimal.Decimal, exponent decimal.Decimal) decimal.Decimal
func IsEqualSafe(a, b decimal.Decimal, precision int32) bool
```

### 统计函数

```go
// 基础统计函数
func Average[T constraints.Integer | constraints.Float](ns ...T) float64
func StandardDeviation[T constraints.Integer | constraints.Float](ns ...T) float64
func Max[T constraints.Ordered](ns ...T) T
func Min[T constraints.Ordered](ns ...T) T
func Sum[T constraints.Integer | constraints.Float](ns ...T) T

// Safe 版本的统计函数
func AverageSafe(ds ...decimal.Decimal) decimal.Decimal
func MaxSafe(ds ...decimal.Decimal) decimal.Decimal
func MinSafe(ds ...decimal.Decimal) decimal.Decimal
func SumSafe(ds ...decimal.Decimal) decimal.Decimal
```

**使用示例**：
```go
// 基础统计
scores := []float64{85, 92, 78, 96, 88}
avg := mathx.Average(scores...)
max := mathx.Max(scores...)

// 高精度统计
values := []decimal.Decimal{
    decimal.NewFromInt(85),
    decimal.NewFromInt(92),
    decimal.NewFromInt(78),
}
avg := mathx.AverageSafe(values...)
max := mathx.MaxSafe(values...)
```

### 工具函数

```go
// 范围限制
func Clamp(value, min, max float64) float64
func ClampSafe(a decimal.Decimal, min, max decimal.Decimal) decimal.Decimal

// 线性插值
func Lerp(a, b, t float64) float64

// 固定精度
func ToFixed(value float64, places int32) float64

// 符号判断
func IsPositive(value float64) bool
func IsNegative(value float64) bool
func Sign(value float64) int
```

### 格式化函数

```go
// 货币格式化
func FormatCurrency(amount float64, decimalPlaces int32) string
func FormatMoney(amount float64, decimalPlaces int32) string
func FormatMoneyInt(amount int64, decimalPlaces int32) string

// 字符串转换
func ToString(value float64) string
func ToStringFixed(value float64, places int32) string
func ToStringBank(value float64, places int32) string

// 清理函数
func RemoveTrailingZeros(value float64) string
func RemoveTrailingZerosFixed(value float64, places int32) string
func CleanFloat(value float64) float64
func CleanFloatString(value float64) string

// 解析函数
func ParseFloat(s string) (float64, error)
```

### Result 类型方法

#### 创建 Result

```go
// 从 float64 创建
func NewResult(value float64) Result

// 从字符串创建（保持精度）
func NewResultFromString(value string) (Result, error)

// 获取底层 decimal.Decimal
func (r Result) Decimal() decimal.Decimal
```

#### 字符串转换

```go
func (r Result) String() string
func (r Result) ToString() string
func (r Result) ToStringFixed(places int32) string
func (r Result) ToStringBank(places int32) string
```

#### 格式化

```go
// 货币格式化
func (r Result) FormatMoney(decimalPlaces int32) string

// 清理尾随零
func (r Result) Clean() Result
```

#### 数学运算

```go
// 基本运算（接受 decimal.Decimal）
func (r Result) Add(other decimal.Decimal) Result
func (r Result) Sub(other decimal.Decimal) Result
func (r Result) Mul(other decimal.Decimal) Result
func (r Result) Div(other decimal.Decimal, precision int32) Result
func (r Result) DivTrunc(other decimal.Decimal, precision int32) Result

// 数学函数
func (r Result) Round(places int32) Result
func (r Result) Truncate(places int32) Result
func (r Result) Abs() Result
func (r Result) Neg() Result
```

#### 值提取

```go
func (r Result) Float64() float64
```

## 使用示例

### 1. 精确计算

```go
// 避免浮点数精度问题
result1 := 0.1 + 0.2  // 原生 Go: 0.30000000000000004
result2 := mathx.Add(0.1, 0.2).ToString()  // MathX: "0.3"

// 使用 Safe 函数获得最大精度
a, _ := decimal.NewFromString("0.1")
b, _ := decimal.NewFromString("0.2")
result3 := mathx.AddSafe(a, b).String()  // "0.3"（完全精确）
```

### 2. 链式操作

```go
// 复杂的链式计算
result := mathx.Add(100, 50).
    Add(decimal.NewFromFloat(10)).  // 使用 decimal.Decimal
    Div(decimal.NewFromFloat(3), 2).
    Round(2).
    FormatMoney(2)

fmt.Printf("最终结果: $%s\n", result)
// 输出: 最终结果: $55.00
```

### 3. 高精度统计

```go
// 使用 Safe 函数进行高精度统计
values := []decimal.Decimal{
    decimal.RequireFromString("85.5"),
    decimal.RequireFromString("92.3"),
    decimal.RequireFromString("78.9"),
}

avg := mathx.AverageSafe(values...)
max := mathx.MaxSafe(values...)
min := mathx.MinSafe(values...)
sum := mathx.SumSafe(values...)

fmt.Printf("平均值: %s\n", avg.String())
fmt.Printf("最大值: %s\n", max.String())
fmt.Printf("最小值: %s\n", min.String())
fmt.Printf("总和: %s\n", sum.String())
```

### 4. 金融计算

```go
// 计算小费
bill := 50.00
tipPercent := 18.0
tip := mathx.Mul(bill, tipPercent/100)
total := mathx.Add(bill, tip.Float64())

fmt.Printf("账单: $%.2f\n", bill)
fmt.Printf("小费 (%.0f%%): $%.2f\n", tipPercent, tip.Float64())
fmt.Printf("总计: $%.2f\n", total.Float64())
```

### 5. 数据清理

```go
// 清理浮点数尾随零
value := 3.140000
cleanValue := mathx.RemoveTrailingZeros(value)
fmt.Printf("清理前: %f\n", value)
fmt.Printf("清理后: %s\n", cleanValue)
// 输出:
// 清理前: 3.140000
// 清理后: 3.14

// 使用链式操作清理
result := mathx.NewResult(3.140000).
    Clean().
    ToString()
fmt.Printf("链式清理: %s\n", result)
// 输出: 链式清理: 3.14
```

### 6. 整数运算

```go
// 整数除法
result := mathx.Int64Div(1, 3, 4)
fmt.Printf("1 / 3 (4位小数): %.4f\n", result)
// 输出: 1 / 3 (4位小数): 0.3333

// 整数截断除法
result := mathx.Int64DivTrunc(10, 3, 2)
fmt.Printf("10 / 3 (截断2位): %.2f\n", result)
// 输出: 10 / 3 (截断2位): 3.33

// 整数与浮点数乘法
result := mathx.Int64MulFloat64(100, 0.1)
fmt.Printf("100 * 0.1: %.1f\n", result)
// 输出: 100 * 0.1: 10.0
```

### 7. 高精度比较

```go
// 使用 IsEqualSafe 进行高精度比较
a := decimal.RequireFromString("3.14")
b := decimal.RequireFromString("3.1400000001")
equal := mathx.IsEqualSafe(a, b, 8)
fmt.Printf("是否相等（8位精度）: %v\n", equal)
// 输出: 是否相等（8位精度）: true
```

## 性能对比

MathX 使用高精度 decimal 库，虽然比原生 Go 数学运算稍慢，但提供了更高的精度：

```go
// 基准测试结果 (Apple M4 Pro)
BenchmarkAdd-12                     	 1658899	       737.8 ns/op
BenchmarkMul-12                     	 1613665	       734.8 ns/op
BenchmarkDiv-12                     	 4808594	       260.8 ns/op
BenchmarkChainableOperations-12     	 1000000	      1029 ns/op

// 原生 Go 运算对比
BenchmarkNativeAdd-12               	1000000000	         0.2595 ns/op
BenchmarkNativeMul-12               	1000000000	         0.2586 ns/op
```

**性能建议**：
- 对于精度要求不高的场景，可以使用基础函数
- 对于精度要求高的场景（如金融计算），使用 Safe 函数
- 在性能关键路径中，考虑直接使用 `float64` 返回值而不是链式操作

## 精度说明

### 何时使用基础函数

- 精度要求不高的一般计算
- 性能优先的场景
- 输入已经是 `float64` 类型

### 何时使用 Safe 函数

- 金融计算等精度敏感场景
- 需要避免 float64 转换精度丢失
- 输入可以转换为 `decimal.Decimal` 类型

### 最佳实践

```go
// ✅ 推荐：使用字符串创建 Result，然后使用链式操作
a, _ := mathx.NewResultFromString("0.1")
b, _ := mathx.NewResultFromString("0.2")
result := a.Add(b.Decimal()).Mul(decimal.NewFromInt(10))

// ✅ 推荐：使用 Safe 函数进行高精度计算
a, _ := decimal.NewFromString("0.1")
b, _ := decimal.NewFromString("0.2")
result := mathx.AddSafe(a, b)

// ⚠️ 注意：基础函数可能有精度损失
result := mathx.Add(0.1, 0.2)  // 输入精度已丢失
```

## 测试覆盖率

- **测试覆盖率**: 包含单元测试、示例函数和基准测试
- **测试用例**: 覆盖所有 API 函数
- **示例函数**: 提供完整的使用示例

## 依赖

- `github.com/shopspring/decimal` - 高精度十进制运算库
- `golang.org/x/exp/constraints` - Go 泛型约束

## 错误处理

大多数函数直接返回 `float64` 或 `Result` 类型。可能失败的函数（如 `ParseFloat` 和 `NewResultFromString`）返回 `(value, error)` 元组。

```go
value, err := mathx.ParseFloat("invalid")
if err != nil {
    // 处理错误
}

result, err := mathx.NewResultFromString("invalid")
if err != nil {
    // 处理错误
}
```

## 许可证

MIT.

## 贡献

欢迎提交 Issue 和 Pull Request！
