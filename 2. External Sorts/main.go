package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var array = GenerateRandomNumbers(10000)

var data = [][]string{
	{"", "Сравнения", "Присвоения", "Время", "Отсортировано?"},
	{"Простое 2Ф", "", "", "", ""},
	{"Простое 1Ф", "", "", "", ""},
	{"Естественное 2Ф", "", "", "", ""},
	{"Естественное 1Ф", "", "", "", ""},
	{"Поглощение", "", "", "", ""},
}

func main() {
	myApp := app.New()

	myWindow := myApp.NewWindow("Алгоритмы внешней сортировки")
	myWindow.CenterOnScreen()

	checkGroup := widget.NewCheckGroup([]string{"Простое 2Ф", "Простое 1Ф", "Естественное 2Ф", "Естественное 1Ф", "Поглощение"}, nil)
	checkGroupContainer := container.New(layout.NewGridLayoutWithRows(1), checkGroup)

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			w := widget.NewLabel("\t\t\t")
			return w
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	splitContainer := container.NewHSplit(checkGroupContainer, table)
	splitContainer.SetOffset(0.0)
	topContainer := container.NewVScroll(splitContainer)
	topContainer.SetMinSize(fyne.Size{Width: 870, Height: 235})

	sizeBinding := binding.NewString()
	percentBinding := binding.NewString()

	buttonSort := widget.NewButton("Сортировать", func() {
		n, _ := sizeBinding.Get()
		size, _ := strconv.ParseInt(n, 10, 64)

		block, _ := percentBinding.Get()
		blockSize, _ := strconv.ParseInt(block, 10, 64)
		if blockSize == 0 {
			blockSize = 1
		}

		array = GenerateRandomNumbers(int(size))

		selectedSortings := checkGroup.Selected
		if len(selectedSortings) != 0 {
			array_copy := make([]int, len(array))
			copy(array_copy, array)

			for _, v := range selectedSortings {
				switch v {
				case "Простое 2Ф":
					comparisons, permutations, time, isSorted := TwoPhaseMergeSort(array_copy)
					data[1][1], data[1][2] = strconv.FormatInt(int64(comparisons), 10), strconv.FormatInt(int64(permutations), 10)
					data[1][3], data[1][4] = strconv.FormatInt(time, 10), strconv.FormatBool(isSorted)
				case "Простое 1Ф":
					comparisons, permutations, time, isSorted := OnePhaseMergeSort(array_copy)
					data[2][1], data[2][2] = strconv.FormatInt(int64(comparisons), 10), strconv.FormatInt(int64(permutations), 10)
					data[2][3], data[2][4] = strconv.FormatInt(time, 10), strconv.FormatBool(isSorted)
				case "Естественное 2Ф":
					comparisons, permutations, time, isSorted := NaturalTwoPhaseMergeSort(array_copy)
					data[3][1], data[3][2], data[3][3], data[3][4] = strconv.FormatInt(int64(comparisons), 10), strconv.FormatInt(int64(permutations), 10),
						strconv.FormatInt(time, 10), strconv.FormatBool(isSorted)
				case "Естественное 1Ф":
					comparisons, permutations, time, isSorted := NaturalOnePhaseMergeSort(array_copy)
					data[4][1], data[4][2], data[4][3], data[4][4] = strconv.FormatInt(int64(comparisons), 10), strconv.FormatInt(int64(permutations), 10),
						strconv.FormatInt(time, 10), strconv.FormatBool(isSorted)
				case "Поглощение":
					comparisons, permutations, time, isSorted := MergeInsertionSort(array_copy, int(float64(len(array_copy))*float64(blockSize)*0.01))
					data[5][1], data[5][2] = strconv.FormatInt(int64(comparisons), 10), strconv.FormatInt(int64(permutations), 10)
					data[5][3], data[5][4] = strconv.FormatInt(time, 10), strconv.FormatBool(isSorted)
				}
			}
			table.Refresh()
		}
	})

	arraySize := widget.NewEntry()
	arraySize.SetPlaceHolder("Размер массива")
	arraySize.Bind(sizeBinding)

	percent := widget.NewEntry()
	percent.SetPlaceHolder("% ОП")
	percent.Bind(percentBinding)

	buttonExit := widget.NewButton("Выход", func() {
		myApp.Quit()
	})

	inputContainer := container.New(layout.NewAdaptiveGridLayout(4), buttonSort, arraySize, percent, buttonExit)

	content := container.NewVBox(
		topContainer,
		inputContainer,
	)

	myWindow.Resize(fyne.Size{Width: 870, Height: 300})
	myWindow.SetContent(container.NewVScroll(content))
	myWindow.ShowAndRun()
}
