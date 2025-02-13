package com.example.nacosdiscovery.t1;

import java.util.PriorityQueue;

/**
 * @author wangxiang
 * @description
 * @create 2025/1/14 09:46
 */
public class Solution {
    public int minOperations(int[] nums, int k) {
        PriorityQueue<Long> pq = new PriorityQueue<>();
        for (int num : nums) {
            pq.add((long) num);
        }
        int res = 0;
        while (pq.size() > 2) {
            long num1 = pq.poll();
            long num2 = pq.poll();
            if (num1 >= k && num2 >= k) {
                break;
            }
            pq.add(num1 * 2 + num2);
            res++;
        }
        return res;
    }

    public static void main(String[] args) {
        System.out.println("new Solution().minOperations(new int[]{1, 2, 3, 4, 5}, 2) = " + new Solution().minOperations(new int[]{1000000000, 999999999, 1000000000, 999999999, 1000000000, 999999999}, 1000000000));
    }
}
