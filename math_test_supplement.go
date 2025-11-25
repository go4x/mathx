package mathx

import (
	"math"
	"testing"

	"github.com/shopspring/decimal"
)

// ========== Safe 函数测试 ==========

func TestAddSafe(t *testing.T) {
	tests := []struct {
		name     string
		a        decimal.Decimal
		b        decimal.Decimal
		expected string
	}{
		{"positive addition", decimal.NewFromInt(3), decimal.NewFromInt(2), "5"},
		{"decimal addition", decimal.RequireFromString("0.1"), decimal.RequireFromString("0.2"), "0.3"},
		{"negative addition", decimal.RequireFromString("-3.14"), decimal.RequireFromString("2.86"), "-0.28"},
		{"zero addition", decimal.Zero, decimal.RequireFromString("5.5"), "5.5"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddSafe(tt.a, tt.b)
			if got := result.String(); got != tt.expected {
				t.Errorf("AddSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSubSafe(t *testing.T) {
	tests := []struct {
		name     string
		a        decimal.Decimal
		b        decimal.Decimal
		expected string
	}{
		{"positive subtraction", decimal.NewFromInt(5), decimal.NewFromInt(2), "3"},
		{"negative result", decimal.NewFromInt(2), decimal.NewFromInt(5), "-3"},
		{"zero result", decimal.NewFromInt(5), decimal.NewFromInt(5), "0"},
		{"decimal precision", decimal.RequireFromString("0.3"), decimal.RequireFromString("0.1"), "0.2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubSafe(tt.a, tt.b)
			if got := result.String(); got != tt.expected {
				t.Errorf("SubSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMulSafe(t *testing.T) {
	tests := []struct {
		name     string
		a        decimal.Decimal
		b        decimal.Decimal
		expected string
	}{
		{"positive multiplication", decimal.NewFromInt(3), decimal.NewFromInt(4), "12"},
		{"negative multiplication", decimal.RequireFromString("-3"), decimal.NewFromInt(4), "-12"},
		{"zero multiplication", decimal.NewFromInt(5), decimal.Zero, "0"},
		{"decimal multiplication", decimal.RequireFromString("0.1"), decimal.RequireFromString("0.2"), "0.02"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MulSafe(tt.a, tt.b)
			if got := result.String(); got != tt.expected {
				t.Errorf("MulSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDivTruncSafe(t *testing.T) {
	tests := []struct {
		name      string
		a         decimal.Decimal
		b         decimal.Decimal
		precision int32
		expected  string
	}{
		{"simple truncation", decimal.NewFromInt(10), decimal.NewFromInt(3), 2, "3.33"},
		{"negative truncation", decimal.RequireFromString("-10"), decimal.NewFromInt(3), 2, "-3.33"},
		{"zero precision", decimal.NewFromInt(10), decimal.NewFromInt(3), 0, "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DivTruncSafe(tt.a, tt.b, tt.precision)
			if got := result.String(); got != tt.expected {
				t.Errorf("DivTruncSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestRoundSafe(t *testing.T) {
	tests := []struct {
		name      string
		a         decimal.Decimal
		precision int32
		expected  string
	}{
		{"round up", decimal.RequireFromString("3.145"), 2, "3.15"},
		{"round down", decimal.RequireFromString("3.144"), 2, "3.14"},
		{"round to integer", decimal.RequireFromString("3.6"), 0, "4"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RoundSafe(tt.a, tt.precision)
			if got := result.String(); got != tt.expected {
				t.Errorf("RoundSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestTruncateSafe(t *testing.T) {
	tests := []struct {
		name      string
		a         decimal.Decimal
		precision int32
		expected  string
	}{
		{"truncate positive", decimal.RequireFromString("3.145"), 2, "3.14"},
		{"truncate negative", decimal.RequireFromString("-3.145"), 2, "-3.14"},
		{"truncate to integer", decimal.RequireFromString("3.9"), 0, "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TruncateSafe(tt.a, tt.precision)
			if got := result.String(); got != tt.expected {
				t.Errorf("TruncateSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAbsSafe(t *testing.T) {
	tests := []struct {
		name     string
		a        decimal.Decimal
		expected string
	}{
		{"positive number", decimal.RequireFromString("3.14"), "3.14"},
		{"negative number", decimal.RequireFromString("-3.14"), "3.14"},
		{"zero", decimal.Zero, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AbsSafe(tt.a)
			if got := result.String(); got != tt.expected {
				t.Errorf("AbsSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCeilSafe(t *testing.T) {
	tests := []struct {
		name     string
		a        decimal.Decimal
		expected string
	}{
		{"positive decimal", decimal.RequireFromString("3.2"), "4"},
		{"negative decimal", decimal.RequireFromString("-3.2"), "-3"},
		{"integer", decimal.NewFromInt(3), "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CeilSafe(tt.a)
			if got := result.String(); got != tt.expected {
				t.Errorf("CeilSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestFloorSafe(t *testing.T) {
	tests := []struct {
		name     string
		a        decimal.Decimal
		expected string
	}{
		{"positive decimal", decimal.RequireFromString("3.8"), "3"},
		{"negative decimal", decimal.RequireFromString("-3.8"), "-4"},
		{"integer", decimal.NewFromInt(3), "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FloorSafe(tt.a)
			if got := result.String(); got != tt.expected {
				t.Errorf("FloorSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestPowSafe(t *testing.T) {
	tests := []struct {
		name     string
		base     decimal.Decimal
		exponent decimal.Decimal
		expected string
	}{
		{"positive power", decimal.NewFromInt(2), decimal.NewFromInt(3), "8"},
		{"negative power", decimal.NewFromInt(2), decimal.RequireFromString("-1"), "0.5"},
		{"zero power", decimal.NewFromInt(5), decimal.Zero, "1"},
		{"fractional power", decimal.NewFromInt(4), decimal.RequireFromString("0.5"), "2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PowSafe(tt.base, tt.exponent)
			if got := result.String(); got != tt.expected {
				t.Errorf("PowSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestIsEqualSafe(t *testing.T) {
	tests := []struct {
		name      string
		a         decimal.Decimal
		b         decimal.Decimal
		precision int32
		expected  bool
	}{
		{"exact equal", decimal.RequireFromString("3.14"), decimal.RequireFromString("3.14"), 2, true},
		{"very close", decimal.RequireFromString("3.14"), decimal.RequireFromString("3.1400000001"), 8, true},
		{"different", decimal.RequireFromString("3.14"), decimal.RequireFromString("3.15"), 2, false},
		{"zero comparison", decimal.Zero, decimal.Zero, 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEqualSafe(tt.a, tt.b, tt.precision); got != tt.expected {
				t.Errorf("IsEqualSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestClampSafe(t *testing.T) {
	tests := []struct {
		name     string
		value    decimal.Decimal
		min      decimal.Decimal
		max      decimal.Decimal
		expected string
	}{
		{"within range", decimal.NewFromInt(5), decimal.Zero, decimal.NewFromInt(10), "5"},
		{"below min", decimal.RequireFromString("-5"), decimal.Zero, decimal.NewFromInt(10), "0"},
		{"above max", decimal.NewFromInt(15), decimal.Zero, decimal.NewFromInt(10), "10"},
		{"at min", decimal.Zero, decimal.Zero, decimal.NewFromInt(10), "0"},
		{"at max", decimal.NewFromInt(10), decimal.Zero, decimal.NewFromInt(10), "10"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClampSafe(tt.value, tt.min, tt.max)
			if got := result.String(); got != tt.expected {
				t.Errorf("ClampSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSumSafe(t *testing.T) {
	tests := []struct {
		name     string
		values   []decimal.Decimal
		expected string
	}{
		{"positive numbers", []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(3)}, "6"},
		{"negative numbers", []decimal.Decimal{decimal.RequireFromString("-1"), decimal.RequireFromString("-2")}, "-3"},
		{"mixed numbers", []decimal.Decimal{decimal.NewFromInt(1), decimal.RequireFromString("-2"), decimal.NewFromInt(3)}, "2"},
		{"single value", []decimal.Decimal{decimal.NewFromInt(42)}, "42"},
		{"empty slice", []decimal.Decimal{}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SumSafe(tt.values...)
			if got := result.String(); got != tt.expected {
				t.Errorf("SumSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMaxSafe(t *testing.T) {
	tests := []struct {
		name     string
		values   []decimal.Decimal
		expected string
	}{
		{"positive numbers", []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(5), decimal.NewFromInt(3)}, "5"},
		{"negative numbers", []decimal.Decimal{decimal.RequireFromString("-1"), decimal.RequireFromString("-5")}, "-1"},
		{"mixed numbers", []decimal.Decimal{decimal.RequireFromString("-1"), decimal.NewFromInt(5)}, "5"},
		{"single value", []decimal.Decimal{decimal.NewFromInt(42)}, "42"},
		{"empty slice", []decimal.Decimal{}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxSafe(tt.values...)
			if got := result.String(); got != tt.expected {
				t.Errorf("MaxSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestMinSafe(t *testing.T) {
	tests := []struct {
		name     string
		values   []decimal.Decimal
		expected string
	}{
		{"positive numbers", []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(5), decimal.NewFromInt(3)}, "1"},
		{"negative numbers", []decimal.Decimal{decimal.RequireFromString("-1"), decimal.RequireFromString("-5")}, "-5"},
		{"mixed numbers", []decimal.Decimal{decimal.RequireFromString("-1"), decimal.NewFromInt(5)}, "-1"},
		{"single value", []decimal.Decimal{decimal.NewFromInt(42)}, "42"},
		{"empty slice", []decimal.Decimal{}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinSafe(tt.values...)
			if got := result.String(); got != tt.expected {
				t.Errorf("MinSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAverageSafe(t *testing.T) {
	tests := []struct {
		name     string
		values   []decimal.Decimal
		expected string
	}{
		{"simple average", []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(3)}, "2"},
		{"decimal average", []decimal.Decimal{decimal.RequireFromString("1.5"), decimal.RequireFromString("2.5"), decimal.RequireFromString("3.5")}, "2.5"},
		{"single value", []decimal.Decimal{decimal.NewFromInt(42)}, "42"},
		{"empty slice", []decimal.Decimal{}, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AverageSafe(tt.values...)
			if got := result.String(); got != tt.expected {
				t.Errorf("AverageSafe() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// ========== Int64 函数测试 ==========

func TestInt64Div(t *testing.T) {
	tests := []struct {
		name      string
		dividend  int64
		divisor   int64
		precision int32
		expected  float64
	}{
		{"simple division", 10, 2, 2, 5.0},
		{"decimal division", 1, 3, 2, 0.33},
		{"negative division", -10, 2, 2, -5.0},
		{"precision test", 1, 3, 4, 0.3333},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int64Div(tt.dividend, tt.divisor, tt.precision)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Int64Div() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestInt64DivTrunc(t *testing.T) {
	tests := []struct {
		name      string
		dividend  int64
		divisor   int64
		precision int32
		expected  float64
	}{
		{"simple truncation", 10, 3, 2, 3.33},
		{"negative truncation", -10, 3, 2, -3.33},
		{"zero precision", 10, 3, 0, 3.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int64DivTrunc(tt.dividend, tt.divisor, tt.precision)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Int64DivTrunc() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestInt64MulFloat64(t *testing.T) {
	tests := []struct {
		name         string
		multiplicand int64
		multiplier   float64
		expected     float64
	}{
		{"positive multiplication", 10, 2.5, 25.0},
		{"negative multiplication", -10, 2.5, -25.0},
		{"zero multiplication", 10, 0, 0.0},
		{"decimal multiplication", 100, 0.1, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Int64MulFloat64(tt.multiplicand, tt.multiplier)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("Int64MulFloat64() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// ========== 格式化函数测试 ==========

func TestFormatMoney(t *testing.T) {
	tests := []struct {
		name          string
		amount        float64
		decimalPlaces int32
		expected      string
	}{
		{"thousands separator", 1234567.89, 2, "1,234,567.89"},
		{"no thousands separator", 123.45, 2, "123.45"},
		{"integer", 1000, 0, "1,000"},
		{"negative", -1234567.89, 2, "-1,234,567.89"},
		{"small number", 12.5, 2, "12.50"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatMoney(tt.amount, tt.decimalPlaces)
			if result != tt.expected {
				t.Errorf("FormatMoney() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFormatMoneyInt(t *testing.T) {
	tests := []struct {
		name          string
		amount        int64
		decimalPlaces int32
		expected      string
	}{
		{"thousands separator", 1234567, 2, "1,234,567.00"},
		{"no thousands separator", 123, 2, "123.00"},
		{"integer", 1000, 0, "1,000"},
		{"negative", -1234567, 2, "-1,234,567.00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatMoneyInt(tt.amount, tt.decimalPlaces)
			if result != tt.expected {
				t.Errorf("FormatMoneyInt() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRemoveTrailingZeros(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected string
	}{
		{"trailing zeros", 3.14000, "3.14"},
		{"no trailing zeros", 3.14, "3.14"},
		{"integer", 42.0, "42"},
		{"zero", 0.0, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveTrailingZeros(tt.value)
			if result != tt.expected {
				t.Errorf("RemoveTrailingZeros() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestRemoveTrailingZerosFixed(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		places   int32
		expected string
	}{
		{"trailing zeros", 3.14000, 4, "3.14"},
		{"no trailing zeros", 3.14, 2, "3.14"},
		{"integer", 42.0, 2, "42"},
		{"zero", 0.0, 2, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveTrailingZerosFixed(tt.value, tt.places)
			if result != tt.expected {
				t.Errorf("RemoveTrailingZerosFixed() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCleanFloat(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected float64
	}{
		{"normal float", 3.14, 3.14},
		{"trailing zeros", 3.14000, 3.14},
		{"zero", 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CleanFloat(tt.value)
			if math.Abs(result-tt.expected) > 1e-10 {
				t.Errorf("CleanFloat() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestCleanFloatString(t *testing.T) {
	tests := []struct {
		name     string
		value    float64
		expected string
	}{
		{"normal float", 3.14, "3.14"},
		{"trailing zeros", 3.14000, "3.14"},
		{"zero", 0.0, "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CleanFloatString(tt.value)
			if result != tt.expected {
				t.Errorf("CleanFloatString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// ========== Result.DivTrunc 测试 ==========

func TestResult_DivTrunc(t *testing.T) {
	tests := []struct {
		name      string
		value     float64
		other     decimal.Decimal
		precision int32
		expected  string
	}{
		{"simple truncation", 10.0, decimal.NewFromInt(3), 2, "3.33"},
		{"negative truncation", -10.0, decimal.NewFromInt(3), 2, "-3.33"},
		{"zero precision", 10.0, decimal.NewFromInt(3), 0, "3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewResult(tt.value)
			quotient := result.DivTrunc(tt.other, tt.precision)
			if got := quotient.String(); got != tt.expected {
				t.Errorf("Result.DivTrunc() = %v, want %v", got, tt.expected)
			}
		})
	}
}

