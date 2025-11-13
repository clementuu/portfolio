<svelte:options customElement="header-portfolio" />

<nav>
    <a href="/">Accueil</a>
    <div class="dropdown">
        <a href="/projets/projets.html">Projets</a>
        <div class="dropdown-content">
            <a href="/projets/detail.html?id=projet-a">Projet A</a>
            <a href="/projets/detail.html?id=projet-b">Projet B</a>
            <a href="/projets/detail.html?id=projet-c">Projet C</a>
        </div>
    </div>
    <div class="dropdown">
        <a href="/competences/competences.html">Compétences</a>
        <div class="dropdown-content">
            <a href="/competences/competences.html?type=technique">Compétences techniques</a>
            <a href="/competences/competences.html?type=humain">Compétences humaines</a>
        </div>
    </div>
    <a href="/contacts.html">Contacts</a>
</nav>

<style>
    @import '/style/style.css';
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

    nav {
        background-color: #333;
        padding: 1em;
        color: white;
        display: flex;
        justify-content: center;
        align-items: center;
        gap: 1em;
    }

    a {
        color: white;
        text-decoration: none;
        padding: 0.5em 1em;
        border-radius: 5px;
    }

    a:hover {
        background-color: #555;
    }

    /* Dropdown styles */
    .dropdown {
        position: relative;
    }

    .dropdown-content {
        display: none;
        position: absolute;
        background-color: #f9f9f9;
        min-width: 160px;
        box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
        z-index: 1;
        border-radius: 5px;
        top: 2em;
    }

    .dropdown-content a {
        color: black;
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
</style>