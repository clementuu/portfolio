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

	<section class="formations">
		<h2>Formations</h2>
		<article>
			<h3>Mastère Expert en Ingénierie Logicielle</h3>
			<p><strong>ISCOD | 2023 - 2025</strong></p>
		</article>
		<article>
			<h3>Licence de mathématiques et informatique appliquées</h3>
			<p><strong>Paul Valéry Montpellier 3 | 2019 - 2022</strong></p>
		</article>
		<article>
			<h3>Baccalauréat Scientifique</h3>
			<p><strong>Lycée Laetitia Bonaparte | 2013 - 2016</strong></p>
		</article>
	</section>
</div>

<style>
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
</style>