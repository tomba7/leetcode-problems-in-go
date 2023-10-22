/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     struct ListNode *next;
 * };
 */
bool kNodesExist(struct ListNode *node, int k) {
    for (; node && k > 0; node = node->next, k--) {}
    return k == 0? true: false;
}

struct ListNode* reverseKGroup(struct ListNode* head, int k) {
    struct ListNode dummy = {.next = head}, *grpTail, *prevGrpTail = &dummy, *prev, *curr = head, *next;
    while (curr) {
        grpTail = curr;
        if (!kNodesExist(curr, k)) break;
        prev = NULL;
        for (int i = k; curr && i > 0; i--) {
            next = curr->next;
            curr->next = prev;
            prev = curr;
            curr = next;
        }
        prevGrpTail->next = prev;
        grpTail->next = curr;
        prevGrpTail = grpTail;
    }
    return dummy.next;
}
