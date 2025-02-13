package com.example.nacosdiscovery.t2;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/16 09:12
 */
class Solution {
    public int minimumSubarrayLength(int[] nums, int k) {
        int res = Integer.MAX_VALUE;
        for (int i = 0; i < nums.length; i++) {
            int num = 0;
            for (int j = i; j < nums.length; j++) {
                num |= nums[j];
                if(num >= k) res = Math.min(res, j - i + 1);
            }
        }
        return res == Integer.MAX_VALUE ? -1 : res;
    }

    public static void main(String[] args) {
        System.out.println(12 & 8);
        System.out.println(12 | 10);
        // System.out.println(new Solution().minimumSubarrayLength(new int[]{16, 1, 2, 20, 32}, 45));
    }
}