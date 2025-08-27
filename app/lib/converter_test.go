package lib

import (
	"reflect"
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/uuid"
)

func TestMD5FromInt(t *testing.T) {
	value := 1
	gen := MD5FromInt(value)
	gen2 := MD5FromInt(value)
	utils.AssertEqual(t, gen2, gen)
}

func TestMD5FromString(t *testing.T) {
	value := "development usage"
	gen := MD5FromString(value)
	gen2 := MD5FromString(value)
	utils.AssertEqual(t, gen2, gen)
}

func TestSHA1FromString(t *testing.T) {
	value := "development usage"
	gen := SHA1FromString(value)
	gen2 := SHA1FromString(value)
	utils.AssertEqual(t, gen2, gen)
}

func TestSHA256FromString(t *testing.T) {
	value := "development usage"
	gen := SHA256FromString(value)
	gen2 := SHA256FromString(value)
	utils.AssertEqual(t, gen2, gen)
}

func TestIntToStr(t *testing.T) {
	value := 1
	res := IntToStr(value)
	utils.AssertEqual(t, "1", res)
}

func TestInt64ToStr(t *testing.T) {
	value := int64(1)
	res := Int64ToStr(value)
	utils.AssertEqual(t, "1", res)
}

func TestStrToInt(t *testing.T) {
	value := "1"
	res := StrToInt(value)
	utils.AssertEqual(t, 1, res)
}

func TestStrToInt64(t *testing.T) {
	value := "1"
	res := StrToInt64(value)
	utils.AssertEqual(t, int64(1), res)
}

func TestStrToFloat(t *testing.T) {
	value := "1"
	res := StrToFloat(value)
	utils.AssertEqual(t, float64(1), res)
}

func TestStrToBool(t *testing.T) {
	value := "true"
	res := StrToBool(value)
	utils.AssertEqual(t, true, res)
}

func TestFloatToStr(t *testing.T) {
	value := 1.2
	res := FloatToStr(value, 6)
	utils.AssertEqual(t, "1.200000", res)
}

func TestFloatToFormattedStr(t *testing.T) {
	mapTest := map[float64]string{
		10000:    "10,000",
		12500.12: "12,501",
		500:      "500",
		10.1:     "11",
		19191919: "19,191,919",
	}

	for inp, expected := range mapTest {
		utils.AssertEqual(t, expected, FloatToFormattedStr(inp), "Test FloatToFormattedStr")
	}
}

func TestStrToTime(t *testing.T) {
	expect := time.Date(2021, 5, 19, 11, 56, 30, 0, time.UTC)
	res := StrToTime("2021-05-19 11:56:30", "2006-01-02 15:04:05")
	utils.AssertEqual(t, expect, res, "Test StrToTime")

	expect = time.Date(2021, 5, 19, 0, 0, 0, 0, time.UTC)
	res = StrToTime("2021-05-19", "2006-01-02")
	utils.AssertEqual(t, expect, res, "Test StrToTime with custom layout")

	expect = time.Date(2021, 5, 19, 0, 0, 0, 0, time.UTC)
	res = StrToTime("2021-05-19")
	utils.AssertEqual(t, expect, res, "Test StrToTime without custom layout")
}

func TestStrToStrfmtDate(t *testing.T) {
	expect := strfmt.Date(time.Date(2023, 8, 15, 0, 0, 0, 0, time.UTC))
	res := StrToStrfmtDate("2023-08-15")
	utils.AssertEqual(t, expect, res, "Test StrToStrfmtDate without custom layout")
}

func TestStrToStrfmtDateTime(t *testing.T) {
	expect := strfmt.DateTime(time.Date(2023, 8, 15, 10, 8, 15, 0, time.UTC))
	res := StrToStrfmtDateTime("2023-08-15 10:08:15")
	utils.AssertEqual(t, expect, res, "Test StrToStrfmtDateTime without custom layout")
}

func TestIntSliceToStr(t *testing.T) {
	value := []int{1, 2, 3, 4}
	res := IntSliceToStr(value, ",")
	utils.AssertEqual(t, "1,2,3,4", res)
}

func TestStrSliceToStr(t *testing.T) {
	value := []string{"active", "inactive", "suspend"}
	res := StrSliceToStr(value, ",")
	utils.AssertEqual(t, "active,inactive,suspend", res)
}

func TestUUIDSliceToStr(t *testing.T) {
	id1 := *GenUUID()
	id2 := *GenUUID()
	id3 := *GenUUID()

	str := id1.String() + "," + id2.String() + "," + id3.String()
	value := []uuid.UUID{id1, id2, id3}
	res := UUIDSliceToStr(value, ",")
	utils.AssertEqual(t, str, res)
}

