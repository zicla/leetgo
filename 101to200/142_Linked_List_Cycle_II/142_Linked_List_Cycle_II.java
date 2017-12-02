package com.lish.practice.leetcode.p142;

public class Solution {

    //具体分析参见：https://www.cnblogs.com/hiddenfox/p/3408931.html
    public ListNode detectCycle(ListNode head) {

        ListNode slow = head;
        ListNode fast = head;
        ListNode meet = null;
        int step = 0;
        while (slow != null && fast != null && fast.next != null) {
            slow = slow.next;
            fast = fast.next.next;
            step++;
            if (slow == fast) {
                meet = slow;
                break;
            }
        }

        if (meet != null) {
            slow = head;
            while (slow != meet) {
                slow = slow.next;
                meet = meet.next;
            }
            return slow;
        }


        return null;

    }

    public static class ListNode {
        int val;
        ListNode next;

        ListNode(int x) {
            val = x;
            next = null;
        }
    }

}