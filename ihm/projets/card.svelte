<script>
    export let projet;
</script>

<button on:click={() => window.location=`/projets/detail.html?id=${projet.id}`} class="project-card">
    <div class="image-container">
        <img src={projet.image} alt={projet.name} />
    </div>
    <h3>{projet.name}</h3>
    <br>
    <p>{projet.resume}</p>
    <div class="competences">
        {#each projet.competences as competence}
            <span class={`card-span ${competence.type === 'Dev' ? 'technique' : competence.type === 'DevOps' ? 'devops' : 'humain'}`}>{competence.name}</span>
        {/each}
    </div>
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

    .image-container {
        height: 70px;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        margin-bottom: 1em;
    }

    .project-card img {
        max-width: 100%;
        max-height: 100%;
        object-fit: contain;
        border-radius: 4px;
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