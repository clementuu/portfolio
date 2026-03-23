<script>
    import Card from '../elements/card.svelte';
    export let experience;
</script>

<Card id={experience.ID}>
    <div slot="hawkins" >
        <div class="flex">
            <div>
                <h3>      
                    <span>{experience.Intitule} [{experience.Type}]</span>
                </h3>
                <p>
                    <strong>
                        <span><a href="{experience.URL}" class="link">{experience.Structure}</a> | {experience.Periode}</span>
                    </strong>
                </p>
            </div>
            <div class="img-container">
                <img src="{experience.Logo}" alt="{experience.Intitule}">
            </div>
        </div>
        
        <!-- liste des taches -->
        <ul>
            {#each experience.Taches as tache}
                <li>{tache}</li>
            {/each}
        </ul>
    </div>
    <div slot="upside-down">
        {@html experience.Description}
    </div>
</Card>

<style>
    img {
        max-height: 60px;
    }
</style>