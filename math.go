package mathx

import (
	"fmt"
	"math"
	"strings"

	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
)

// Add adds two float64 values using decimal precision and returns a Result
func Add(a, b float64) Result {
	result := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b))
	return Result{v: result}
}

// AddSafe adds two decimal values and returns a Result
func AddSafe(a, b decimal.Decimal) Result {
	result := a.Add(b)
	return Result{v: result}
}

// Sub subtracts two float64 values using decimal precision and returns a Result
func Sub(a, b float64) Result {
	result := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b))
	return Result{v: result}
}

// SubSafe subtracts two decimal values and returns a Result
func SubSafe(a, b decimal.Decimal) Result {
	result := a.Sub(b)
	return Result{v: result}
}

// Mul multiplies two float64 values using decimal precision and returns a Result
func Mul(a, b float64) Result {
	result := decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b))
	return Result{v: result}
}

// MulSafe multiplies two decimal values and returns a Result
func MulSafe(a, b decimal.Decimal) Result {
	result := a.Mul(b)
	return Result{v: result}
}

// Div divides two float64 values using decimal precision and returns a Result
func Div(a, b float64, precision int32) Result {
	result := decimal.NewFromFloat(a).DivRound(decimal.NewFromFloat(b), precision)
	return Result{v: result}
}

// DivSafe divides two decimal values and returns a Result
func DivSafe(a, b decimal.Decimal, precision int32) Result {
	result := a.DivRound(b, precision)
	return Result{v: result}
}

// DivTrunc truncates the division of two float64 values and returns a Result
func DivTrunc(a, b float64, precision int32) Result {
	result := decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)).Truncate(precision)
	return Result{v: result}
}

// DivTruncSafe truncates the division of two decimal values and returns a Result
func DivTruncSafe(a, b decimal.Decimal, precision int32) Result {
	result := a.Div(b).Truncate(precision)
	return Result{v: result}
}

// Round rounds a float64 to specified precision and returns a Result
func Round(value float64, precision int32) Result {
	result := decimal.NewFromFloat(value).Round(precision)
	return Result{v: result}
}

// RoundSafe rounds a decimal value to specified precision and returns a Result
func RoundSafe(a decimal.Decimal, precision int32) Result {
	result := a.Round(precision)
	return Result{v: result}
}

// Truncate truncates a float64 to specified precision and returns a Result
func Truncate(value float64, precision int32) Result {
	if precision < 0 {
		// For negative precision, truncate to integer places
		// e.g., precision -1 means truncate to tens place
		multiplier := decimal.NewFromFloat(math.Pow(10, float64(-precision)))
		result := decimal.NewFromFloat(value).Div(multiplier).Truncate(0).Mul(multiplier)
		return Result{v: result}
	}
	result := decimal.NewFromFloat(value).Truncate(precision)
	return Result{v: result}
}

// TruncateSafe truncates a decimal value to specified precision and returns a Result
func TruncateSafe(a decimal.Decimal, precision int32) Result {
	if precision < 0 {
		// For negative precision, truncate to integer places
		// e.g., precision -1 means truncate to tens place
		multiplier := decimal.NewFromFloat(math.Pow(10, float64(-precision)))
		result := a.Div(multiplier).Truncate(0).Mul(multiplier)
		return Result{v: result}
	}
	result := a.Truncate(precision)
	return Result{v: result}
}

// Int64Div divides two int64 values with specified precision
func Int64Div(dividend, divisor int64, precision int32) float64 {
	result := decimal.NewFromInt(dividend).DivRound(decimal.NewFromInt(divisor), precision)
	f, _ := result.Float64()
	return f
}

// Int64DivTrunc truncates the division of two int64 values
func Int64DivTrunc(dividend, divisor int64, precision int32) float64 {
	result := decimal.NewFromInt(dividend).Div(decimal.NewFromInt(divisor)).Truncate(precision)
	f, _ := result.Float64()
	return f
}

// Int64MulFloat64 multiplies int64 and float64 using decimal precision
func Int64MulFloat64(multiplicand int64, multiplier float64) float64 {
	result := decimal.NewFromInt(multiplicand).Mul(decimal.NewFromFloat(multiplier))
	f, _ := result.Float64()
	return f
}

