package mathx_test

import (
	"fmt"

	"github.com/go4x/mathx"
	"github.com/shopspring/decimal"
)

func ExampleAdd() {
	// Basic addition with precision
	result := mathx.Add(0.1, 0.2)
	fmt.Printf("0.1 + 0.2 = %.10f\n", result.Float64())
	// Output: 0.1 + 0.2 = 0.3000000000
}

func ExampleAdd_chainable() {
	// Chainable operations
	result := mathx.Add(0.1, 0.2).
		Mul(decimal.NewFromFloat(10)).
		Div(decimal.NewFromFloat(3), 2).
		Round(2).
		ToStringFixed(2)
	fmt.Printf("Result: %s\n", result)
	// Output: Result: 1.00
}

func ExampleMul() {
	// Money formatting with thousands separator
	price := mathx.Mul(99.99, 1.15).
		Round(2).
		FormatMoney(2)
	fmt.Printf("Price: $%s\n", price)
	// Output: Price: $114.99
}

func ExampleAverage() {
	// Statistical calculations
	scores := []float64{85, 92, 78, 96, 88}

	avg := mathx.Average(scores...)
	std := mathx.StandardDeviation(scores...)

	fmt.Printf("Average: %.2f\n", avg)
	fmt.Printf("Standard Deviation: %.2f\n", std)
	// Output:
	// Average: 87.80
	// Standard Deviation: 6.87
}

func ExampleAdd_clean() {
	// Remove trailing zeros
	result := mathx.Add(0.1, 0.2).
		Clean().
		ToString()
	fmt.Printf("Clean result: %s\n", result)
	// Output: Clean result: 0.3
}

func ExampleToString() {
	// Various string conversion methods
	value := 3.14159

	fmt.Printf("ToString: %s\n", mathx.ToString(value))
	fmt.Printf("ToStringFixed(2): %s\n", mathx.ToStringFixed(value, 2))
	fmt.Printf("ToStringBank(2): %s\n", mathx.ToStringBank(value, 2))
	// Output:
	// ToString: 3.14159
	// ToStringFixed(2): 3.14
	// ToStringBank(2): 3.14
}

func ExampleAbs() {
	// Mathematical functions
	fmt.Printf("Abs(-3.14): %.2f\n", mathx.Abs(-3.14))
	fmt.Printf("Ceil(3.14): %.0f\n", mathx.Ceil(3.14))
	fmt.Printf("Floor(3.14): %.0f\n", mathx.Floor(3.14))
	fmt.Printf("Pow(2, 3): %.0f\n", mathx.Pow(2, 3))
	fmt.Printf("Sqrt(16): %.0f\n", mathx.Sqrt(16))
	// Output:
	// Abs(-3.14): 3.14
	// Ceil(3.14): 4
	// Floor(3.14): 3
	// Pow(2, 3): 8
	// Sqrt(16): 4
}

func ExampleClamp() {
	// Utility functions
	value := 15.0
	min := 0.0
	max := 10.0

	clamped := mathx.Clamp(value, min, max)
	fmt.Printf("Clamp(%.0f, %.0f, %.0f): %.0f\n", value, min, max, clamped)

	// Linear interpolation
	lerp := mathx.Lerp(0.0, 10.0, 0.5)
	fmt.Printf("Lerp(0, 10, 0.5): %.1f\n", lerp)
	// Output:
	// Clamp(15, 0, 10): 10
	// Lerp(0, 10, 0.5): 5.0
}

func ExampleParseFloat() {
	// Safe string parsing
	value, err := mathx.ParseFloat("3.14")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Parsed value: %.2f\n", value)
	// Output: Parsed value: 3.14
}

func ExampleAddSafe() {
	// High precision addition using decimal.Decimal
	a, _ := decimal.NewFromString("0.1")
	b, _ := decimal.NewFromString("0.2")
	result := mathx.AddSafe(a, b)
	fmt.Printf("0.1 + 0.2 = %s\n", result.String())
	// Output: 0.1 + 0.2 = 0.3
}

func ExampleInt64Div() {
	// Integer division with precision
	result := mathx.Int64Div(1, 3, 4)
	fmt.Printf("1 / 3 (4 decimal places) = %.4f\n", result)
	// Output: 1 / 3 (4 decimal places) = 0.3333
}

func ExampleFormatMoney() {
	// Format currency with thousands separator
	amount := 1234567.89
	formatted := mathx.FormatMoney(amount, 2)
	fmt.Printf("Amount: $%s\n", formatted)
	// Output: Amount: $1,234,567.89
}

func ExampleFormatMoneyInt() {
	// Format integer as currency
	amount := int64(1234567)
	formatted := mathx.FormatMoneyInt(amount, 2)
	fmt.Printf("Amount: $%s\n", formatted)
	// Output: Amount: $1,234,567.00
}

func ExampleSumSafe() {
	// High precision sum using decimal.Decimal
	values := []decimal.Decimal{
		decimal.RequireFromString("0.1"),
		decimal.RequireFromString("0.2"),
		decimal.RequireFromString("0.3"),
	}
	sum := mathx.SumSafe(values...)
	fmt.Printf("Sum: %s\n", sum.String())
	// Output: Sum: 0.6
}

func ExampleMaxSafe() {
	// High precision max using decimal.Decimal
	values := []decimal.Decimal{
		decimal.RequireFromString("3.14"),
		decimal.RequireFromString("2.71"),
		decimal.RequireFromString("1.41"),
	}
	max := mathx.MaxSafe(values...)
	fmt.Printf("Max: %s\n", max.String())
	// Output: Max: 3.14
}

func ExampleAverageSafe() {
	// High precision average using decimal.Decimal
	values := []decimal.Decimal{
		decimal.NewFromInt(85),
		decimal.NewFromInt(92),
		decimal.NewFromInt(78),
	}
	avg := mathx.AverageSafe(values...)
	fmt.Printf("Average: %s\n", avg.String())
	// Output: Average: 85
}

func ExampleClampSafe() {
	// High precision clamp using decimal.Decimal
	value := decimal.RequireFromString("15.5")
	min := decimal.NewFromInt(0)
	max := decimal.NewFromInt(10)
	clamped := mathx.ClampSafe(value, min, max)
	fmt.Printf("Clamped: %s\n", clamped.String())
	// Output: Clamped: 10
}

func ExampleRemoveTrailingZeros() {
	// Remove trailing zeros from float64
	value := 3.14000
	cleaned := mathx.RemoveTrailingZeros(value)
	fmt.Printf("Cleaned: %s\n", cleaned)
	// Output: Cleaned: 3.14
}

func ExampleCleanFloatString() {
	// Get clean string representation
	value := 3.14000
	cleaned := mathx.CleanFloatString(value)
	fmt.Printf("Cleaned: %s\n", cleaned)
	// Output: Cleaned: 3.14
}

func ExampleIsEqualSafe() {
	// High precision comparison
	a := decimal.RequireFromString("3.14")
	b := decimal.RequireFromString("3.1400000001")
	equal := mathx.IsEqualSafe(a, b, 8)
	fmt.Printf("Equal (8 decimal places): %v\n", equal)
	// Output: Equal (8 decimal places): true
}
