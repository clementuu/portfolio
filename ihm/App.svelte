<!--permet d'acaller le composant svelte dans le fichier html avec une simple balise-->
<svelte:options customElement="index-portfolio" />

<script>	
	import { onMount } from 'svelte';

	let name = "Clément Calia";
	let inge = "Expert en ingénierie logicielle";
	let fullstack = "Développeur Fullstack";

	let currentH2Text = inge;
	let isFading = false;
	const transitionDuration = 500; // ms
	const displayDuration = 3000; // ms

	onMount(() => {
		const interval = setInterval(() => {
			isFading = true; // Start fade-out

			setTimeout(() => {
				currentH2Text = (currentH2Text === inge) ? fullstack : inge;
				isFading = false; // Start fade-in
			}, transitionDuration); // Wait for fade-out to complete before changing text
		}, displayDuration + transitionDuration); // Total cycle duration

		return () => clearInterval(interval); // Cleanup on component unmount
	});
</script>

<div class="banner">
	<div class="banner-content">
		<h1 class="h1">{name}</h1>
		<h2 class="h2" class:fade-out={isFading}>{currentH2Text}</h2>
	</div>
</div>

<style>
	@import './style/style.css';
	@import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

	.banner {
		position: relative;
		height: 300px;
		display: flex;
		align-items: center;
		overflow: hidden;
		background-image: url('../assets/background.webp');
		background-size: auto;
		background-position: center;
	}

	.banner::before {
		content: '';
		position: absolute;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		background-image: inherit;
		background-size: cover;
		background-position: center;
		filter: blur(2px);
		z-index: 0;
	}

	.banner-content {
		position: absolute;
		z-index: 1;
		color: white;
		text-align: left;
		text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.9);
		min-width: fit-content;
		padding-left: 20px; /* Adjust as needed */
	}

	.banner-content h1, .banner-content h2 {
		margin: 0;
		padding: 0.2em 0;
		text-align: left;
		transition: opacity 0.5s ease-in-out;
	}

	.banner-content h2.fade-out {
		opacity: 0;
	}
</style>