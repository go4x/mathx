# MathX - High Precision Decimal Math Library

A high-precision decimal math library for Go that provides accurate floating-point arithmetic operations using the `shopspring/decimal` library under the hood. MathX offers a clean, chainable API for mathematical operations with precise decimal calculations.

## Features

- **High Precision**: Uses `shopspring/decimal` for accurate decimal arithmetic
- **Chainable API**: Fluent interface for mathematical operations
- **Type Safety**: Compile-time type checking with generic support
- **Comprehensive**: Basic operations, statistical functions, formatting utilities
- **Safe Functions**: High-precision functions that accept `decimal.Decimal` directly
- **Zero Dependencies**: Only depends on `shopspring/decimal` and standard library

## Installation

```bash
go get github.com/go4x/mathx
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/go4x/mathx"
    "github.com/shopspring/decimal"
)

func main() {
    // Basic arithmetic with precision
    result := mathx.Add(0.1, 0.2)
    fmt.Println(result) // 0.3 (not 0.30000000000000004)
    
    // Chainable operations
    formatted := mathx.Add(0.1, 0.2).
        Add(decimal.NewFromFloat(10)).
        Div(decimal.NewFromFloat(3), 2).
        Round(2).
        ToString()
    fmt.Println(formatted) // "1.00"
    
    // Money formatting
    price := mathx.Mul(99.99, 1.15).
        Round(2).
        FormatMoney(2)
    fmt.Printf("Price: $%s\n", price) // Price: $114.99
    
    // High precision with Safe functions
    a, _ := decimal.NewFromString("0.1")
    b, _ := decimal.NewFromString("0.2")
    result := mathx.AddSafe(a, b)
    fmt.Println(result.String()) // "0.3"
}
```

## Core Concepts

### Result Type

All mathematical operations return a `Result` type that can be used in two ways:

1. **Direct usage**: Returns `float64` when used directly
2. **Chainable methods**: Call methods for formatting and further operations

```go
// Direct usage
sum := mathx.Add(1, 2) // Returns Result, but can be used as float64
fmt.Println(sum) // 3

// Chainable usage
str := mathx.Add(1, 2).ToString() // "3"
```

### Creating Results

```go
// From float64
result := mathx.NewResult(3.14)

// From string (preserves precision)
result, err := mathx.NewResultFromString("3.14159265358979323846")
```

## API Reference

### Basic Arithmetic

#### Add
```go
func Add(a, b float64) Result
```
Adds two numbers with decimal precision.

```go
result := mathx.Add(0.1, 0.2) // 0.3
```

#### Sub
```go
func Sub(a, b float64) Result
```
Subtracts two numbers with decimal precision.

```go
result := mathx.Sub(5.0, 2.0) // 3.0
```

#### Mul
```go
func Mul(a, b float64) Result
```
Multiplies two numbers with decimal precision.

```go
result := mathx.Mul(3.0, 4.0) // 12.0
```

#### Div
```go
func Div(a, b float64, precision int32) Result
```
Divides two numbers with specified precision.

```go
result := mathx.Div(10.0, 3.0, 2) // 3.33
```

#### DivTrunc
```go
func DivTrunc(a, b float64, precision int32) Result
```
Divides two numbers and truncates to specified precision.

```go
result := mathx.DivTrunc(10.0, 3.0, 2) // 3.33
```

### Safe Functions (High Precision)

Safe functions accept `decimal.Decimal` directly, avoiding float64 conversion precision loss:

#### AddSafe
```go
func AddSafe(a, b decimal.Decimal) Result
```
High precision addition using decimal.Decimal.

```go
a, _ := decimal.NewFromString("0.1")
b, _ := decimal.NewFromString("0.2")
result := mathx.AddSafe(a, b) // "0.3"
```

#### SubSafe
```go
func SubSafe(a, b decimal.Decimal) Result
```
High precision subtraction.

#### MulSafe
```go
func MulSafe(a, b decimal.Decimal) Result
```
High precision multiplication.

#### DivSafe
```go
func DivSafe(a, b decimal.Decimal, precision int32) Result
```
High precision division.

#### DivTruncSafe
```go
func DivTruncSafe(a, b decimal.Decimal, precision int32) Result
```
High precision truncating division.

#### RoundSafe
```go
func RoundSafe(a decimal.Decimal, precision int32) Result
```
High precision rounding.

