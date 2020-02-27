// delete Node from singly linked Node
function deleteNode(head, position) {
    //if position is head reassign head
    if (position === 0) {
        return head.next;
    }

    //count the index delete positio assign previous node to the next one
    let i = 1;
    let prev = head;
    let curr = head.next;
    while (i <= position) {
        //if we reached the node link prev and next
        if (i === position) {
            prev.next = curr.next;
            return head;
        }
        i++;
        prev = curr;
        curr = curr.next;
    }
}
