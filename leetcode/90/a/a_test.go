// Code generated by copypasta/template/leetcode/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`"ab"`, `"ba"`, 
			`true`,
		},
		{
			`"ab"`, `"ab"`, 
			`false`,
		},
		{
			`"aa"`, `"aa"`, 
			`true`,
		},
		{
			`"aaaaaaabc"`, `"aaaaaaacb"`, 
			`true`,
		},
		{
			`""`, `"aa"`, 
			`false`,
		},
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, buddyStrings, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://leetcode-cn.com/contest/weekly-contest-90/problems/buddy-strings/
