<script>
    import { onMount } from "svelte";
    import { projets, loadProjets } from "../projets/store.js";

    onMount(async ()=> {
        await loadProjets();
    })
</script>

<nav class="header">
    <div class="acceuil-link">
        <a href="/"><img class="my-pic mini" src="../assets/photo_cv.png" alt=""><span class="link-text">Clément Calia</span></a>
    </div>
    <div class="header-right">
        <a href="/about.html"><i class="bi bi-person-vcard-fill"></i><span class="link-text">À propos</span></a>
        <div class="dropdown">
            <a href="/projets/projets.html"><i class="bi bi-kanban-fill"></i><span class="link-text">Projets</span></a>
            <div class="dropdown-content">
                <ul class="list-group list-group-flush">
                    {#each $projets as projet}
                        <li class="list-group-item"><a class="dropdown-link" href="/projets/detail.html?id={projet.ID}">{projet.Name}</a></li>
                    {/each}
                </ul>
            </div>
        </div>
        <div class="dropdown">
            <a href="/competences/competences.html"><i class="bi bi-person-badge"></i><span class="link-text">Compétences</span></a>
            <div class="dropdown-content">
                <ul class="list-group list-group-flush">
                    <li class="list-group-item"><a class="dropdown-link" href="/competences/competences.html?type=dev">Développement</a></li>
                    <li class="list-group-item"><a class="dropdown-link" href="/competences/competences.html?type=devops">DevOps</a></li>
                    <li class="list-group-item"><a class="dropdown-link" href="/competences/competences.html?type=humain">Soft skills</a></li>
                </ul>
            </div>
        </div>
        <a href="/contacts.html"><i class="bi bi-envelope-fill"></i><span class="link-text">Contacts</span></a>
    </div>
</nav>

<style>
    @import "../style/style.css";
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");
    @import url("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.11.3/font/bootstrap-icons.min.css");

    nav.header {
        position: fixed;
        top: 0;
        bottom: auto;
        left: 0;
        right: 0;
        z-index: 1001;
        background-color: #333;
        padding: 1em;
        height: var(--header-height);
        color: white;
        display: flex;
        justify-content: space-between;
        align-items: center;
        gap: 1em;
    }

    .header-right {
        display: flex;
        align-items: center;
        gap: 1em;
        margin-right: 1em;
    }

    .acceuil-link {
        display: flex;
        align-items: center;
        margin-left: 1em;
    }
    
    li {
        margin: 0;
        padding: 0;
    }

    a {
        color: white;
        text-decoration: none;
        font-size: 16px;
        padding: 0.5em 1em;
        border-radius: 5px;
        display: inline-flex;
        align-items: center;
        gap: 0.5em;
    }

    a.dropdown-link {
        border-radius: 0;
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
        top: 2.5em;
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

    @media (max-width: 800px) {
        nav.header {
            position: fixed;
            top: auto;
            bottom: 0;
            left: 0;
            right: 0;
            justify-content: center;
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

        .header-right{
            gap: 0;
        }
    }

    .mini{
        width: 32px;
        height: 32px;
        object-position: center;
        border: none;
        box-shadow: none;
    }
</style>