#### TruncateSafe
```go
func TruncateSafe(a decimal.Decimal, precision int32) Result
```
High precision truncation.

#### AbsSafe
```go
func AbsSafe(a decimal.Decimal) decimal.Decimal
```
High precision absolute value.

#### CeilSafe
```go
func CeilSafe(a decimal.Decimal) decimal.Decimal
```
High precision ceiling.

#### FloorSafe
```go
func FloorSafe(a decimal.Decimal) decimal.Decimal
```
High precision floor.

#### PowSafe
```go
func PowSafe(a decimal.Decimal, exponent decimal.Decimal) decimal.Decimal
```
High precision power.

#### IsEqualSafe
```go
func IsEqualSafe(a, b decimal.Decimal, precision int32) bool
```
High precision comparison.

#### ClampSafe
```go
func ClampSafe(a decimal.Decimal, min, max decimal.Decimal) decimal.Decimal
```
High precision clamping.

### Integer Functions

#### Int64Div
```go
func Int64Div(dividend, divisor int64, precision int32) float64
```
Divides two int64 values with specified precision.

```go
result := mathx.Int64Div(1, 3, 4) // 0.3333
```

#### Int64DivTrunc
```go
func Int64DivTrunc(dividend, divisor int64, precision int32) float64
```
Truncates the division of two int64 values.

```go
result := mathx.Int64DivTrunc(10, 3, 2) // 3.33
```

#### Int64MulFloat64
```go
func Int64MulFloat64(multiplicand int64, multiplier float64) float64
```
Multiplies int64 and float64 using decimal precision.

```go
result := mathx.Int64MulFloat64(100, 0.1) // 10.0
```

### Mathematical Functions

#### Round
```go
func Round(value float64, precision int32) Result
```
Rounds a number to specified decimal places.

```go
result := mathx.Round(3.145, 2) // 3.15
```

#### Truncate
```go
func Truncate(value float64, precision int32) Result
```
Truncates a number to specified decimal places.

```go
result := mathx.Truncate(3.145, 2) // 3.14
```

#### Abs
```go
func Abs(value float64) float64
```
Returns the absolute value of a number.

```go
result := mathx.Abs(-3.14) // 3.14
```

#### Ceil
```go
func Ceil(value float64) float64
```
Returns the smallest integer greater than or equal to the value.

```go
result := mathx.Ceil(3.2) // 4.0
```

#### Floor
```go
func Floor(value float64) float64
```
Returns the largest integer less than or equal to the value.

```go
result := mathx.Floor(3.8) // 3.0
```

#### Pow
```go
func Pow(base, exponent float64) float64
```
Raises a number to the power of another.

```go
result := mathx.Pow(2.0, 3.0) // 8.0
```

#### Sqrt
```go
func Sqrt(value float64) float64
```
Returns the square root of a number.

```go
result := mathx.Sqrt(16.0) // 4.0
```

#### IsEqual
```go
func IsEqual(a, b float64, precision int32) bool
```
Checks if two numbers are equal within specified precision.

```go
equal := mathx.IsEqual(3.14, 3.1400001, 6) // true
```

### Statistical Functions

#### Average
```go
func Average[T constraints.Integer | constraints.Float](ns ...T) float64
```
Calculates the average of a slice of numbers.

```go
avg := mathx.Average(1, 2, 3, 4, 5) // 3.0
```

#### AverageSafe
```go
func AverageSafe(ds ...decimal.Decimal) decimal.Decimal
```
High precision average calculation.

```go
values := []decimal.Decimal{
    decimal.NewFromInt(85),
    decimal.NewFromInt(92),
    decimal.NewFromInt(78),
}
avg := mathx.AverageSafe(values...) // 85
```

#### StandardDeviation
```go
func StandardDeviation[T constraints.Integer | constraints.Float](ns ...T) float64
```
Calculates the standard deviation of a slice of numbers.

```go
std := mathx.StandardDeviation(1, 2, 3, 4, 5) // 1.4142135623730951
```

#### Max
```go
func Max[T constraints.Ordered](ns ...T) T
```
Returns the maximum value from a slice of numbers.

```go
max := mathx.Max(1, 5, 3, 9, 2) // 9.0
```

#### MaxSafe
```go
func MaxSafe(ds ...decimal.Decimal) decimal.Decimal
```
High precision maximum value.

