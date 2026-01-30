<script>
    import Card from './card.svelte';
    export let experience;
</script>

<Card>
    <span slot="title">{experience.Intitule} [{experience.Type}]</span>
    <span slot="subtitle"><span style="color: var(--primary-color);">{experience.Structure}</span> | {experience.Periode}</span>
    <ul>
        {#each experience.Taches as tache}
            <li>{tache}</li>
        {/each}
    </ul>
</Card>
