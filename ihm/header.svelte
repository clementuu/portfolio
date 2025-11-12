<svelte:options customElement="header-portfolio" />

<script>
</script>

<nav>
  <a href="/">Accueil</a>
  <a href="/projets.html">Projets</a>
  <a href="/competences.html">Compétences</a>
  <a href="/contacts.html">Contacts</a>
</nav>

<style>
  nav {
    background-color: #333;
    padding: 1em;
    color: white;
    display: flex;
    justify-content: center;
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
</style>