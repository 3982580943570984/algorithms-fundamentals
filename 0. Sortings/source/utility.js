/**
* Generate random array
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

/**
 * Replace inner text in HTML tag
 */
export const replaceText = (selector, text) => {
    const element = document.getElementById(selector);
    if (element) element.innerText = text;
}