// Max returns the maximum value from a slice of numbers
func Max[T constraints.Ordered](ns ...T) T {
	if len(ns) == 0 {
		var zero T
		return zero
	}
	max := ns[0]
	for _, n := range ns[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

// Min returns the minimum value from a slice of numbers
func Min[T constraints.Ordered](ns ...T) T {
	if len(ns) == 0 {
		var zero T
		return zero
	}
	min := ns[0]
	for _, n := range ns[1:] {
		if n < min {
			min = n
		}
	}
	return min
}

// Sum returns the sum of a slice of numbers
func Sum[T constraints.Integer | constraints.Float](ns ...T) T {
	var sum T
	for _, n := range ns {
		sum += n
	}
	return sum
}

// SumSafe returns the sum of decimal values
func SumSafe(ds ...decimal.Decimal) decimal.Decimal {
	if len(ds) == 0 {
		return decimal.Zero
	}
	sum := ds[0]
	for _, d := range ds[1:] {
		sum = sum.Add(d)
	}
	return sum
}

// MaxSafe returns the maximum decimal value
func MaxSafe(ds ...decimal.Decimal) decimal.Decimal {
	if len(ds) == 0 {
		return decimal.Zero
	}
	max := ds[0]
	for _, d := range ds[1:] {
		if d.GreaterThan(max) {
			max = d
		}
	}
	return max
}

// MinSafe returns the minimum decimal value
func MinSafe(ds ...decimal.Decimal) decimal.Decimal {
	if len(ds) == 0 {
		return decimal.Zero
	}
	min := ds[0]
	for _, d := range ds[1:] {
		if d.LessThan(min) {
			min = d
		}
	}
	return min
}

// Abs returns the absolute value of a number
func Abs(value float64) float64 {
	if value < 0 {
		return -value
	}
	return value
}

// AbsSafe returns the absolute value of a decimal value
func AbsSafe(a decimal.Decimal) decimal.Decimal {
	return a.Abs()
}

// Ceil returns the smallest integer greater than or equal to the value
func Ceil(value float64) float64 {
	result := decimal.NewFromFloat(value).Ceil()
	f, _ := result.Float64()
	return f
}

// CeilSafe returns the smallest integer greater than or equal to the value of a decimal value
func CeilSafe(a decimal.Decimal) decimal.Decimal {
	return a.Ceil()
}

// Floor returns the largest integer less than or equal to the value
func Floor(value float64) float64 {
	result := decimal.NewFromFloat(value).Floor()
	f, _ := result.Float64()
	return f
}

// FloorSafe returns the largest integer less than or equal to the value of a decimal value
func FloorSafe(a decimal.Decimal) decimal.Decimal {
	return a.Floor()
}

// Pow raises a number to the power of another
func Pow(base, exponent float64) float64 {
	result := decimal.NewFromFloat(base).Pow(decimal.NewFromFloat(exponent))
	f, _ := result.Float64()
	return f
}

// PowSafe raises a number to the power of another of a decimal value
func PowSafe(a decimal.Decimal, exponent decimal.Decimal) decimal.Decimal {
	return a.Pow(exponent)
}

// Sqrt returns the square root of a number
func Sqrt(value float64) float64 {
	return math.Sqrt(value)
}

func IsEqual(a, b float64, precision int32) bool {
	return Abs(a-b) < math.Pow(10, -float64(precision))
}

func IsEqualSafe(a, b decimal.Decimal, precision int32) bool {
	return a.Sub(b).Abs().LessThan(decimal.NewFromFloat(math.Pow(10, -float64(precision))))
}

// Clamp clamps a value between min and max
func Clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ClampSafe clamps a decimal value between min and max
func ClampSafe(a decimal.Decimal, min, max decimal.Decimal) decimal.Decimal {
	if a.LessThan(min) {
		return min
	}
	if a.GreaterThan(max) {
		return max
	}
	return a
}

// Lerp performs linear interpolation between two values
func Lerp(a, b, t float64) float64 {
	return Add(a, Mul(Sub(b, a).Float64(), t).Float64()).Float64()
}

// Average calculates the average of a slice of numbers
func Average[T constraints.Integer | constraints.Float](ns ...T) float64 {
	if len(ns) == 0 {
		return 0
	}
	sum := Sum(ns...)
	return Div(float64(sum), float64(len(ns)), 32).Float64()
}

// AverageSafe calculates the average of a slice of decimal values
func AverageSafe(ds ...decimal.Decimal) decimal.Decimal {
	if len(ds) == 0 {
		return decimal.Zero
	}
	sum := SumSafe(ds...)
	return DivSafe(sum, decimal.NewFromInt(int64(len(ds))), 32).Decimal()
}

// StandardDeviation calculates the standard deviation of a slice of numbers
func StandardDeviation[T constraints.Integer | constraints.Float](ns ...T) float64 {
	if len(ns) == 0 {
		return 0
	}
	if len(ns) == 1 {
		return 0
	}

	avg := Average(ns...)
	var sum float64
	for _, n := range ns {
		diff := Sub(float64(n), avg).Float64()
		sum = Add(sum, Mul(diff, diff).Float64()).Float64()
	}

	// Use sample standard deviation (n-1)
	variance := Div(sum, float64(len(ns)-1), 10).Float64()
	return Sqrt(variance)
}

// FormatCurrency formats a number as currency with specified decimal places
func FormatCurrency(amount float64, decimalPlaces int32) string {
	rounded := Round(amount, decimalPlaces)
	if decimalPlaces == 0 {
		return fmt.Sprintf("%.0f", rounded.Float64())
	}
	format := fmt.Sprintf("%%.%df", decimalPlaces)
	return fmt.Sprintf(format, rounded.Float64())
}

// ParseFloat safely parses a string to float64
func ParseFloat(s string) (float64, error) {
	d, err := decimal.NewFromString(s)
	if err != nil {
		return 0, err
	}
	f, _ := d.Float64()
	return f, nil
}

// ToFixed formats a number to a fixed number of decimal places
func ToFixed(value float64, places int32) float64 {
	return Round(value, places).Float64()
}

// IsPositive checks if a number is positive
func IsPositive(value float64) bool {
	return value > 0
}

// IsNegative checks if a number is negative
func IsNegative(value float64) bool {
	return value < 0
}

// Sign returns the sign of a number (-1, 0, or 1)
func Sign(value float64) int {
	if value > 0 {
		return 1
	}
	if value < 0 {
		return -1
	}
	return 0
}

// ToString converts a float64 to string
func ToString(value float64) string {
	return decimal.NewFromFloat(value).String()
}

// ToStringFixed converts a float64 to string with fixed decimal places
func ToStringFixed(value float64, places int32) string {
	return decimal.NewFromFloat(value).StringFixed(places)
}

// ToStringBank converts a float64 to string with banker's rounding
func ToStringBank(value float64, places int32) string {
	return decimal.NewFromFloat(value).StringFixedBank(places)
}

// FormatMoney formats a number as currency with thousands separator
func FormatMoney(amount float64, decimalPlaces int32) string {
	d := decimal.NewFromFloat(amount)
	rounded := d.Round(decimalPlaces)

	// 转换为字符串并添加千位分隔符
	str := rounded.String()

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

// FormatMoneyInt formats an int64 as currency with thousands separator
func FormatMoneyInt(amount int64, decimalPlaces int32) string {
	d := decimal.NewFromInt(amount)
	rounded := d.Round(decimalPlaces)

	// 转换为字符串并添加千位分隔符
	str := rounded.String()

	// 分离整数和小数部分
	parts := strings.Split(str, ".")
	integerPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	} else {
		decimalPart = ".00"
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

// RemoveTrailingZeros removes trailing zeros from a float64 string representation
func RemoveTrailingZeros(value float64) string {
	str := decimal.NewFromFloat(value).String()

	// 如果包含小数点
	if strings.Contains(str, ".") {
		// 移除尾随的零
		str = strings.TrimRight(str, "0")
		// 如果小数点后没有数字了，也移除小数点
		str = strings.TrimRight(str, ".")
	}

	return str
}

// RemoveTrailingZerosFixed removes trailing zeros from a float64 with fixed decimal places
func RemoveTrailingZerosFixed(value float64, places int32) string {
	str := decimal.NewFromFloat(value).StringFixed(places)

	// 如果包含小数点
	if strings.Contains(str, ".") {
		// 移除尾随的零
		str = strings.TrimRight(str, "0")
		// 如果小数点后没有数字了，也移除小数点
		str = strings.TrimRight(str, ".")
	}

	return str
}

// CleanFloat removes trailing zeros and returns a clean float64
func CleanFloat(value float64) float64 {
	// 先转换为decimal，再转回float64，这样可以去除浮点数精度问题
	d := decimal.NewFromFloat(value)
	f, _ := d.Float64()
	return f
}

// CleanFloatString removes trailing zeros and returns a clean string representation
func CleanFloatString(value float64) string {
	d := decimal.NewFromFloat(value)
	return d.String()
}
