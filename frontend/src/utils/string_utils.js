export function convertListToString(list, delimiter) {
    return list.join(delimiter);
}

export function convertStringToList(str, delimiter) {
    return str.trimEnd().split(delimiter);
}

export function getNumberFromString(value) {
    const match = value.match(/(\d+)$/);
    if (match) {
        return parseInt(match[1], 10);
    }
    return 0;
}

export function normalizeUmlauts(str) {
    return str
        .toLowerCase()
        .replace(/ä/g, "ae")
        .replace(/ö/g, "oe")
        .replace(/ü/g, "ue")
        .replace(/ß/g, "ss");
}