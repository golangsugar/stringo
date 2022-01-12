package stringo

import (
	"testing"
)

func TestCheckStr(t *testing.T) {
	testlist := []struct {
		summary        string
		input          string
		minLen         uint
		maxLen         uint
		flags          ChkRule
		expectedOutput ChkResult
	}{
		{"perfect string", "ЀЁЂЃЄЅІЇЈЉЊЋЌЍЎЏ АБВГДЕЖЗИЙКЛМНОП РСТУФХЦЧШЩЪЫ ЬЭЮЯ абвгдежзийкл мнопрстуфхцчшщ ъыьэюяѐёђѓєѕіїј љњћќѝўџѠѡѢѣѤѥѦѧѨѩѪѬѭѮѯ", 0, 118, 0, ChkOk},
		{"allowed empty string", "", 0, 1000, ChkAllowEmpty, ChkOk},
		{"fails on empty string", "", 0, 100, 0, ChkEmptyDenied},
		{"fails on min length", "four", 5, 100, 0, ChkTooShort},
		{"fails on max length", "four", 4, 3, 0, ChkTooLong},
		{"denies space", "two words", 0, 100, ChkDenySpaces, ChkSpaceDenied},
		{"denies tab", "two	words", 0, 100, ChkDenySpaces, ChkSpaceDenied},
		{`denies \n`, "row\nrow", 0, 100, ChkDenySpaces, ChkSpaceDenied},
		{`denies \r`, "row\rrow", 0, 100, ChkDenySpaces, ChkSpaceDenied},
		{"denies numbers", "this is a number: 9", 0, 100, ChkDenyNumbers, ChkNumbersDenied},
		{"denies numbers UNICODE", " 九 に 三 Ⅷ'", 0, 100, ChkDenyNumbers, ChkNumbersDenied},
		{"denies letters", "o123456789", 0, 100, ChkDenyLetters, ChkLettersDenied},
		{"denies symbols", "a symbol %", 0, 100, ChkDenySymbols, ChkSymbolsDenied},
		{"denies more than one word", "two words", 0, 100, ChkDenyMoreThanOneWord, ChkMoreThanOneWordDenied},
		{"denies words separated by new line", "one\ntwo", 0, 100, ChkDenyMoreThanOneWord, ChkMoreThanOneWordDenied},
		{"denies uppercase", "upper Case", 0, 100, ChkDenyUpperCase, ChkUpperCaseDenied},
		{"denies lowercase", "LOWER cASE", 0, 100, ChkDenyLowercase, ChkLowercaseDenied},
		{"denies unicode", "TAB	ÇÂÖÉд", 0, 100, ChkDenyUnicode, ChkUnicodeDenied},
		{"missed numbers", "NO NUMBERS", 0, 100, ChkRequireNumbers, ChkNumbersNotFound},
		{"missed letters", " 87 %323232	", 0, 100, ChkRequireLetters, ChkLettersNotFound},
		{"missed symbols", "NO SYMBOLS 123", 0, 100, ChkRequireSymbols, ChkSymbolsNotFound},
		{"missed more than one word", "FHFJKDHFSDJKH012308312-0=-0=-00", 0, 100, ChkRequireMoreThanOneWord, ChkMoreThanOneWordNotFound},
		{"missed uppercase", "all lowercase 456", 0, 100, ChkRequireUpperCase, ChkUpperCaseNotFound},
		{"missed lowercase", "ALL UPPERCASE 789", 0, 100, ChkRequireLowercase, ChkLowercaseNotFound},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := CheckStr(tst.input, tst.minLen, tst.maxLen, tst.flags)

			if tr != tst.expectedOutput {
				t.Errorf("Failed with input %s, want %d and got %d instead", tst.input, tst.expectedOutput, tr)
			}
		})
	}
}

func TestStrContainsEmail(t *testing.T) {
	testList := []defaultTestStruct{
		{"empty string", "", false},
		{"no valid email", "email-gmail.com", false},
		{"the string is an email address", "email@gmail.com", true},
		{"valid email within", "dasdsdsdsda-*9email@gmail.comdsdsds.88", true},
	}

	for _, tst := range testList {
		t.Run(tst.summary, func(t *testing.T) {
			if StrContainsEmail(tst.input.(string)) != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tEmail: %s, \n\tExpected: %v", tst.input, tst.expectedOutput)
			}
		})
	}
}
