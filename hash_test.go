package stringo

import (
	"strings"
	"testing"
	"time"
)

type defaultTestStruct struct {
	summary        string
	input          interface{}
	expectedOutput interface{}
}

func TestCheckEmail(t *testing.T) {
	testList := []defaultTestStruct{
		{"send empty string", "", false},
		{"send invalid address", "email-gmail.com", false},
		{"send valid address", "email@gmail.com", true},
	}

	for _, tst := range testList {
		t.Run(tst.summary, func(t *testing.T) {
			tr := ValidateEmail(tst.input.(string))

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tEmail: %s, \n\tExpected: %s", tst.input, tst.expectedOutput)
			}
		})
	}
}

func TestCheckNewPassword(t *testing.T) {
	testlist := []struct {
		summary        string
		password       string
		checkpassword  string
		minimumlength  uint
		flag           uint8
		expectedOutput uint8
	}{
		{"test lowest flag", "1234AB", "1234AB", 6, CheckNewPasswordComplexityLowest, CheckNewPasswordResultOK},
		{"test check password", "1234AB", "1234A", 6, CheckNewPasswordComplexityLowest, CheckNewPasswordResultDivergent},
		{"Only Numbers with Default Flag", "1234", "1234", 4, CheckNewPasswordComplexityLowest, CheckNewPasswordResultOK},
		{"Only letters with Default Flag", "lala", "lala", 4, CheckNewPasswordComplexityLowest, CheckNewPasswordResultOK},
		{"testing minimum length", "1234", "1234", 2, CheckNewPasswordComplexityLowest, CheckNewPasswordResultOK},
		{"testing minimum length for password", "123", "123", 2, CheckNewPasswordComplexityLowest, CheckNewPasswordResultTooShort},
		{"test require letter success", "1234AB", "1234AB", 4, CheckNewPasswordComplexityRequireLetter, CheckNewPasswordResultOK},
		{"test require letter error", "1234", "1234", 4, CheckNewPasswordComplexityRequireLetter, CheckNewPasswordResultTooSimple},
		{"test require uppercase success", "1234Ab", "1234Ab", 4, CheckNewPasswordComplexityRequireUpperCase | CheckNewPasswordComplexityRequireLetter, CheckNewPasswordResultOK},
		{"test require uppercase error", "1234ab", "1234ab", 4, CheckNewPasswordComplexityRequireUpperCase | CheckNewPasswordComplexityRequireLetter, CheckNewPasswordResultTooSimple},
		{"test require number success", "abc1", "abc1", 4, CheckNewPasswordComplexityRequireNumber, CheckNewPasswordResultOK},
		{"test require number error", "abcd", "abcd", 4, CheckNewPasswordComplexityRequireNumber, CheckNewPasswordResultTooSimple},
		{"test require space success", "abc d", "abc d", 4, CheckNewPasswordComplexityRequireSpace, CheckNewPasswordResultOK},
		{"test require space error", "abcd", "abcd", 4, CheckNewPasswordComplexityRequireSpace, CheckNewPasswordResultTooSimple},
		{"test require symbol success", "abc#", "abc#", 4, CheckNewPasswordComplexityRequireSymbol, CheckNewPasswordResultOK},
		{"test require symbol error", "abcd", "abcd", 4, CheckNewPasswordComplexityRequireSymbol, CheckNewPasswordResultTooSimple},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := CheckNewPassword(tst.password, tst.checkpassword, tst.minimumlength, tst.flag)

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %d, \n\tGot: %d", tst.password, tst.expectedOutput, tr)
			}
		})
	}
}

