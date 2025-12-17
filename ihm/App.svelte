<!--permet d'acaller le composant svelte dans le fichier html avec une simple balise-->
<svelte:options customElement="index-portfolio" />

<script>	
	import { onMount } from 'svelte';
  	import Formations from './formations/formations.svelte';
  	import Experiences from './experiences/experiences.svelte';

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
	<div class="its-me">
		<img class="my-pic" src="./assets/icons8-male-user.svg" alt="">
		<div class="banner-content">
			<h1 class="h1">{name}</h1>
			<h2 class="h2" class:fade-out={isFading}>{currentH2Text}</h2>
		</div>
	</div>
	<span class="parcours">Mon parcours</span>
	<div class="scroll-indicator">
		<i class="bi bi-arrow-down-circle"></i>
	</div>
</div>

<Formations/>
<Experiences/>

<style>
	@import url("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css");
	@import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

	.banner {
		position: relative;
		height: calc(100vh - var(--header-height));
		display: flex;
		flex-direction: column;
		overflow: hidden;
		background: url('../assets/background.webp') center/cover no-repeat;
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
		z-index: 0;
	}

	.banner-content {
		position: relative; /* Changed from absolute to relative */
		z-index: 1;
		color: white;
		text-align: left;
		text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.9);
		min-width: fit-content;
		padding-left: 20px;
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

	.parcours {
		align-self: center;
		margin-top: auto;
		margin-bottom: 0.5em;
		z-index: 1;
		color: white;
		font-size: 1.5rem;
		text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.9);
	}

	.scroll-indicator {
		align-self: center;
		margin-bottom: 1rem;
		font-size: 1.5rem;
		color: white;
		text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.9);
		z-index: 1;
		animation: bounce 2s infinite;
	}

	@keyframes bounce {
		0%, 20%, 50%, 80%, 100% {
			transform: translateY(0);
		}
		40% {
			transform: translateY(-10px);
		}
		60% {
			transform: translateY(-5px);
		}
	}

	.its-me {
		display: flex;
		z-index: 1;
		margin: auto auto auto 0;
	}

	.my-pic {
		margin-left: 2rem;
		width: 150px;
		height: 150px;
		background-color: whitesmoke;
		border-radius: 50%;
		border: 3px solid rgba(255, 255, 255, 0.7);
		box-shadow: 0 0 0 5px rgba(255, 255, 255, 0.3),
		            0 4px 10px rgba(0, 0, 0, 0.5);
		object-fit: cover;
		transition: all 0.3s ease;
	}

	.my-pic:hover {
		transform: scale(1.05);
		box-shadow: 0 0 0 8px rgba(255, 255, 255, 0.4),
		            0 6px 15px rgba(0, 0, 0, 0.6);
	}
</style>