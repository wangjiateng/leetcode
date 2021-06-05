fn main() {
    let head = Some(Box::new(ListNode {
        val: 1,
        next: Some(Box::new(ListNode {
            val: 2,
            next: Some(Box::new(ListNode { val: 2, next: Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode { val: 1, next: Some(Box::new(ListNode {
                    val: 2,
                    next: Some(Box::new(ListNode { val: 1, next: Some(Box::new(ListNode {
                        val: 2,
                        next: Some(Box::new(ListNode { val: 1, next: None })),
                    })) })),
                })) })),
            })) })),
        })),
    }));
    let mut head = remove_elements(head, 1);
    while let Some(node) = head.as_ref() {
        println!("val: {}", node.val);
        head = node.next.to_owned();
    }
}

/// 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
///
#[doc = "https://leetcode-cn.com/problems/remove-linked-list-elements/"]
pub fn remove_elements(head: Option<Box<ListNode>>, val: i32) -> Option<Box<ListNode>> {
    let mut head = Some(Box::new(ListNode { val: val + 1, next: head }));
	let mut node = &mut head;

    'OUTER:
    while let Some(cur) = node {
        while let Some(next) = &mut cur.next {
            if next.val == val {
                cur.next = next.next.take();
            } else {
                node = &mut cur.next;
                continue 'OUTER;
            }
        }
        break;
    }
    
    head.unwrap().next
}

/// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}