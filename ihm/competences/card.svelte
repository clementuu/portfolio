<script>
  export let competence;

  // Function to generate star icons
  function getStars(rating) {
    let stars = '';
    for (let i = 0; i < 5; i++) {
      if (i < rating) {
        stars += '⭐'; // Filled star
      } else {
        stars += '☆'; // Empty star
      }
    }
    return stars;
  }

  function handleClick() {
    window.location.href = `detail.html?id=${competence.id}`;
  }
</script>

<button class="competence-card" on:click={handleClick} type="button" tabindex="0">
    <div class="image-container">
        <img src={competence.image} alt={competence.name} />
    </div>
    <h5>{competence.name}</h5>
    <div class="stars">{getStars(competence.rating)}</div>
    <span class={`card-span ${competence.type === 'Dev' ? 'technique' : competence.type === 'DevOps' ? 'devops' : 'humain'}`}>#{competence.type}</span>
</button>

<style>
	@import '../style/style.css';
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

    .competence-card {
        border: 1px solid #eee;
        border-radius: 8px;
        padding: 2em 0;
        margin: 0.8em;
        box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.1);
        display: flex;
        flex-direction: column;
        align-items: center;
        text-align: center;
        width: 150px;
        transition: transform 0.2s;
        cursor: pointer;
    }

    .competence-card:hover {
        transform: scale(1.05);
    }

    .competence-card img {
        max-width: 100px;
        max-height: 100%;
        object-fit: contain;
        border-radius: 4px;
    }

    .image-container {
        height: 50px;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
        margin-bottom: 1em;
    }

    .competence-card h5 {
        margin-top: 0;
        color: #333;
    }

    .competence-card .stars {
        font-size: 1.2em;
        color: gold;
    }
</style>