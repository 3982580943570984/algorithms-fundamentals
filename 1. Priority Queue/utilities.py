import dearpygui.dearpygui as dpg
from priority_queue import *

priority_queue = PriorityQueue()

priority_queue.insert(45)
priority_queue.insert(20)
priority_queue.insert(14)
priority_queue.insert(12)
priority_queue.insert(31)
priority_queue.insert(7)
priority_queue.insert(11)
priority_queue.insert(13)
priority_queue.insert(7)

selection = []


def create_queue(sender, data):
    selection.clear()
    priority_queue.clear()

    update_images()
    reload_images()


def clear_queue(sender, data):
    selection.clear()
    priority_queue.clear()

    update_images()
    reload_images()


def extract_max(sender, data):
    if priority_queue.heap:
        selection.append(priority_queue.extract_max())

    update_images()
    reload_images()


def insert(sender, data, user_data):
    if (len(priority_queue.heap) < 15):
        value = dpg.get_value("Insert Widget")
        priority_queue.insert(value=value)

        update_images()
        reload_images()
    else:
        dpg.configure_item("modal_id", show=True)


def change_prority(sender, data):
    index = dpg.get_value("Index Change")
    if (index < len(priority_queue.heap)):
        value = dpg.get_value("Index Value")
        priority_queue.change_priority(index, value)

        update_images()
        reload_images()
    else:
        dpg.configure_item("modal_id2", show=True)


def graph_array():
    dot = Digraph(name="Selection")
    dot.graph_attr = {"bgcolor": "#252525"}
    dot.edge_attr = {"color": "white"}
    dot.node_attr = {"shape": "record", "color": "white", "fontcolor": "white"}
    dot.attr("node", shape="record")
    result = str()
    for i in range(0, len(selection)):
        result += str(selection[i]) + "|"
    for i in range(0, 15 - len(selection)):
        result += "|"
    dot.node(name="node", label=result[:-1])
    dot.render(directory="graph", format="png")


def load_image(image_path):
    width, height, channels, data = dpg.load_image(image_path)

    with dpg.texture_registry() as reg_id:
        texture_id = dpg.add_static_texture(width, height, data, parent=reg_id)

    return texture_id


def update_images():
    priority_queue.graph_array()
    priority_queue.graph_tree()
    graph_array()


def reload_images():
    dpg.delete_item("Array")
    texture_array = load_image(image_path="./graph/Array.gv.png")
    dpg.add_image(tag="Array", texture_tag=texture_array, before=1)

    dpg.delete_item("Tree")
    texture_tree = load_image(image_path="./graph/Tree.gv.png")
    dpg.add_image(tag="Tree", texture_tag=texture_tree, before=2)

    dpg.delete_item("Selection")
    texture_selection = load_image(image_path="./graph/Selection.gv.png")
    dpg.add_image(tag="Selection", texture_tag=texture_selection, before=3)


def exit(sender, data):
    dpg.stop_dearpygui()
