<svelte:options customElement="header-portfolio" />

<script>
  import { onMount } from "svelte";

    let projets = [];

    async function getProjetsNames() {
        try {
            const response = await fetch('/projets/names');

            if (!response.ok) {
                throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
            }

            const data = await response.json();
            projets = data;
        } catch (error) {
            console.error("Échec de la récupération des compétences :", error.message || error);
        }
    }

    onMount(async ()=> {
        await getProjetsNames();
    })
</script>

<nav class="header">
    <a href="/"><i class="bi bi-house-door-fill"></i><span class="link-text">Accueil</span></a>
    <div class="dropdown">
        <a href="/projets/projets.html"><i class="bi bi-kanban-fill"></i><span class="link-text">Projets</span></a>
        <div class="dropdown-content">
            {#each projets as projet}
                <a href="/projets/detail.html?id={projet.Id}">{projet.Name}</a>
            {/each}
        </div>
    </div>
    <div class="dropdown">
        <a href="/competences/competences.html"><i class="bi bi-person-badge"></i><span class="link-text">Compétences</span></a>
        <div class="dropdown-content">
            <a href="/competences/competences.html?type=technique">Hard skills</a>
            <a href="/competences/competences.html?type=humain">Soft skills</a>
        </div>
    </div>
    <a href="/contacts.html"><i class="bi bi-envelope-fill"></i><span class="link-text">Contacts</span></a>
</nav>

<style>
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");
    @import url("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css");

    nav.header {
        background-color: #333;
        padding: 1em;
        height: var(--header-height);
        color: white;
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 1em;
    }

    a {
        color: white;
        text-decoration: none;
        font-size: 1.2em;
        font-weight: bold;
        padding: 0.5em 1em;
        border-radius: 5px;
        display: inline-flex;
        align-items: center;
        gap: 0.5em;
    }

    a:hover {
        background-color: #555;
    }

    .dropdown {
        position: relative;
    }

    .dropdown-content {
        display: none;
        position: absolute;
        background-color: #f9f9f9;
        min-width: 160px;
        box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
        z-index: 1000;
        border-radius: 5px;
        top: 3em;
    }

    .dropdown-content a {
        color: black;
        font-size: 14px;
        padding: 12px 16px;
        text-decoration: none;
        display: block;
        text-align: left;
    }

    .dropdown-content a:hover {
        background-color: #ddd;
    }

    .dropdown:hover .dropdown-content {
        display: block;
    }

    @media (max-width: 768px) {
        nav.header {
            position: fixed;
            bottom: 0;
            left: 0;
            right: 0;
            justify-content: space-around;
            z-index: 1001;
            height: var(--header-height);
            padding: 0.5em;
        }

        .link-text {
            display: none;
        }

        nav.header a, .dropdown {
            font-size: 1.5rem;
            display: flex;
            justify-content: center;
            flex-grow: 1;
        }

        .dropdown:hover .dropdown-content {
            display: none;
        }
    }
</style>