func TestStringHash(t *testing.T) {
	testcases := []defaultTestStruct{
		{"Normal Test", "Handy", "E80649A6418B6C24FCCB199DAB7CB5BD6EC37593EA0285D52D717FCC7AEE5FB3"},
		{"string with number", "123456", "8D969EEF6ECAD3C29A3A629280E686CF0C3F5D5A86AFF3CA12020C923ADC6C92"},
		{"mashup", "Handy12345", "C82333DB3A6D91F98BE188C6C7B928DF4960B9EC3F3EB8CB50293368C673BE3D"},
		{"with symbols", "#handy_12Ax", "507512071AAEA24A94ECBB0F32EE74169FD59160EE9232819C504F39656E61F7"},
	}

	for _, tc := range testcases {
		t.Run(tc.summary, func(t *testing.T) {
			r := Sha256Hash(tc.input.(string))

			if r != strings.ToLower(tc.expectedOutput.(string)) {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %d, \n\tGot: %s", tc.input, tc.expectedOutput, r)
			}
		})
	}
}

func TestOnlyLetters(t *testing.T) {
	tcs := []defaultTestStruct{
		{"empty", "", ""},
		{"only letters", "haoplhu", "haoplhu"},
		{"letters and numbers", "hlo1234", "hlo"},
		{"symbols", "$#@", ""},
		{"numbers", "1234", ""},
		{"with space", "with space", "withspace"},
		{"A full phrase", "Hello Sr! Tell me, how are you?", "HelloSrTellmehowareyou"},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			r := OnlyLetters(tc.input.(string))

			if r != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %s, \n\tGot: %s", tc.input, tc.expectedOutput, r)
			}
		})
	}
}

func TestOnlyDigits(t *testing.T) {
	tcs := []defaultTestStruct{
		{"empty", "", ""},
		{"only letters", "haoplhu", ""},
		{"letters and numbers", "hlo1234", "1234"},
		{"symbols", "$#@", ""},
		{"numbers", "1234", "1234"},
		{"with space", "with space 10", "10"},
		{"A full phrase", "Hello Sr! I'm 24 Years Old!", "24"},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			r := OnlyDigits(tc.input.(string))

			if r != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %s, \n\tGot: %s", tc.input, tc.expectedOutput, r)
			}
		})
	}
}

func TestOnlyLettersAndNumbers(t *testing.T) {
	tcs := []defaultTestStruct{
		{"empty", "", ""},
		{"only letters", "haoplhu", "haoplhu"},
		{"letters and numbers", "hlo1234", "hlo1234"},
		{"symbols", "$#@", ""},
		{"numbers", "1234", "1234"},
		{"with space", "with space 10", "withspace10"},
		{"A full phrase", "Hello Sr! I'm 24 Years Old!", "HelloSrIm24YearsOld"},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			r := OnlyLettersAndNumbers(tc.input.(string))

			if r != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tInput: %s,\n\tExpected: %s, \n\tGot: %s", tc.input, tc.expectedOutput, r)
			}
		})
	}
}

func TestRandomInt(t *testing.T) {
	tcs := []struct {
		summary string
		min     int
		max     int
	}{
		{"normal test", 10, 20},
		{"big range", 10, 1000},
		{"negative", -10, 1000},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			r := RandomInt(tc.min, tc.max)

			if r < tc.min || r > tc.max {
				t.Errorf("Test has failed!\n\tMin: %d, \n\tMax: %d, \n\tGot: %d", tc.min, tc.max, r)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tcs := []struct {
		summary        string
		input          string
		limit          int
		trim           bool
		expectedOutput string
	}{
		{"normal Test", "The Go programming language is an open source project to make programmers more productive.", 25, false, "The Go programming langua"},
		{"normal Test with trim", "   The Go programming language is an open source project to make programmers more productive.", 45, true, "The Go programming language is an open sou"},
		{"zero", "The Go programming language is an open source project to make programmers more productive.", 0, true, ""},
		{"zero zero", "", 45, true, ""},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := Truncate(tc.input, tc.limit, tc.trim)
			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %s, \n\tGot: %s, \n\tInput: %s, \n\tlimit: %d, \n\ttrim: %t", tc.expectedOutput, tr, tc.input, tc.limit, tc.trim)
			}
		})
	}
}

