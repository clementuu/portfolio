import { writable } from "svelte/store";

const cached = localStorage.getItem("projets");
export const projets = writable(cached ? JSON.parse(cached) : []);

let loaded = cached !== null;

export async function loadProjets() {
    if (loaded) return;

    try {
        const response = await fetch('/projets/names');
        const data = await response.json();

        projets.set(data);
        localStorage.setItem("projets", JSON.stringify(data));

        loaded = true;
    } catch (err) {
        console.error(err);
    }
}
