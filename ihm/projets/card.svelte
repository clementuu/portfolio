<script>
  export let title;
  export let summary;
  export let image;
</script>

<div class="project-card">
  <img src={image} alt={title} />
  <h3>{title}</h3>
  <p>{summary}</p>
</div>

<style>
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
    max-width: 300px; /* Adjust as needed */
  }

  .project-card img {
    max-width: 100%;
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