```go
values := []decimal.Decimal{
    decimal.RequireFromString("3.14"),
    decimal.RequireFromString("2.71"),
}
max := mathx.MaxSafe(values...) // 3.14
```

#### Min
```go
func Min[T constraints.Ordered](ns ...T) T
```
Returns the minimum value from a slice of numbers.

```go
min := mathx.Min(1, 5, 3, 9, 2) // 1.0
```

#### MinSafe
```go
func MinSafe(ds ...decimal.Decimal) decimal.Decimal
```
High precision minimum value.

#### Sum
```go
func Sum[T constraints.Integer | constraints.Float](ns ...T) T
```
Returns the sum of a slice of numbers.

```go
sum := mathx.Sum(1, 2, 3, 4, 5) // 15.0
```

#### SumSafe
```go
func SumSafe(ds ...decimal.Decimal) decimal.Decimal
```
High precision sum calculation.

```go
values := []decimal.Decimal{
    decimal.RequireFromString("0.1"),
    decimal.RequireFromString("0.2"),
    decimal.RequireFromString("0.3"),
}
sum := mathx.SumSafe(values...) // 0.6
```

### Utility Functions

#### Clamp
```go
func Clamp(value, min, max float64) float64
```
Clamps a value between min and max.

```go
result := mathx.Clamp(15.0, 0.0, 10.0) // 10.0
```

#### Lerp
```go
func Lerp(a, b, t float64) float64
```
Performs linear interpolation between two values.

```go
result := mathx.Lerp(0.0, 10.0, 0.5) // 5.0
```

#### ToFixed
```go
func ToFixed(value float64, places int32) float64
```
Formats a number to a fixed number of decimal places.

```go
result := mathx.ToFixed(3.14159, 2) // 3.14
```

#### IsPositive
```go
func IsPositive(value float64) bool
```
Checks if a number is positive.

#### IsNegative
```go
func IsNegative(value float64) bool
```
Checks if a number is negative.

#### Sign
```go
func Sign(value float64) int
```
Returns the sign of a number (-1, 0, or 1).

### String Conversion

#### ToString
```go
func ToString(value float64) string
```
Converts a float64 to string.

```go
str := mathx.ToString(3.14) // "3.14"
```

#### ToStringFixed
```go
func ToStringFixed(value float64, places int32) string
```
Converts a float64 to string with fixed decimal places.

```go
str := mathx.ToStringFixed(3.14159, 2) // "3.14"
```

#### ToStringBank
```go
func ToStringBank(value float64, places int32) string
```
Converts a float64 to string with banker's rounding.

```go
str := mathx.ToStringBank(3.145, 2) // "3.14"
```

### Formatting Functions

#### FormatMoney
```go
func FormatMoney(amount float64, decimalPlaces int32) string
```
Formats a number as currency with thousands separator.

```go
str := mathx.FormatMoney(1234567.89, 2) // "1,234,567.89"
```

#### FormatMoneyInt
```go
func FormatMoneyInt(amount int64, decimalPlaces int32) string
```
Formats an int64 as currency with thousands separator.

```go
str := mathx.FormatMoneyInt(1234567, 2) // "1,234,567.00"
```

#### FormatCurrency
```go
func FormatCurrency(amount float64, decimalPlaces int32) string
```
Formats a number as currency with specified decimal places.

```go
str := mathx.FormatCurrency(123.45, 2) // "123.45"
```

#### RemoveTrailingZeros
```go
func RemoveTrailingZeros(value float64) string
```
Removes trailing zeros from a float64 string representation.

```go
str := mathx.RemoveTrailingZeros(3.14000) // "3.14"
```

#### RemoveTrailingZerosFixed
```go
func RemoveTrailingZerosFixed(value float64, places int32) string
```
Removes trailing zeros from a float64 with fixed decimal places.

```go
str := mathx.RemoveTrailingZerosFixed(3.14000, 4) // "3.14"
```

#### CleanFloat
```go
func CleanFloat(value float64) float64
```
Removes trailing zeros and returns a clean float64.

```go
result := mathx.CleanFloat(3.14000) // 3.14
```

#### CleanFloatString
```go
func CleanFloatString(value float64) string
```
Removes trailing zeros and returns a clean string representation.

```go
str := mathx.CleanFloatString(3.14000) // "3.14"
```

