package stringo

import "testing"

func TestCheckPersonName(t *testing.T) {
	type TestStructForCheckPersonName struct {
		summary        string
		name           string
		acceptEmpty    bool
		expectedOutput ChkPersonNameResult
	}

	testlist := []TestStructForCheckPersonName{
		{"Only two letters", "T S", false, ChkPersonNameTooSimple},
		{"only four letters", "AB CD", false, ChkPersonNameTooSimple},
		{"five letters with non-ascii runes", "ça vá", false, ChkPersonNameTooSimple},
		{"mixing letters and numbers", "W0RDS W1TH NUMB3RS", false, ChkPersonNamePolluted},
		{"Sending and accepting empty string", "", true, ChkPersonNameOK},
		{"Sending spaces-only string and accepting empty", "     ", true, ChkPersonNameOK},
		{"Sending but not accepting empty string", " ", false, ChkPersonNameTooShort},
		{"Sending spaces-only string and refusing empty", "     ", false, ChkPersonNameTooShort},
		{"Sending numbers, expecting false", " 5454 ", true, ChkPersonNamePolluted},
		{"OneWorded string", "ONEWORD", false, ChkPersonNameTooFewWords},
		{"Minimum acceptable", "AB CDE", false, ChkPersonNameOK},
		{"Non-ascii stuff", "ÑÔÑÀSÇÏÏ ÇÃO ÀË", false, ChkPersonNameOK},
		{"Words with symbols. Expecting true", "WORDS-WITH SYMBOLS'", false, ChkPersonNameOK},
		{"Words with symbols. Expecting false", "WORDS WITH SYMBOLS`", false, ChkPersonNamePolluted},
		{"less than two letters", "a", false, ChkPersonNameTooFewWords},
		{"Sending numbers, expecting false", "5454", false, ChkPersonNamePolluted},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			tr := ChkPersonName(tst.name, tst.acceptEmpty)

			if tr != tst.expectedOutput {
				t.Errorf("Test has failed!\n\tName: %s\n\tAcceptEmpty: %t, \n\tExpected: %d, \n\tGot: %d,", tst.name, tst.acceptEmpty, tst.expectedOutput, tr)
			}
		})
	}
}

//// StringSlicesAreEqual compares two string slices and returns true if they have the same elements, in same order
//func StringSlicesAreEqual(x, y []string) bool {
//	if ((x == nil) != (y == nil)) || (len(x) != len(y)) {
//		return false
//	}
//
//	for i := range y {
//		if x[i] != y[i] {
//			return false
//		}
//	}
//
//	return true
//}

func TestNameFirstAndLast(t *testing.T) {
	type TestNameFirstAndLastStruct struct {
		summary         string
		name            string
		transformFlags  TransformFlag
		expectedOutputS string
	}

	testlist := []TestNameFirstAndLastStruct{
		{"Only two letters", "x Y", TransformNone, `x Y`},
		{"one word name", "namë", TransformNone, `namë`},
		{"all non-ascii runes", "çá öáã àÿ", TransformNone, `çá àÿ`},
		{"all non-ascii runes to upper", "çá öáã àÿ", TransformUpperCase, `ÇÁ ÀŸ`},
		{"mixing letters and numbers and then filtering digits off", "W0RDS W1TH NUMB3RS", TransformRemoveDigits, `WRDS NUMBRS`},
		{"empty string", "", TransformNone, ``},
		{"only spaces", "     ", TransformNone, ``},
		{"with spaces and tabs", " FIRST NAME - MIDDLENAME 	LAST	 ", TransformNone, `FIRST LAST`},
		{"last name single rune", "NAME X", TransformNone, `NAME X`},
		{"only symbols", "5454#@$", TransformNone, `5454#@$`},
		{"single letter", "x", TransformNone, `x`},
		{"only spaces empty return", " 		 ", TransformNone, ``},
		{"regular name to upper", "name lastname", TransformUpperCase, `NAME LASTNAME`},
		{"regular name to title", "name LASTNAME", TransformTitleCase, `Name Lastname`},
		{"REGULAR Name to lOwEr", "name LASTNAME", TransformLowerCase, `name lastname`},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			s := NameFirstAndLast(tst.name, tst.transformFlags)

			if s != tst.expectedOutputS {
				t.Errorf(`[%s] Test has failed! Given name: "%s", Expected string: "%s", Got: "%s"`, tst.summary, tst.name, tst.expectedOutputS, s)
			}
		})
	}
}

