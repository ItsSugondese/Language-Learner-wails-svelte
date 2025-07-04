import {TranslateFromToEnums} from "../../enums/translate_from_to_enums.js";

export function getSplitWordMeaningObject(str, delimiter, direction) {
    let words = str.split(delimiter) ?? null;

    if (words === null) {
        return null;
    }

    return direction === TranslateFromToEnums.EnglishToGerman.name ? Object.freeze({word: words[1].trim(), meaning: words[0].trim()}) : Object.freeze({word: words[0].trim(), meaning: words[1].trim()});
}