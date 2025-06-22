export function convertListToString(list, delimiter) {
    return list.join(delimiter);
}

export function convertStringToList(str, delimiter) {
    return str.trimEnd().split(delimiter);
}