func TestNameFirst(t *testing.T) {
	type TestNameFirstStruct struct {
		summary         string
		name            string
		transformFlags  TransformFlag
		expectedOutputS string
	}

	testlist := []TestNameFirstStruct{
		{"Only two letters", "x Y", TransformNone, `x`},
		{"one word name", "namë", TransformNone, `namë`},
		{"all non-ascii runes", "çá öáã àÿ", TransformNone, `çá`},
		{"all non-ascii runes to upper", "çá öáã àÿ", TransformUpperCase, `ÇÁ`},
		{"mixing letters and numbers and then filtering digits off", "W0RDS W1TH NUMB3RS", TransformRemoveDigits, `WRDS`},
		{"empty string", "", TransformNone, ``},
		{"only spaces", "     ", TransformNone, ``},
		{"with spaces and tabs", " FIRST NAME - MIDDLENAME 	LAST	 ", TransformNone, `FIRST`},
		{"last name single rune", "NAME X", TransformNone, `NAME`},
		{"only symbols", "5454#@$", TransformNone, `5454#@$`},
		{"single letter", "x", TransformNone, `x`},
		{"only spaces empty return", " 		 ", TransformNone, ``},
		{"regular name to upper", "name lastname", TransformUpperCase, `NAME`},
		{"regular name to title", "name LASTNAME", TransformTitleCase, `Name`},
		{"REGULAR Name to lOwEr", "name LASTNAME", TransformLowerCase, `name`},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			s := NameFirst(tst.name, tst.transformFlags)

			if s != tst.expectedOutputS {
				t.Errorf(`[%s] Test has failed! Given name: "%s", Expected string: "%s", Got: "%s"`, tst.summary, tst.name, tst.expectedOutputS, s)
			}
		})
	}
}

func TestNameInitials(t *testing.T) {
	type tStruct struct {
		summary        string
		name           string
		transformFlags TransformFlag
		expectedOutput string
	}

	testlist := []tStruct{
		{`simplest 2 words name`, `miguel pragier`, TransformNone, `m p`},
		{`3 words name separated`, `ivan alexandrovitch kleshtakov`, TransformNone, `i a k`},
		{`3 words with unicode`, `Ívän Âlexandrovitch Çzelyatchenko`, TransformNone, `Í Â Ç`},
		{`3 words with unicode title-case`, `ívän âlexandrovitch çzelyatchenko`, TransformTitleCase, `Í Â Ç`},
		{`empty string`, ``, TransformNone, ``},
		{`dot`, `.`, TransformNone, `.`},
		{`spaces and tabs`, "  \t\t \n", TransformNone, ``},
		{`name with tabs`, "richard\t\tstallmann", TransformNone, `r s`},
		{`noble name with 1`, `dom pedro 1`, TransformNone, `d p 1`},
		{`noble name with I uppercase`, `dom pedro I`, TransformUpperCase, `D P I`},
		{`3 letters`, `x y z`, TransformNone, `x y z`},
		{`one word`, `asingleword`, TransformNone, `a`},
		{`comma separators`, `name,with,comma,separators`, TransformNone, `n`},
	}

	for _, tst := range testlist {
		t.Run(tst.summary, func(t *testing.T) {
			s := Initials(tst.name)

			if s != tst.expectedOutput {
				t.Errorf(`[%s] Test has failed! Given name: "%s", Expected string: "%s", Got: "%s"`, tst.summary, tst.name, tst.expectedOutput, s)
			}
		})
	}
}
