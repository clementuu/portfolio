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
                    <span style="color: var(--primary-color);">
                        {formation.Etablissement}
                    </span>
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

    .img-container {
        display: none;
        justify-content: center;
        align-items: center;
        height: 100px;
        width: 100px;
        border-radius: 10%;
        background-color: white;
    }

    .flex {
        display: flex;
        justify-content: space-between;
    }

    @media (min-width: 800px) {
		.img-container {
			display: flex;
		}
	}
</style>