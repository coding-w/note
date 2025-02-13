package com.example.nacosdiscovery.t3;


class Solution {
    // https://leetcode.cn/problems/find-closest-number-to-zero/?envType=daily-question&envId=2025-01-20
    public int findClosestNumber(int[] nums) {
        int minLen = Integer.MAX_VALUE;
        int maxNum = Integer.MAX_VALUE;
        for (int num : nums) {
            int abs = Math.abs(num);
            if (abs < minLen) {
                minLen = abs;
                maxNum = num;
            }
            if (abs == minLen) {
                maxNum = Math.max(maxNum, num);
            }
        }
        return maxNum;
    }
}