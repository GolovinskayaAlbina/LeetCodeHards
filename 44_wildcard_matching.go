////////////////////////////////////////////////////////////////////////////////////////
//
// https://leetcode.com/problems/wildcard-matching/
//
// Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*' where:
//
// '?' Matches any single character.
// '*' Matches any sequence of characters (including the empty sequence).
// The matching should cover the entire input string (not partial).
//
// Constraints:
//
// 0 <= s.length, p.length <= 2000
// s contains only lowercase English letters.
// p contains only lowercase English letters, '?' or '*'.
//
//////////////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

func isMatch(s string, p string) bool {
	res, _ := checkIfMatch(s, p, 0, 0)
	return res
}

func checkIfMatch(s, p string, sPos, pPos int) (bool, bool) {
	for pPos < len(p) {
		if p[pPos] == '*' {
			if pPos == len(p)-1 {
				//last * in pattern and string fin
				return true, false
			}
			offset := sPos
			for offset <= len(s) {
				res, fin := checkIfMatch(s, p, offset, pPos+1)
				if fin || res {
					return res, fin
				}
				offset++
			}
			return false, true
		} else {
			if sPos >= len(s) {
				// string was ended
				return false, false
			}
			fmt.Printf("%d[%c]:%d[%c]", sPos, s[sPos], pPos, p[pPos])
			if p[pPos] != '?' && s[sPos] != p[pPos] {
				fmt.Printf(": false\n")
				return false, false
			}
			fmt.Printf(": true\n")
			sPos += 1
		}
		pPos += 1
	}
	if sPos < len(s) {
		return false, false
	}
	return true, true
}