func TestTransform(t *testing.T) {
	tcs := []struct {
		summary        string
		input          string
		max            int
		flags          TransformFlag
		expectedOutput string
	}{
		{"without flags", "The Go programming language is an open source project to make programmers more productive.", 20, TransformNone, "The Go programming l"},
		{"with trim", "   The Go programming language is an open source project to make programmers more productive.", 20, TransformTrim, "The Go programming l"},
		{"with lower case", "The Go programming language is an open source project to make programmers more productive.", 20, TransformLowerCase, "the go programming l"},
		{"with upper case", "The Go programming language is an open source project to make programmers more productive.", 20, TransformUpperCase, "THE GO PROGRAMMING L"},
		{"with Only Digits", "The Go is the 1º programming language is an open source project to make programmers more productive.", 20, TransformOnlyDigits, "1"},
		{"with Only Letters", "The Go is the 1º programming language is an open source project to make programmers more productive.", 20, TransformOnlyLetters, "TheGoistheºprogrammi"},
		{"with Only Letters and Numbers", "The Go is the 1º programming language is an open source project to make programmers more productive.", 20, TransformOnlyLettersAndDigits, "TheGoisthe1ºprogramm"},
		{"with Only Hash", "The Go is the 1º programming language is an open source project to make programmers more productive.", 20, TransformHash, "e68e17f094e7c05eb7c9"},
		{"with Only Hash and letters", "The Go is the 1º programming language is an open source project to make programmers more productive.", 20, TransformHash | TransformOnlyLetters, "a29f4806226150623d9d"},
		{"empty", "", 20, TransformHash | TransformOnlyLetters, ""},
		{"spacing", " ", 1, TransformOnlyLettersAndDigits | TransformOnlyLetters | TransformOnlyDigits | TransformOnlyLetters | TransformTrim | TransformLowerCase | TransformUpperCase, ""},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := Transform(tc.input, tc.max, tc.flags)

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %s, \n\tGot: %s, \n\tInput: %s, \n\tlimit: %d, \n\tflags: %d", tc.expectedOutput, tr, tc.input, tc.max, tc.flags)
			}
		})
	}
}

func TestHasOnlyNumbers(t *testing.T) {
	tcs := []struct {
		summary        string
		input          string
		expectedOutput bool
	}{
		{"normal test", "20", true},
		{"with string", "The Go programming language ", false},
		{"with part of a string", "20The Go programming language ", false},
		{"empty", "", false},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := HasOnlyNumbers(tc.input)

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %t, \n\tGot: %t, \n\tInput: %s", tc.expectedOutput, tr, tc.input)
			}
		})
	}
}

func TestHasOnlyLetters(t *testing.T) {
	tcs := []struct {
		summary        string
		input          string
		expectedOutput bool
	}{
		{"normal test", "TheGoprogramminglanguage", true},
		{"normal test with spaces", "The Go programming language", false},
		{"with numbers", "20", false},
		{"with part of a string", "20The Go programming language ", false},
		{"empty", "", false},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := HasOnlyLetters(tc.input)

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %t, \n\tGot: %t, \n\tInput: %s", tc.expectedOutput, tr, tc.input)
			}
		})
	}
}

func TestTrimLen(t *testing.T) {
	tcs := []struct {
		summary        string
		input          string
		expectedOutput int
	}{
		{"normal test", "TheGoprogramminglanguage", 24},
		{"normal test with spaces", "The Go programming language", 27},
		{"with numbers", "20", 2},
		{"with part of a string", "20The Go programming language ", 29},
		{"empty", "", 0},
		{"space", " ", 0},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := TrimLen(tc.input)

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %d, \n\tGot: %d, \n\tInput: %s", tc.expectedOutput, tr, tc.input)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	tcs := []defaultTestStruct{
		{"normal test", "Miguel", "leugiM"},
		{"2 chars", "Fe", "eF"},
		{"With spaces", "Lorem ipsum nibh sem laoreet taciti mattis neque ut, ornare cursus aenean inceptos suspendisse est hac hendrerit malesuada, luctus malesuada sit maecenas lorem arcu justo.", ".otsuj ucra merol saneceam tis adauselam sutcul ,adauselam tirerdneh cah tse essidnepsus sotpecni naenea susruc eranro ,tu euqen sittam iticat teeroal mes hbin muspi meroL"},
		{"String Number", "Ha1", "1aH"},
		{"empty", "", ""},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := Reverse(tc.input.(string))
			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %s, \n\tGot: %s, \n\tInput: %s", tc.expectedOutput, tr, tc.input)
			}
		})
	}
}

