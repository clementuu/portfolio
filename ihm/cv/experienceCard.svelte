<script>
    import Card from './card.svelte';
    export let experience;
</script>

<Card>
    <span slot="title">{experience.Intitule} [{experience.Type}]</span>
    <span slot="subtitle">{experience.Structure} | {experience.Periode}</span>
    <ul>
        {#each experience.Taches as tache}
            <li>{tache}</li>
        {/each}
    </ul>
</Card>