#### ParseFloat
```go
func ParseFloat(s string) (float64, error)
```
Safely parses a string to float64.

```go
value, err := mathx.ParseFloat("3.14")
if err == nil {
    fmt.Println(value) // 3.14
}
```

### Result Methods

The `Result` type provides chainable methods for further operations:

#### Creating Results
```go
func NewResult(value float64) Result
func NewResultFromString(value string) (Result, error)
func (r Result) Decimal() decimal.Decimal
```

#### String Conversion
```go
func (r Result) String() string
func (r Result) ToString() string
func (r Result) ToStringFixed(places int32) string
func (r Result) ToStringBank(places int32) string
```

#### Formatting
```go
func (r Result) FormatMoney(decimalPlaces int32) string
func (r Result) Clean() Result // Removes trailing zeros
```

#### Mathematical Operations
```go
func (r Result) Add(other decimal.Decimal) Result
func (r Result) Sub(other decimal.Decimal) Result
func (r Result) Mul(other decimal.Decimal) Result
func (r Result) Div(other decimal.Decimal, precision int32) Result
func (r Result) DivTrunc(other decimal.Decimal, precision int32) Result
func (r Result) Round(places int32) Result
func (r Result) Truncate(places int32) Result
func (r Result) Abs() Result
func (r Result) Neg() Result
```

#### Value Extraction
```go
func (r Result) Float64() float64
```

## Examples

### Financial Calculations

```go
package main

import (
    "fmt"
    "github.com/go4x/mathx"
)

func main() {
    // Calculate tax
    price := 99.99
    taxRate := 0.08
    tax := mathx.Mul(price, taxRate).Round(2)
    
    // Calculate total
    total := mathx.Add(price, tax.Float64())
    
    // Format as currency
    formatted := total.FormatMoney(2)
    fmt.Printf("Total: $%s\n", formatted) // Total: $107.99
}
```

### High Precision Calculations

```go
package main

import (
    "fmt"
    "github.com/go4x/mathx"
    "github.com/shopspring/decimal"
)

func main() {
    // Use Safe functions for maximum precision
    a, _ := decimal.NewFromString("0.1")
    b, _ := decimal.NewFromString("0.2")
    result := mathx.AddSafe(a, b)
    fmt.Println(result.String()) // "0.3"
    
    // Chain Safe operations
    values := []decimal.Decimal{
        decimal.RequireFromString("0.1"),
        decimal.RequireFromString("0.2"),
        decimal.RequireFromString("0.3"),
    }
    sum := mathx.SumSafe(values...)
    fmt.Println(sum.String()) // "0.6"
}
```

### Statistical Analysis

```go
package main

import (
    "fmt"
    "github.com/go4x/mathx"
)

func main() {
    scores := []float64{85, 92, 78, 96, 88}
    
    avg := mathx.Average(scores...)
    std := mathx.StandardDeviation(scores...)
    
    fmt.Printf("Average: %.2f\n", avg)
    fmt.Printf("Standard Deviation: %.2f\n", std)
}
```

### Chainable Operations

```go
package main

import (
    "fmt"
    "github.com/go4x/mathx"
    "github.com/shopspring/decimal"
)

func main() {
    // Complex calculation with formatting
    result := mathx.Add(0.1, 0.2).
        Add(decimal.NewFromFloat(100)).
        Div(decimal.NewFromFloat(3), 4).
        Round(2).
        Clean().
        ToString()
    
    fmt.Println(result) // "10.00"
}
```

## Performance Considerations

- MathX uses `shopspring/decimal` internally for high precision
- String conversions may have slight overhead compared to native float64 operations
- For performance-critical applications, consider using the direct `float64` return values
- Chainable operations create new `Result` instances, so use judiciously in tight loops
- Safe functions avoid float64 conversion overhead and provide maximum precision

## Error Handling

Most functions return `float64` or `Result` types directly. Functions that can fail (like `ParseFloat` and `NewResultFromString`) return `(value, error)` tuples.

```go
value, err := mathx.ParseFloat("invalid")
if err != nil {
    // Handle error
}

result, err := mathx.NewResultFromString("invalid")
if err != nil {
    // Handle error
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT.

## Acknowledgments

- Built on top of the excellent `shopspring/decimal` library
- Inspired by the need for precise decimal arithmetic in Go applications
