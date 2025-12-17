<!--permet d'acaller le composant svelte dans le fichier html avec une simple balise-->
<svelte:options customElement="index-portfolio" />

<script>	
	import { onMount } from 'svelte';
  import Formations from './formations/formations.svelte';

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
<div class="cv-container">

	<section class="experiences">
		<h2>Expériences professionnelles</h2>
		<article>
			<h3>Développeur fullstack [Alternance]</h3>
			<p><strong>Softinnov | mars 2024 à mars 2026</strong></p>
			<ul>
				<li>Conception et développement d'applications et de solutions web pour divers clients.</li>
				<li>Administration de systèmes, migration de base de données.</li>
			</ul>
		</article>
		<article>
			<h3>Etude du gaspillage alimentaire [Service Civique]</h3>
			<p><strong>Mairie de Saint-Jean de Védas | novembre 2022 à juillet 2023</strong></p>
			<ul>
				<li>Diagnostique du gaspillage alimentaire à travers la collecte et l'étude de données.</li>
				<li>Rédaction d'un rapport et présentation des résultats aux élus locaux.</li>
				<li>Création d'un jeu vidéo éducatif en javascript pour sensibiliser aux écogestes.</li>
			</ul>
		</article>
		<article>
			<h3>Support sur un projet d'innovation [Stage]</h3>
			<p><strong>IDEMIA – R&D Sophia-Antipolis | mai 2022 à août 2022</strong></p>
			<ul>
				<li>Création et automatisation de tests d'UI.</li>
				<li>Utilisation d'un framework de test « End to End » (Python, Selenium, Serenity, Behave).</li>
			</ul>
		</article>
	</section>
</div>

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

    .cv-container {
        display: flex;
        flex-wrap: wrap;
        max-width: 1200px;
        margin: 2rem auto;
        overflow: hidden;
    }

    .cv-container h2, .cv-container h3 {
        padding-bottom: 0.5rem;
        margin-bottom: 1rem;
    }
    
    .cv-container h2 {
        font-size: 1.8rem;
    }

    .cv-container h3 {
        font-size: 1.2rem;
        border-bottom: none;
        margin-top: 1.5rem;
    }

    .cv-container section {
        margin-bottom: 2rem;
    }

    .cv-container ul {
        padding-left: 20px;
    }

    .cv-container li {
        margin-bottom: 0.5rem;
    }

    .cv-container article {
        margin-bottom: 1.5rem;
    }

    @media (max-width: 768px) {
        .cv-container {
            flex-direction: column;
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
		border-radius: 50%; /* Make it circular */
		border: 3px solid rgba(255, 255, 255, 0.7); /* Softer, slightly transparent border */
		box-shadow: 0 0 0 5px rgba(255, 255, 255, 0.3), /* Halo effect */
		            0 4px 10px rgba(0, 0, 0, 0.5); /* Depth shadow */
		object-fit: cover; /* Ensures image covers the circular area */
		transition: all 0.3s ease; /* Smooth transitions for hover effects */
	}

	.my-pic:hover {
		transform: scale(1.05);
		box-shadow: 0 0 0 8px rgba(255, 255, 255, 0.4), /* More pronounced halo */
		            0 6px 15px rgba(0, 0, 0, 0.6); /* More depth */
	}
</style>