<svelte:options customElement="index-portfolio" />

<script>	
	import { onMount } from 'svelte';
    import Cv from './cv/cv.svelte';
 	import Header from './elements/header.svelte';

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

<Header/>

<div class="banner">
	<div class="its-me">
		<img class="my-pic" src="./assets/photo_cv.png" alt="">
		<div class="banner-content">
			<h1 class="name"><b>{name}</b></h1>
			<h2 class="metier" class:fade-out={isFading}><b>{currentH2Text}</b></h2>
		</div>
	</div>
	<span class="parcours">Mon parcours</span>
	<div class="scroll-indicator">
		<i class="bi bi-arrow-down-circle"></i>
	</div>
</div>

<div class="cv-container">
	<Cv/>
</div>

<style>
	@import "./style/style.css";
	@import url("https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css");
	@import url("https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css");

	.cv-container {
		padding-top: var(--header-height);
	}

	.banner {
		position: relative;
		top: var(--header-height);
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
		text-align: center;
		text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.9);
		min-width: 0;
		max-width: 100%;
		margin-left: 0;
	}

	.banner-content h1, .banner-content h2 {
		margin: 0;
		padding: 0.2em 0;
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

	@media (max-width: 768px) {
		.banner {
			top:0
		}

		.cv-container {
			padding-top: 0;
		}
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
		flex-direction: column;
		align-items: center;
		z-index: 1;
		margin: 10% 10% 0 10%;
	}

	@media (min-width: 1000px) {
		.its-me {
			flex-direction: row;
		}

		.my-pic {
			margin-left: 2.5rem;
		}

		.banner-content {
			text-align: left;
			padding-left: 20px;
		}
	}

	.metier, .name {
		white-space: normal;
		word-wrap: break-word;
	}

	.name {
		font-size: clamp(30px, 15vw, 60px);
	}

	.metier {
		font-size: clamp(15px, 8vw, 40px);
	}
</style>