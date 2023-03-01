import dearpygui.dearpygui as dpg
from utilities import *
from constants import *

update_images()

dpg.create_context()
dpg.create_viewport(title="Очередь с приоритетом",
                    width=SCREEN_WIDTH,
                    height=SCREEN_HEIGHT,
                    min_width=SCREEN_WIDTH,
                    min_height=SCREEN_HEIGHT,
                    max_width=SCREEN_WIDTH,
                    max_height=SCREEN_HEIGHT,
                    resizable=True,
                    x_pos=DISPLAY_CENTER_X,
                    y_pos=DISPLAY_CENTER_Y,
                    small_icon='./icons/llvm.ico',
                    large_icon='./icons/llvm.ico')

with dpg.window(tag="Primary"):

    with dpg.font_registry():
        with dpg.font(file="C:/Windows/Fonts/JetBrains Mono Thin Nerd Font Complete Mono.ttf", size=15, parent=dpg.last_item()) as default_font:
            dpg.add_font_range_hint(hint=dpg.mvFontRangeHint_Cyrillic)

    dpg.bind_font(default_font)

    with dpg.group(horizontal=True):
        dpg.add_button(label="Создать очередь",
                       width=BUTTON_WIDTH,
                       callback=create_queue)
        dpg.add_button(label="Очистить очередь",
                       width=BUTTON_WIDTH,
                       callback=clear_queue)
        dpg.add_button(label="Извлечь наиб. элемент",
                       width=BUTTON_WIDTH,
                       callback=extract_max)

    with dpg.group(horizontal=True):
        dpg.add_button(label="Вставить значение",
                       width=BUTTON_WIDTH, callback=insert)
        insert_widget = dpg.add_input_int(tag="Insert Widget", width=242)

    with dpg.group(horizontal=True, tag=0):
        dpg.add_button(label="Изменить приоритет",
                       width=BUTTON_WIDTH, callback=change_prority)
        dpg.add_text(default_value="Индекс")
        dpg.add_input_int(width=192, min_value=0, max_value=15,
                          min_clamped=True, max_clamped=True, tag="Index Change")
        dpg.add_text(default_value="Значение")
        dpg.add_input_int(width=178, tag="Index Value")

    dpg.add_spacer(height=SPACER_HEIGHT)
    dpg.add_text(default_value="Представление в виде массива")
    texture_array = load_image(image_path="./graph/Array.gv.png")
    dpg.add_image(tag="Array", texture_tag=texture_array)

    dpg.add_spacer(height=SPACER_HEIGHT, tag=1)
    dpg.add_text(default_value="Представление в виде дерева")
    texture_tree = load_image(image_path="./graph/Tree.gv.png")
    dpg.add_image(tag="Tree", texture_tag=texture_tree)

    dpg.add_spacer(height=SPACER_HEIGHT, tag=2)
    dpg.add_text(default_value="Результат выборки")
    texture_selection = load_image(image_path="./graph/Selection.gv.png")
    dpg.add_image(tag="Selection", texture_tag=texture_selection)

    dpg.add_spacer(height=SPACER_HEIGHT, tag=3)
    dpg.add_button(label="Выход", callback=exit)

with dpg.window(modal=True, show=False, tag="modal_id", no_title_bar=True, width=300, height=160, pos=[SCREEN_WIDTH / 2 - 150, SCREEN_HEIGHT / 2 - 80]):
    dpg.add_text("Невозможно добавить значение")
    dpg.add_text("Очередь уже заполнена")
    with dpg.group(horizontal=True):
        dpg.add_button(label="OK", width=80, callback=lambda: dpg.configure_item(
            "modal_id", show=False), pos=[120, 130])

with dpg.window(modal=True, show=False, tag="modal_id2", no_title_bar=True, width=300, height=160, pos=[SCREEN_WIDTH / 2 - 150, SCREEN_HEIGHT / 2 - 80]):
    dpg.add_text("Невозможно изменить приоритет")
    dpg.add_text("Указан неверный индекс")
    with dpg.group(horizontal=True):
        dpg.add_button(label="OK", width=80, callback=lambda: dpg.configure_item(
            "modal_id2", show=False), pos=[120, 130])

dpg.setup_dearpygui()
dpg.show_viewport()
dpg.set_primary_window("Primary", True)
dpg.start_dearpygui()
dpg.destroy_context()
