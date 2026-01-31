<script>
  import Tooltip from '../elements/tooltip.svelte';
    import Card from './card.svelte';
    export let formation;
</script>

<Card className="flex">
    <Tooltip html={formation.Presentation} placement="top">
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
        
    </Tooltip>
    <div class="img-container">
        <img src="{formation.Logo}" alt="{formation.Intitule}">
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

    @media (min-width: 800px) {
		.img-container {
			display: flex;
		}
	}
</style>
