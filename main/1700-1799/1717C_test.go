// Code generated by copypasta/template/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// https://codeforces.com/problemset/problem/1717/C
// https://codeforces.com/problemset/status/1717/problem/C
func Test_cf1717C(t *testing.T) {
	testCases := [][2]string{
		{
			`5
3
1 2 5
1 2 5
2
2 2
1 3
4
3 4 1 2
6 4 2 5
3
2 4 1
4 5 3
5
1 2 3 4 5
6 5 6 7 6`,
			`YES
NO
NO
NO
YES`,
		},
		{
			`1
5
3 4 1 4 1
4 4 2 4 2`,
			`YES`,
		},
	}
	testutil.AssertEqualStringCase(t, testCases, 0, cf1717C)
}
