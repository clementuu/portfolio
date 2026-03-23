<script>
    import { onMount } from 'svelte';
    import FormationCard from './formationCard.svelte';
    import ExperienceCard from './experienceCard.svelte';
    import OnScrollAppear from '../elements/OnScrollAppear.svelte';
    import CardGroup from '../elements/cardGroup.svelte';

    let cv = {Formations: [], Experiences: []};
    let error = '';
    let loaded = false;

    async function getCv() {
        try {
            const response = await fetch('/cv');
            if (!response.ok) {
                throw new Error(`Erreur HTTP ${response.status}: ${response.statusText}`);
            }
            const data = await response.json();
            cv = data;
        } catch (e) {
            console.error("Échec de la récupération du cv :", e.message || e);
            error = "Impossible de charger le cv. Veuillez réessayer plus tard.";
        }
    }

    onMount(()=>{
        getCv();
        loaded = true;
    });

    // calculs d'offset
    let formationTitleIndex = 0;
    let formationStartIndex = formationTitleIndex + 1;
    $: experienceTitleIndex = formationStartIndex + cv.Formations.length;
    $: experienceStartIndex = experienceTitleIndex + 1;
</script>

<CardGroup>
    <section>
        <OnScrollAppear index={formationTitleIndex}>
            <h2 id="formations-title"><b>Formations</b></h2>
        </OnScrollAppear>
        {#if error}
            <p class="error">{error}</p>
        {:else if cv.Formations.length > 0}
            {#each cv.Formations as formation, i}
                <OnScrollAppear index={formationStartIndex + i}>
                    <FormationCard {formation}/>
                </OnScrollAppear>
            {/each}
        {:else}
            <p>Chargement des formations...</p>
        {/if}
    </section>

    <section>
        <OnScrollAppear index={experienceTitleIndex}>
            <h2><b>Expériences professionnelles</b></h2>
        </OnScrollAppear>
        {#if error}
            <p class="error">{error}</p>
        {:else if cv.Experiences.length > 0}
            {#each cv.Experiences as experience, i}
                <OnScrollAppear index={experienceStartIndex + i}>
                    <ExperienceCard {experience}/>
                </OnScrollAppear>
            {/each}
        {:else}
            <p>Chargement des expériences...</p>
        {/if}
    </section>
</CardGroup>

<style>
    section {
        margin: 2rem;
    }
    h2 {
        font-size: clamp(1rem, 8vw, 2.5rem); /* Larger font size */
        font-weight: 700; /* Bolder */
        text-align: center; /* Center the title */
        padding-bottom: 0.5em;
        margin-bottom: 1rem; /* More space below the title */
        padding: 0;
    }

    .error {
        color: red;
        text-align: center;
    }
</style>