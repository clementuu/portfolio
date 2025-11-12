<svelte:options customElement="projets-portfolio" />

<script>
  import Card from './card.svelte';

  const projects = [
    {
      title: "Projet Alpha",
      summary: "Un projet passionnant qui a exploré les limites de la technologie X.",
      image: ""
    },
    {
      title: "Projet Beta",
      summary: "Développement d'une application mobile innovante pour la gestion de tâches.",
      image: ""
    },
    {
      title: "Projet Gamma",
      summary: "Optimisation des performances d'une base de données distribuée.",
      image: ""
    }
  ];
</script>

<div>
    <h2>Mes Projets</h2>
    <div class="projects-grid">
        {#each projects as project}
            <Card
                title={project.title}
                summary={project.summary}
                image={project.image}
            />
        {/each}
    </div>
</div>

<style>
  .projects-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
    gap: 1em;
    justify-items: center;
    padding: 1em;
  }
</style>