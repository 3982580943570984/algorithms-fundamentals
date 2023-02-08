/**
* Bubble Sort === Сортировка Обменом
*/
const bubbleComparisons = 0;
const bubbleTransitions = 0;
const bubbleTime = 0;
export const bubblesort = (arr) => {
    const array = arr.slice();
    let swapped = false;
    for (let i = 0; i < array.length; i++) {
        swapped = false;
        for (let j = 0; j < (array.length - i - 1); j++) {
            if (array[j] > array[j + 1]) {
                [array[j], array[j + 1]] = [array[j + 1], array[j]];
                swapped = true;
            }
        }
        if (!swapped) break;
    }
    return [bubbleComparisons, bubbleTransitions, bubbleTime];
};

/**
* Selection Sort === Сортировка Выбором
*/
const selectionComparisons = 0;
const selectionTransitions = 0;
const selectionTime = 0;
export const selectionsort = (arr) => {
    const array = arr.slice();
    for (let i = 0; i < array.length; i++) {
        let index = i;
        for (let j = i + 1; j < array.length; j++)
            if (array[j] < array[index])
                index = j;
        if (index != i)
            [array[index], array[i]] = [array[i], array[index]];
    }
    return [selectionComparisons, selectionTransitions, selectionTime];
};

/**
* Insertion Sort === Сортировка Включением
*/
const insertionComparisons = 0;
const insertionTransitions = 0;
const insertionTime = 0;
export const insertionsort = (arr) => {
    const array = arr.slice();
    for (let i = 1; i < array.length; i++) {
        let j = i;
        while (j > 0 && array[j - 1] > array[j]) {
            [array[j - 1], array[j]] = [array[j], array[j - 1]];
            --j;
        }
    }
    return [insertionComparisons, insertionTransitions, insertionTime];
};

/**
* Quick Sort === Быстрая Сортировка
*/
const quickComparisons = 0;
const quickTransitions = 0;
const quickTime = 0;
export const quicksort = (arr, start, end) => {
    if (start >= end) return;
    let index = partition(arr, start, end);
    quicksort(arr, start, index - 1);
    quicksort(arr, index + 1, end);

    return [quickComparisons, quickTransitions, quickTime];
}

const partition = (arr, start, end) => {
    const pivot = arr[end];
    let index = start;
    for (let i = start; i < end; i++) {
        if (arr[i] < pivot) {
            [arr[i], arr[index]] = [arr[index], arr[i]];
            ++index;
        }
    }

    [arr[index], arr[end]] = [arr[end], arr[index]];

    return index;
}
