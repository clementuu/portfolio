<script>
    import Card from './card.svelte';
    export let formation;
</script>

<Card>
    <span slot="title">{formation.Intitule}</span>
    <span slot="subtitle"><span style="color: var(--primary-color);">{formation.Etablissement}</span> | {formation.Periode}</span>
</Card>
