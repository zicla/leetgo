package com.lish.practice.leetcode.p151;

import java.util.ArrayList;
import java.util.List;

public class Solution {

    public String reverseWords(String s) {
        List<String> list = new ArrayList<>();

        int N = s.length();
        for (int i = 0; i < N; i++) {
            if (s.charAt(i) == ' ') {
                continue;
            }
            for (int j = i; j < N; j++) {
                //形成一个单词了
                if (j == N - 1) {
                    if (s.charAt(j) == ' ') {
                        list.add(s.substring(i, j));
                    } else {
                        list.add(s.substring(i, N));
                    }
                    i = j;
                    break;
                } else {
                    if (s.charAt(j) == ' ') {
                        list.add(s.substring(i, j));
                        i = j;
                        break;
                    }
                }
            }
        }

        StringBuilder a = new StringBuilder();
        for (int i = list.size() - 1; i >= 0; i--) {
            if (i == list.size() - 1) {
                a = new StringBuilder(list.get(i));
            } else {
                a.append(" ").append(list.get(i));
            }
        }
        return a.toString();
    }

    public static void main(String[] args) {
        System.out.println(new Solution().reverseWords("the sky is blue"));
    }

}