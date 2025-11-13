<script>
  export let projet;
</script>

<button on:click={() => window.location=`/projets/detail.html?id=${projet.id}`} class="project-card">
  <img src={projet.image} alt={projet.name} />
  <h3>{projet.name}</h3>
  <p>{projet.desc}</p>
</button>

<style>
    @import '../style/style.css';
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

  .project-card {
    border: 1px solid #eee;
    border-radius: 8px;
    padding: 1em;
    margin: 1em;
    box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.1);
    display: flex;
    flex-direction: column;
    align-items: center;
    text-align: center;
    transition: transform 0.2s;
    max-width: 300px; 
    cursor: pointer;
    color: inherit; /* Make link text color same as parent */
    text-decoration: none; /* Remove underline from link */
  }

  .project-card:hover {
    transform: scale(1.05);
    text-decoration: none; /* Ensure no underline on hover */
  }

  .project-card img {
    max-width: 70px;
    height: auto;
    border-radius: 4px;
    margin-bottom: 1em;
  }

  .project-card h3 {
    margin-top: 0;
    color: #333;
  }

  .project-card p {
    color: #666;
    font-size: 0.9em;
  }
</style>