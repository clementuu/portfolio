import { writable } from "svelte/store";

const cached = sessionStorage.getItem("articles");
export const articles = writable(cached ? JSON.parse(cached) : []);

let loaded = cached !== null;

export async function loadArticles() {
    if (loaded) return;

    try {
        const response = await fetch('/articles');
        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }
        const data = await response.json();

        articles.set(data);
        sessionStorage.setItem("articles", JSON.stringify(data));

        loaded = true;
    } catch (err) {
        console.error('Error loading articles:', err);
        throw err;
    }
}
