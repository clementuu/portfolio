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
    sessionStorage.setItem('selectedCompetence', JSON.stringify(competence));
    window.location.href = 'detail.html';
  }
</script>

<button class="competence-card" on:click={handleClick} type="button" tabindex="0">
    <div class="image-wrapper">
        <img src={competence.image} alt={competence.name} />
    </div>
    <h4>{competence.name}</h4>
    <div class="stars">{getStars(competence.rating)}</div>
    <span class={competence.type === 'Technique' ? 'technique' : 'humain'}>#{competence.type}</span>
</button>

<style>
	@import '../style/style.css';
    @import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

    .competence-card {
        border: 1px solid #eee;
        border-radius: 8px;
        padding: 1em;
        margin: 1em;
        box-shadow: 2px 2px 8px rgba(0, 0, 0, 0.1);
        display: flex;
        flex-direction: column;
        align-items: center;
        text-align: center;
        width: 180px;
        transition: transform 0.2s;
        cursor: pointer;
    }

    .competence-card:hover {
        transform: scale(1.05);
    }

    .competence-card img {
        max-width: 50px;
        max-height: 50px;
        margin-bottom: 1em;
    }

    .competence-card span {
        margin-top: 1em;
        padding: 2px;
        border-radius: 5%;
    }

    .technique {
        background-color: darkgrey;
    }

    .humain {
        background-color: darkturquoise;
    }

    .image-wrapper {
        height: 50px;
        display: flex;
        justify-content: center;
    }

    .competence-card h4 {
        margin-top: 0;
        color: #333;
    }

    .competence-card .stars {
        font-size: 1.2em;
        color: gold; /* Or any other color for stars */
    }
</style>