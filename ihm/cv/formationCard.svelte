<script>
    import Card from '../elements/card.svelte';
    export let formation;
</script>

<Card id={formation.ID}>
    <div slot="hawkins" class="flex">
        <div>
            <h3>      
                <span>{formation.Intitule}</span>
            </h3>
            <p>
                <strong>
                    <a href="{formation.URL}" class="link">
                        {formation.Etablissement}
                    </a>
                    <span>| {formation.Periode}</span>
                </strong>
            </p>
        </div>
        <div class="img-container">
            <img src="{formation.Logo}" alt="{formation.Intitule}">
        </div>
    </div>
    <div slot="upside-down">
        {@html formation.Presentation}
    </div>
</Card>

<style>
    img {
        max-height: 60px;
    }
</style>