func TestUUIDSliceToStrSlice(t *testing.T) {
	uuid1 := uuid.New()
	uuid2 := uuid.New()
	uuid3 := uuid.New()

	type args struct {
		listUUID []uuid.UUID
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
	}{
		{
			name: "slice uuid converted",
			args: args{
				listUUID: []uuid.UUID{
					uuid1,
					uuid2,
					uuid3,
				},
			},
			wantResult: []string{
				uuid1.String(),
				uuid2.String(),
				uuid3.String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := UUIDSliceToStrSlice(tt.args.listUUID); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ConvertSliceUUIDToSliceStr() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestRemoveLeadingZeros(t *testing.T) {
	str := "026"
	res := RemoveLeadingZeros(str)
	utils.AssertEqual(t, true, len(res) > 0)
}

func TestForceStr(t *testing.T) {
	str1 := "abcd"

	type args struct {
		input *string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "filled input == result",
			args: args{
				input: &str1,
			},
			wantResult: str1,
		},
		{
			name: "nil input; result = empty string",
			args: args{
				input: nil,
			},
			wantResult: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ForceStr(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("ForceStr() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestForceInt(t *testing.T) {
	num1 := 123

	type args struct {
		input *int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "filled input == result",
			args: args{
				input: &num1,
			},
			wantResult: num1,
		},
		{
			name: "nil input; result = 0",
			args: args{
				input: nil,
			},
			wantResult: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ForceInt(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("ForceInt() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestForceInt64(t *testing.T) {
	num1 := int64(123)

	type args struct {
		input *int64
	}
	tests := []struct {
		name       string
		args       args
		wantResult int64
	}{
		{
			name: "filled input == result",
			args: args{
				input: &num1,
			},
			wantResult: num1,
		},
		{
			name: "nil input; result = 0",
			args: args{
				input: nil,
			},
			wantResult: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ForceInt64(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("ForceInt64() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestForceBool(t *testing.T) {
	cond1 := true
	cond2 := false

	type args struct {
		input *bool
	}
	tests := []struct {
		name       string
		args       args
		wantResult bool
	}{
		{
			name: "filled input == result",
			args: args{
				input: &cond1,
			},
			wantResult: cond1,
		},
		{
			name: "filled input == result",
			args: args{
				input: &cond2,
			},
			wantResult: cond2,
		},
		{
			name: "nil input; result = false",
			args: args{
				input: nil,
			},
			wantResult: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ForceBool(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("ForceBool() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestForceFloat64(t *testing.T) {
	num1 := float64(123.00)

	type args struct {
		input *float64
	}
	tests := []struct {
		name       string
		args       args
		wantResult float64
	}{
		{
			name: "filled input == result",
			args: args{
				input: &num1,
			},
			wantResult: num1,
		},
		{
			name: "nil input; result = 0",
			args: args{
				input: nil,
			},
			wantResult: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ForceFloat64(tt.args.input); gotResult != tt.wantResult {
				t.Errorf("ForceFloat64() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestForceTime(t *testing.T) {
	time1 := time.Now()

	type args struct {
		input *time.Time
	}
	tests := []struct {
		name       string
		args       args
		wantResult time.Time
	}{
		{
			name: "filled input == result",
			args: args{
				input: &time1,
			},
			wantResult: time1,
		},
		{
			name: "nil input; result = time.Time{}",
			args: args{
				input: nil,
			},
			wantResult: time.Time{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ForceTime(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ForceTime() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestForceUUID(t *testing.T) {
	uuid1 := uuid.New()

	type args struct {
		input *uuid.UUID
	}
	tests := []struct {
		name       string
		args       args
		wantResult uuid.UUID
	}{
		{
			name: "filled input == result",
			args: args{
				input: &uuid1,
			},
			wantResult: uuid1,
		},
		{
			name: "nil input; result = uuid.Nil",
			args: args{
				input: nil,
			},
			wantResult: uuid.Nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ForceUUID(tt.args.input); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ForceUUID() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestReplaceSpecialCharacters(t *testing.T) {
	type args struct {
		str         string
		replaceWith string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "normal case, replace with underscore",
			args: args{
				str:         "SQ_66WI54_AIR-1",
				replaceWith: "_",
			},
			wantResult: "SQ_66WI54_AIR_1",
		},
		{
			name: "normal case, replace with space",
			args: args{
				str:         "SQ_66WI54_AIR-1",
				replaceWith: " ",
			},
			wantResult: "SQ 66WI54 AIR 1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ReplaceSpecialCharacters(tt.args.str, tt.args.replaceWith); gotResult != tt.wantResult {
				t.Errorf("ReplaceSpecialCharacters() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
