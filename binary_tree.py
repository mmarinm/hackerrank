class Node:
    def __init__(self, data):
        self.data = data
        self.left = None
        self.right = None


class BST:
    def __init__(self):
        self.root = None

    def add(self, data):
        node = self.root

        if not (node):
            self.root = Node(data)

        else:
            def addToTree(node):
                if (data < node.data):
                    if (node.left):
                        addToTree(node.left)
                    else:
                        node.left = Node(data)
                        return
                elif (data > node.data):
                    if (node.right):
                        addToTree(node.right)
                    else:                    
                        node.right = Node(data)
                        return
                else:
                    return None

            addToTree(self.root)

    def printTree(self):
        current_level = [self.root]
        tree_array = []
        while(current_level):
            next_level = list()
            for n in current_level:
                tree_array.append(n.data)
                if (n.left):
                    next_level.append(n.left)
                if (n.right):
                    next_level.append(n.right)
            current_level = next_level
        print(tree_array)
        return tree_array

    def sortedArrayToBST(self, array, start, end):
        if (start == end):
            return None

        mid = (start + end) >> 1
        root = Node(array[mid])
        root.left = self.sortedArrayToBST(array, start, mid)
        root.right = self.sortedArrayToBST(array, mid + 1, end)
        return root


arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16,
     18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31]
myBST = BST()
myBST.root = myBST.sortedArrayToBST(arr, 0, len(arr))
# myBST.add(4)
# myBST.add(2)
# myBST.add(2)
# myBST.add(6)
# myBST.add(1)
# myBST.add(3)
# myBST.add(5)
# myBST.add(7)
tree = myBST.printTree()

print("Is BST?")
if len(tree) > len(set(tree)):
    print('NO')
else:
    print('YES')

