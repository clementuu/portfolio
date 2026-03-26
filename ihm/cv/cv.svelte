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

<div class="cv-layout">
    <CardGroup>
        <div class="timeline-line"></div>
        <section>
            <OnScrollAppear index={formationTitleIndex}>
                <h2 id="formations-title"><b>Formations</b></h2>
            </OnScrollAppear>
            {#if error}
                <p class="error">{error}</p>
            {:else if cv.Formations.length > 0}
                {#each cv.Formations as formation, i}
                    <OnScrollAppear index={formationStartIndex + i}>
                        <div class="card-with-dot">
                            <FormationCard {formation}/>
                            <div class="dot"></div>
                        </div>
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
                        <div class="card-with-dot">
                            <ExperienceCard {experience}/>
                            <div class="dot"></div>
                        </div>
                    </OnScrollAppear>
                {/each}
            {:else}
                <p>Chargement des expériences...</p>
            {/if}
        </section>
    </CardGroup>
</div>

<style>
    .cv-layout {
        position: relative;
    }

    .timeline-line {
        position: absolute;
        left: 1.5rem;
        top: 4rem;
        bottom: 1rem;
        width: 4px;
        background-color: var(--primary-color);
        z-index: 1;
    }

    .timeline-line::before {
        content: "";
        position: absolute;
        top: -12px;
        left: 50%;
        transform: translateX(-50%);
        border-left: 8px solid transparent;
        border-right: 8px solid transparent;
        border-bottom: 14px solid var(--primary-color);
    }

    .card-with-dot {
        position: relative;
    }

    .dot {
        position: absolute;
        left: -1.88rem;
        top: 50%;
        transform: translateY(-50%);
        width: 1rem;
        height: 1rem;
        background-color: var(--primary-color);
        border-radius: 50%;
        border: 2px solid white;
        box-shadow: 0 0 4px rgba(0,0,0,0.2);
        z-index: 2;
    }

    @media (max-width: 800px) {
        .timeline-line, .dot {
            display: none;
        }
    }

    section {
        margin: 3rem;
    }
    h2 {
        font-size: clamp(1rem, 8vw, 2.5rem);
        font-weight: 700;
        text-align: center;
        padding-bottom: 0.5em;
        margin-bottom: 1rem;
        padding: 0;
    }

    .error {
        color: red;
        text-align: center;
    }
</style>