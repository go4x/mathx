package mathx

import (
	"math"
	"strings"

	"github.com/shopspring/decimal"
)

// Result represents a calculation result with chainable methods
type Result struct {
	v decimal.Decimal
}

// NewResult creates a new Result from a float64
func NewResult(value float64) Result {
	return Result{v: decimal.NewFromFloat(value)}
}

// NewResultFromString creates a new Result from a string
// This is useful for preserving precision when working with very large or very small numbers
func NewResultFromString(value string) (Result, error) {
	d, err := decimal.NewFromString(value)
	if err != nil {
		return Result{}, err
	}
	return Result{v: d}, nil
}

func (r Result) Decimal() decimal.Decimal {
	return r.v
}

// Float64 returns the float64 value
func (r Result) Float64() float64 {
	f, _ := r.v.Float64()
	return f
}

// String returns the string representation
func (r Result) String() string {
	return r.v.String()
}

// ToString returns the string representation
func (r Result) ToString() string {
	return r.v.String()
}

// ToStringFixed returns the string with fixed decimal places
func (r Result) ToStringFixed(places int32) string {
	return r.v.StringFixed(places)
}

// ToStringBank returns the string with banker's rounding
func (r Result) ToStringBank(places int32) string {
	return r.v.StringFixedBank(places)
}

// Clean removes trailing zeros and returns a new Result
func (r Result) Clean() Result {
	// 转换为字符串去除尾随零，再转回decimal
	str := r.v.String()
	if strings.Contains(str, ".") {
		str = strings.TrimRight(str, "0")
		str = strings.TrimRight(str, ".")
	}
	cleanValue, _ := decimal.NewFromString(str)
	return Result{v: cleanValue}
}

// Round rounds to specified precision and returns a new Result
func (r Result) Round(places int32) Result {
	return Result{v: r.v.Round(places)}
}

// Truncate truncates to specified precision and returns a new Result
func (r Result) Truncate(places int32) Result {
	if places < 0 {
		// For negative precision, truncate to integer places
		// e.g., precision -1 means truncate to tens place
		multiplier := decimal.NewFromFloat(math.Pow(10, float64(-places)))
		result := r.v.Div(multiplier).Truncate(0).Mul(multiplier)
		return Result{v: result}
	}
	return Result{v: r.v.Truncate(places)}
}

// FormatMoney formats as currency with thousands separator
func (r Result) FormatMoney(decimalPlaces int32) string {
	rounded := r.v.Round(decimalPlaces)
	str := rounded.StringFixed(decimalPlaces)

	// 分离整数和小数部分
	parts := strings.Split(str, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	// 添加千位分隔符
	if len(integerPart) > 3 {
		var result strings.Builder
		for i, char := range integerPart {
			if i > 0 && (len(integerPart)-i)%3 == 0 {
				result.WriteString(",")
			}
			result.WriteRune(char)
		}
		integerPart = result.String()
	}

	return integerPart + decimalPart
}

// Abs returns the absolute value
func (r Result) Abs() Result {
	return Result{v: r.v.Abs()}
}

// Neg returns the negative value
func (r Result) Neg() Result {
	return Result{v: r.v.Neg()}
}

// Add adds another decimal to this result
func (r Result) Add(other decimal.Decimal) Result {
	return Result{v: r.v.Add(other)}
}

// Sub subtracts another value from this result
func (r Result) Sub(other decimal.Decimal) Result {
	return Result{v: r.v.Sub(other)}
}

// Mul multiplies this result by another value
func (r Result) Mul(other decimal.Decimal) Result {
	return Result{v: r.v.Mul(other)}
}

// Div divides this result by another value
func (r Result) Div(other decimal.Decimal, precision int32) Result {
	return Result{v: r.v.DivRound(other, precision)}
}

// DivTrunc truncates the division
func (r Result) DivTrunc(other decimal.Decimal, precision int32) Result {
	return Result{v: r.v.Div(other).Truncate(precision)}
}
