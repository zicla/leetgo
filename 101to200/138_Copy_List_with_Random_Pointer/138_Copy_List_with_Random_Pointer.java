package com.lish.practice.leetcode.p138;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Definition for singly-linked list with a random pointer.
 * class RandomListNode {
 * int label;
 * RandomListNode next, random;
 * RandomListNode(int x) { this.label = x; }
 * };
 */
public class Solution {


    public RandomListNode copyRandomList(RandomListNode head) {

        if (head == null) {
            return null;
        }

        //old->new
        Map<RandomListNode, RandomListNode> map = new HashMap<>();

        //new all nodes.
        RandomListNode cur = head;
        while (cur != null) {
            map.put(cur, new RandomListNode(cur.label));
            cur = cur.next;
        }

        //assign next and random.
        cur = head;
        while (cur != null) {
            if (cur.next != null) {

                map.get(cur).next = map.get(cur.next);
            }
            if (cur.random != null) {
                map.get(cur).random = map.get(cur.random);
            }
            cur = cur.next;
        }

        return map.get(head);

    }

    /**
     * Definition for singly-linked list with a random pointer.
     */
    public static class RandomListNode {
        int label;
        RandomListNode next, random;

        RandomListNode(int x) {
            this.label = x;
        }
    }

    public static void main(String[] args) {

        RandomListNode node0 = new RandomListNode(0);
        RandomListNode node1 = new RandomListNode(1);
        RandomListNode node2 = new RandomListNode(2);
        RandomListNode node3 = new RandomListNode(3);
        RandomListNode node4 = new RandomListNode(4);
        RandomListNode node5 = new RandomListNode(5);

        node0.next = node1;
        node1.next = node2;
        node2.next = node3;
        node3.next = node4;
        node4.next = node5;
        node0.random = node5;

        RandomListNode randomListNode = new Solution().copyRandomList(node0);
        System.out.println(randomListNode);
    }
}