/*
import { bubblesort, selectionsort, insertionsort, quicksort } from './sortings';
import { replaceText } from './utility';
*/

const randomArray = (length, max) => {
    return Array(length).fill().map(() => Math.round(Math.random() * max));
}

/**
* Check if list is sorted
*/
const isSorted = (arr) => {
    return arr.every((value, index, array) => {
        return index === 0 || array[index - 1] <= value;
    });
}

const replaceText = (selector, text) => {
    const element = document.getElementById(selector);
    if (element) element.innerText = text;
}

const chooseSorting = (name, array) => {
    if (name === 'bubble') return bubblesort(array);
    if (name === 'selection') return selectionsort(array);
    if (name === 'insertion') return insertionsort(array);
    if (name === 'quick') {
        quickComparisons = 0, quickTransitions = 0, quickTime = window.performance.now()
        return quicksort(array, 0, array.length);
    }
}

const performSortings = () => {
    const size = document.getElementById('size').value;
    const array = randomArray(parseInt(size), parseInt(size));
    const sortings = document.querySelectorAll('input[class=sorting]:checked')

    sortings.forEach(element => {
        const [comparisons, transitions, time] = chooseSorting(element.id, array);
        replaceText(`${element.id}-comparisons`, comparisons);
        replaceText(`${element.id}-transitions`, transitions);
        replaceText(`${element.id}-time`, time);
    });
}

/**
* Bubble Sort === Сортировка Обменом
*/
const bubblesort = (arr) => {
    let bubbleComparisons = 0, bubbleTransitions = 0, bubbleTime = window.performance.now();
    const array = arr.slice();
    let swapped = false;
    for (let i = 0; i < array.length; i++) {
        swapped = false;
        for (let j = 0; j < (array.length - i - 1); j++) {
            if (array[j] > array[j + 1]) {
                [array[j], array[j + 1]] = [array[j + 1], array[j]];
                swapped = true;
                ++bubbleComparisons;
                ++bubbleTransitions;
            }
        }
        if (!swapped) break;
    }
    return [bubbleComparisons, bubbleTransitions, parseInt(window.performance.now() - bubbleTime)];
};

/**
* Selection Sort === Сортировка Выбором
*/
const selectionsort = (arr) => {
    let selectionComparisons = 0, selectionTransitions = 0, selectionTime = window.performance.now();
    const array = arr.slice();
    for (let i = 0; i < array.length; i++) {
        let index = i;
        for (let j = i + 1; j < array.length; j++)
            if (array[j] < array[index]) {
                index = j;
                ++selectionComparisons;
            }
        if (index != i) {
            [array[index], array[i]] = [array[i], array[index]];
            ++selectionComparisons;
            ++selectionTransitions;
        }
    }
    return [selectionComparisons, selectionTransitions, parseInt(window.performance.now() - selectionTime)];
};

/**
* Insertion Sort === Сортировка Включением
*/
const insertionsort = (arr) => {
    let insertionComparisons = 0, insertionTransitions = 0, insertionTime = window.performance.now();
    const array = arr.slice();
    for (let i = 1; i < array.length; i++) {
        let j = i;
        while (j > 0 && array[j - 1] > array[j]) {
            [array[j - 1], array[j]] = [array[j], array[j - 1]];
            --j;
            ++insertionTransitions;
            insertionComparisons += 2;
        }
    }
    return [insertionComparisons, insertionTransitions, parseInt(window.performance.now() - insertionTime)];
};

/**
* Quick Sort === Быстрая Сортировка
*/
let quickComparisons = 0, quickTransitions = 0, quickTime = window.performance.now();
const quicksort = (arr, start, end) => {
    if (start >= end) return;
    let index = partition(arr, start, end);
    quicksort(arr, start, index - 1);
    quicksort(arr, index + 1, end);

    return [quickComparisons, quickTransitions, parseInt(window.performance.now() - quickTime)];
}

const partition = (arr, start, end) => {
    const pivot = arr[end];
    let index = start;
    for (let i = start; i < end; i++) {
        if (arr[i] < pivot) {
            [arr[i], arr[index]] = [arr[index], arr[i]];
            ++index;
            ++quickComparisons;
            ++quickTransitions;
        }
    }

    [arr[index], arr[end]] = [arr[end], arr[index]];
    ++quickTransitions;

    return index;
}