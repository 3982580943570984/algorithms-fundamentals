from graphviz import *

# Создание очереди с отображением в виде дерева и массива
# Очистка очереди (очистка массива и дерева в приложении)
# Извлечение наибольшего элемента из очереди
# Добавление нового элемента в очередь (обновление массива и дерева в приложении)
# Изменение приоритета элемента (обновление массива и дерева в приложении)

# Максимальный размер очереди - 15


class PriorityQueue:
    """Priority Queue built on Binary Heap"""

    heap = []

    def __init__(self) -> None:
        self.heap = []

    def parent(self, index: int):
        return (index - 1) // 2

    def left_child(self, index: int):
        return index * 2 + 1

    def right_child(self, index: int):
        return index * 2 + 2

    def shift_up(self, index: int):
        # While not reached root of the heap and parent node value less than child node value
        while (index > 0 and self.heap[self.parent(index)] < self.heap[index]):
            # Swap parent node value and child node value
            self.heap[self.parent(index)], self.heap[index] = self.heap[
                index], self.heap[self.parent(index)]
            # Change index value to point on parent node
            index = self.parent(index)

    def shift_down(self, index: int):
        keep_index = index
        left_child, right_child = self.left_child(
            keep_index), self.right_child(keep_index)

        # If left child node has higher value than current node
        if (left_child < len(self.heap)
                and self.heap[left_child] > self.heap[keep_index]):
            keep_index = left_child

        # If right child node has higher value than current node
        if (right_child < len(self.heap)
                and self.heap[right_child] > self.heap[keep_index]):
            keep_index = right_child

        # Swap node values if current node has smaller value than child node
        if (keep_index != index):
            self.heap[keep_index], self.heap[index] = self.heap[
                index], self.heap[keep_index]
            self.shift_down(keep_index)

    def insert(self, value: int):
        # Insert value as a leaf in the heap
        self.heap.append(value)
        # Move new value as high as possible
        self.shift_up(len(self.heap) - 1)

    def extract_max(self) -> int:
        if (len(self.heap) == 0):
            return

        # Keep root value in a return variable
        result = self.heap[0]
        # Change root with a leaf
        self.heap[0] = self.heap[-1]
        # Remove leaf after transition
        self.heap.pop()
        # Move new root as low as possible
        self.shift_down(0)
        return result

    def change_priority(self, index: int, priority: int):
        if (index >= 0 and index < len(self.heap)):
            old_priority = self.heap[index]
            self.heap[index] = priority
            # Use appropriate function via priority change
            self.shift_up(
                index) if old_priority < priority else self.shift_down(index)

    def graph_tree(self):
        dot = Digraph(name="Tree")
        dot.graph_attr = {"bgcolor": "#252525"}
        dot.edge_attr = {"color": "white"}
        dot.node_attr = {"color": "white", "fontcolor": "white"}
        nodes = [i for i in range(0, len(self.heap))]
        for pos in nodes:
            dot.node(str(pos), str(self.heap[pos]))
            if self.left_child(pos) < len(self.heap):
                dot.edge(str(pos), str(self.left_child(pos)))
            if self.right_child(pos) < len(self.heap):
                dot.edge(str(pos), str(self.right_child(pos)))
        dot.render(directory="graph", format="png")

    def graph_array(self):
        dot = Digraph(name="Array")
        dot.graph_attr = {"bgcolor": "#252525"}
        dot.edge_attr = {"color": "white"}
        dot.node_attr = {
            "shape": "record",
            "color": "white",
            "fontcolor": "white"
        }
        result = str()
        for i in range(0, len(self.heap)):
            result += str(self.heap[i]) + "|"
        for i in range(0, 15 - len(self.heap)):
            result += "|"
        dot.node(name="node", label=result[:-1])
        dot.render(directory="graph", format="png")

    def clear(self):
        self.heap.clear()


if __name__ == "__main__":
    pq = PriorityQueue()

    pq.insert(45)
    pq.insert(20)
    pq.insert(14)
    pq.insert(12)
    pq.insert(31)
    pq.insert(7)
    pq.insert(11)
    pq.insert(13)
    pq.insert(7)

    pq.graph_tree()
    pq.graph_array()
