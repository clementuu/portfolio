<script>
    import { onMount } from 'svelte';
    import FormationCard from './formationCard.svelte';
    import ExperienceCard from './experienceCard.svelte';
    import OnScrollAppear from '../elements/OnScrollAppear.svelte';

    let formations = [];
    let error = '';

    async function getFormations() {
        try {
            const response = await fetch('/formations');
            if (!response.ok) {
                throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
            }
            const data = await response.json();
            formations = data;
        } catch (e) {
            console.error("Échec de la récupération des formations :", e.message || e);
            error = "Impossible de charger les formations. Veuillez réessayer plus tard.";
        }
    }

    let experiences = [];

    async function getExperiences() {
        try {
            const response = await fetch('/experiences');
            if (!response.ok) {
                throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
            }
            const data = await response.json();
            experiences = data;
        } catch (e) {
            console.error("Échec de la récupération des expériences :", e.message || e);
            error = "Impossible de charger les expériences. Veuillez réessayer plus tard.";
        }
    }

    onMount(()=>{
        getFormations();
        getExperiences();
    });
</script>

<section>
    <OnScrollAppear>
        <h2>Formations</h2>
    </OnScrollAppear>
    {#if error}
        <p class="error">{error}</p>
    {:else if formations.length > 0}
        {#each formations as formation, index}
            <OnScrollAppear {index}>
                <FormationCard {formation} />
            </OnScrollAppear>
        {/each}
    {:else}
        <p>Chargement des formations...</p>
    {/if}
</section>

<section>
    <OnScrollAppear>
        <h2>Expériences professionnelles</h2>
    </OnScrollAppear>
    {#if error}
        <p class="error">{error}</p>
    {:else if experiences.length > 0}
        {#each experiences as experience, index}
            <OnScrollAppear {index}>
                <ExperienceCard {experience} />
            </OnScrollAppear>
        {/each}
    {:else}
        <p>Chargement des expériences...</p>
    {/if}
</section>

<style>
    section {
        margin: 2rem;
    }
    h2 {
        font-size: 2.5rem; /* Larger font size */
        font-weight: 700; /* Bolder */
        text-align: center; /* Center the title */
        padding-bottom: 0.5rem;
        margin-bottom: 2rem; /* More space below the title */
    }
    .error {
        color: red;
        text-align: center;
    }
</style>