func TestTransformSerially(t *testing.T) {
	tcs := []struct {
		summary        string
		input          string
		max            int
		flags          []TransformFlag
		expectedOutput string
	}{
		{"without flags", "The Go programming language is an open source project to make programmers more productive.", 20, []TransformFlag{TransformNone}, "The Go programming l"},
		{"with trim and lowercase", "   The Go programming language is an open source project to make programmers more productive.", 20, []TransformFlag{TransformTrim, TransformLowerCase}, "the go programming l"},
		{"with lower case and only letters", "The Go programming language is an open source project to make programmers more productive.", 20, []TransformFlag{TransformLowerCase, TransformOnlyLetters}, "thegoprogramminglang"},
		{"with Only Hash and letters", "The Go is the 1º programming language is an open source project to make programmers more productive.", 20, []TransformFlag{TransformHash, TransformOnlyLetters}, "eefecebcfdbceccbbbcb"},
		{"without string", "", 20, []TransformFlag{TransformNone}, ""},
		{"Only letters and numbers", "The Go is the 1º! programming language is an open source project to make programmers more productive!", 20, []TransformFlag{TransformOnlyLettersAndDigits}, "TheGoisthe1ºprogramm"},
		{"Only numbers", "The Go is the 1º! programming language is an open source project to make programmers more productive!", 20, []TransformFlag{TransformOnlyDigits}, "1"},
		{"Go Upper!", "The Go is the 1º! programming language is an open source project to make programmers more productive!", 20, []TransformFlag{TransformUpperCase}, "THE GO IS THE 1º! PR"},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := TransformSerially(tc.input, tc.max, tc.flags...)

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %s, \n\tGot: %s, \n\tInput: %s, \n\tlimit: %d, \n\tflags: %d", tc.expectedOutput, tr, tc.input, tc.max, tc.flags)
			}
		})
	}
}

func TestStringReplaceAll(t *testing.T) {
	tcs := []struct {
		summary string
		input   string
		pairs   []string
		output  string
	}{
		{"normal test", "test string", []string{"t", "d"}, "desd sdring"},
		{"space test", "test string", []string{" ", "e"}, "testestring"},
		{"a lot of pairs test", "test string", []string{"t", "d", " ", "e"}, "desdesdring"},
		{"empty", "", []string{"t", "d", " ", "e"}, ""},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := ReplaceAll(tc.input, tc.pairs...)

			if tr != tc.output {
				t.Errorf("Error! Expected: %s, Got: %s, Input: %s, Pairs: %s", tc.output, tr, tc.input, tc.pairs)
			}
		})
	}
}

func TestDateTimeAsString(t *testing.T) {
	tcs := []struct {
		summary        string
		time           time.Time
		format         string
		expectedOutput string
	}{
		{"normal test", time.Date(2018, 10, 31, 01, 02, 02, 651387237, time.UTC), "yyyymmdd", "20181031"},
		{"test with format with points", time.Date(2018, 10, 31, 01, 02, 02, 651387237, time.UTC), "yyyy-mm-dd", "2018-10-31"},
		{"brazil test", time.Date(2018, 10, 31, 01, 02, 02, 651387237, time.UTC), "dd/mm/YYYY", "31/10/2018"},
	}

	for _, tc := range tcs {
		t.Run(tc.summary, func(t *testing.T) {
			tr := DateTimeAsString(tc.time, tc.format)

			if tr != tc.expectedOutput {
				t.Errorf("Test has failed!\n\tExpected: %s, \n\tGot: %s, \n\tDate: %s, \n\tFlag: %s", tc.expectedOutput, tr, tc.time, tc.format)
			}
		})
	}
}
