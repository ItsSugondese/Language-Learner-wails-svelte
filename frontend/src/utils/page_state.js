// src/utils/pageState.js

export function getPageKey(pathname) {
    return 'appState:' + pathname;
}

export function savePageState(state, pathname = window.location.pathname) {
    localStorage.setItem(getPageKey(pathname), JSON.stringify(state));
}

export async function loadPageState(pathname = window.location.pathname) {
    const saved = localStorage.getItem(getPageKey(pathname));
    return saved ? JSON.parse(saved) : null;
}
