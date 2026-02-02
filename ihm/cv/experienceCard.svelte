<script>
    import Card from '../elements/card.svelte';
    export let experience;
</script>

<Card id={experience.ID}>
    <div slot="hawkins" >
        <div>
            <h3>      
                <span>{experience.Intitule} [{experience.Type}]</span>
            </h3>
            <p>
                <strong>
                    <span><span style="color: var(--primary-color);">{experience.Structure}</span> | {experience.Periode}</span>
                </strong>
            </p>
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
