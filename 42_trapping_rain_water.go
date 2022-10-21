////////////////////////////////////////////////////////////////////////////////////////
//
// https://leetcode.com/problems/trapping-rain-water/
//
// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.
//
// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1].
//              In this case, 6 units of rain water (blue section) are being trapped.
//
//////////////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

func trap(height []int) int {
	res, count, row, peaks := 0, 1, 1, 0
	for count > 0 || peaks > 1 {
		count, peaks = getRowTrap(row, 0, height)
		res += count
		fmt.Printf("[%d]: count:%d, res:%d, peaks: %d\n", row, count, res, peaks)
		row += 1
	}
	return res
}

func getRowTrap(row, lWall int, height []int) (count int, peaks int) {
	for lWall < len(height) {
		if lWall == 0 {
			//move to first left `wall`
			if height[lWall] < row {
				for lWall < len(height) && height[lWall] < row {
					lWall++
				}
				if lWall < len(height) {
					peaks++
				}
			} else {
				peaks += 1
				for lWall < len(height)-1 && height[lWall+1] >= row {
					peaks++
					lWall++
				}
				if lWall >= len(height)-1 {
					return count, peaks
				}
			}
		}
		//move to right `wall`
		rWall := lWall + 1
		for rWall < len(height) && height[rWall] < row {
			rWall++
		}
		if rWall >= len(height) {
			return count, peaks
		}
		peaks += 1
		count += rWall - lWall - 1
		lWall = rWall
	}
	return count, peaks
}
