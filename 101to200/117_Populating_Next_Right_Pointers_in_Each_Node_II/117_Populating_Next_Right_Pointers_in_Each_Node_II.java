public class Solution {


    public void connect(TreeLinkNode root) {

        if (root == null) {
            return;
        }
        int start = 0;
        int end = 0;

        TreeLinkNode[] queue = new TreeLinkNode[10000];

        TreeLinkNode cur = root;
        queue[end++] = cur;
        queue[end++] = null;


        while (start != end) {
            cur = queue[start++];
            if (cur == null) {
                if (start == end) {
                    break;
                } else {
                    queue[end++] = null;
                    //start到end之间都是同层的师兄妹啊。
                    for (int i = start; i < end - 1; i++) {
                        queue[i].next = queue[i + 1];
                    }
                    continue;
                }

            }
            if (cur.left != null) {
                queue[end++] = cur.left;
            }
            if (cur.right != null) {
                queue[end++] = cur.right;
            }
        }

    }

    public class TreeLinkNode {
        int val;
        TreeLinkNode left, right, next;

        TreeLinkNode(int x) {
            val = x;
        }
    }
}