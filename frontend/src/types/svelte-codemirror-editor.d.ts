declare module 'svelte-codemirror-editor' {
    import { SvelteComponentTyped } from 'svelte';

    export default class CodeMirror extends SvelteComponentTyped<{
        value?: string;
        options?: Record<string, any>;
        [key: string]: any;
    }> {}
}
