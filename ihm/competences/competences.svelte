<svelte:options customElement="competences-portfolio" />

<script>
  import Card from './card.svelte';

  const competences = [
    {
      name: "Go",
      image: "",
      rating: 5
    },
    {
        name: "Java",
        image: "",
        rating: 3
    },
    {
        name: "Spring",
        image: "",
        rating: 3
    },
    
    {
      name: "HTML/CSS",
      image: "",
      rating: 5
    },
    {
      name: "JavaScript",
      image: "",
      rating: 4
    },
    {
      name: "Svelte",
      image: "",
      rating: 4
    },
    {
        name: "SQL",
        image: "",
        rating: 4
    },
    {
      name: "Docker",
      image: "",
      rating: 4
    },
    {
        name: "Kubernetes",
        image: "",
        rating: 2
    }
  ];
</script>

<div>
    <h2>Mes Compétences</h2>
    <div class="competences-grid">
        {#each competences as competence}
            <Card
                name={competence.name}
                image={competence.image}
                rating={competence.rating}
            />
        {/each}
    </div>
</div>

<style>
  .competences-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
    gap: 1em;
    justify-items: center;
    padding: 1em;
  }
</style>