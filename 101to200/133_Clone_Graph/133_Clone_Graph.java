package com.lish.practice.leetcode.p134;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * Definition for undirected graph.
 * class UndirectedGraphNode {
 * int label;
 * List<UndirectedGraphNode> neighbors;
 * UndirectedGraphNode(int x) { label = x; neighbors = new ArrayList<UndirectedGraphNode>(); }
 * };
 */
public class Solution {


    public UndirectedGraphNode cloneGraph(UndirectedGraphNode node) {

        if (node == null) {
            return null;
        }

        Map<Integer, UndirectedGraphNode> map = new HashMap<>();
        //同样的只去一次。
        Map<Integer, Boolean> visited = new HashMap<>();

        return cloneGraphRecursive(node, map, visited);
    }


    public UndirectedGraphNode cloneGraphRecursive(UndirectedGraphNode node, Map<Integer, UndirectedGraphNode> map, Map<Integer, Boolean> visited) {

        if (node == null) {
            return null;
        }


        UndirectedGraphNode newNode;
        if (map.get(node.label) != null) {
            newNode = map.get(node.label);
        } else {
            newNode = new UndirectedGraphNode(node.label);
            map.put(node.label, newNode);
        }


        if (visited.get(node.label) == null || !visited.get(node.label)) {
            visited.put(node.label, true);

            for (UndirectedGraphNode n : node.neighbors) {
                UndirectedGraphNode tmp = cloneGraphRecursive(n, map, visited);
                newNode.neighbors.add(tmp);
            }

            return newNode;
        } else {
            return newNode;
        }
    }


    public static class UndirectedGraphNode {
        int label;
        List<UndirectedGraphNode> neighbors;

        UndirectedGraphNode(int x) {
            label = x;
            neighbors = new ArrayList<UndirectedGraphNode>();
        }
    }

    public static void main(String[] args) {


        UndirectedGraphNode node0 = new UndirectedGraphNode(0);
        UndirectedGraphNode node1 = new UndirectedGraphNode(1);
        UndirectedGraphNode node2 = new UndirectedGraphNode(2);
        UndirectedGraphNode node3 = new UndirectedGraphNode(3);
        UndirectedGraphNode node4 = new UndirectedGraphNode(4);
        UndirectedGraphNode node5 = new UndirectedGraphNode(5);

        node0.neighbors.add(node1);
        node0.neighbors.add(node5);
        node1.neighbors.add(node2);
        node1.neighbors.add(node5);
        node2.neighbors.add(node3);
        node3.neighbors.add(node4);
        node3.neighbors.add(node4);
        node4.neighbors.add(node5);
        node4.neighbors.add(node5);

        UndirectedGraphNode cloneGraph1 = new Solution().cloneGraph(node0);

        System.out.println(cloneGraph1);


        UndirectedGraphNode node = new UndirectedGraphNode(0);
        node.neighbors.add(node);
        node.neighbors.add(node);

        UndirectedGraphNode cloneGraph = new Solution().cloneGraph(node);

        System.out.println(cloneGraph);
    }
}