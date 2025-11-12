<script>
  export let name;
  export let image;
  export let rating; // A number from 0 to 5

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
</script>

<style>
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
    max-width: 200px; /* Adjust as needed */
  }

  .competence-card img {
    max-width: 80px;
    height: auto;
    margin-bottom: 1em;
  }

  .competence-card h3 {
    margin-top: 0;
    color: #333;
  }

  .competence-card .stars {
    font-size: 1.2em;
    color: gold; /* Or any other color for stars */
  }
</style>

<div class="competence-card">
  <img src={image} alt={name} />
  <h3>{name}</h3>
  <div class="stars">{getStars(rating)}</div>
</div>