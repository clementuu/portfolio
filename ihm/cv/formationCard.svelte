<script>
    import Card from './card.svelte';
    export let formation;
</script>

<Card>
    <span slot="title">{formation.Intitule}</span>
    <span slot="subtitle">{formation.Etablissement} | {formation.Periode}</span>
</Card>
