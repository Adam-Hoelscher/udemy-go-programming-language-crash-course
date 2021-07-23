# Python3 code to implement Priority Queue
# using Singly Linked List

# Class to create new node which includes
# Node Data, and Node Priority
class PriorityQueueNode:

    def __init__(self, value, pr):

        self.data = value
        self.priority = pr
        self.next = None


# Implementation of Priority Queue
class PriorityQueue:

    def __init__(self):
        self.front = None

    # Method to check Priority Queue is Empty
    # or not if Empty then it will return True
    # Otherwise False
    def isEmpty(self):
        return self.front is None

    # Method to add items in Priority Queue
    # According to their priority value
    def push(self, value, priority):

        # Creating a new node
        newNode = PriorityQueueNode(
            value,
            priority)

        # Starting at the head of the list, go through until we run out of nodes
        # or locate a node with a lower priority
        after, before = self.front, None
        while after and after.priority <= newNode.priority:
            after, before = after.next, after

        # if nothing is before this node, it is the new front
        if before is None:
            self.front = newNode
        else:
            before.next = newNode

        # The next thing after the new node
        newNode.next = after

        # Returning True for successful execution
        return True

    # Method to remove high priority item
    # from the Priority Queue
    def pop(self):

        # Condition check for checking
        # Priority Queue is empty or not
        if self.isEmpty():
            val = None

        else:
            # Removing high priority node from
            # Priority Queue, and updating front
            # with next node
            val = self.front.data
            self.front = self.front.next

        return val

    # Method to return high priority node
    # value Not removing it
    def peek(self):

        # Condition check for checking Priority
        # Queue is empty or not
        if self.isEmpty() == True:
            return
        else:
            return self.front.data

    # Method to Traverse through Priority
    # Queue
    def traverse(self):

        # Condition check for checking Priority
        # Queue is empty or not
        if self.isEmpty():
            print("Queue is Empty!")
        else:
            temp = self.front
            while temp:
                print(temp.data, end = " ")
                temp = temp.next
            print()

# Driver code
if __name__ == "__main__":

    # Creating an instance of Priority
    # Queue, and adding values
    # 7 -> 4 -> 5 -> 6
    pq = PriorityQueue()
    pq.push(4, 1)
    pq.push(5, 2)
    pq.push("a", 3)
    pq.push("c", 3)
    pq.push("b", 3)
    pq.push(7, 0)

    # Traversing through Priority Queue
    pq.traverse()

    # Removing highest Priority item
    # for priority queue
    while not pq.isEmpty():
        print(pq.pop())

# This code is contributed by himanshu kanojiya
