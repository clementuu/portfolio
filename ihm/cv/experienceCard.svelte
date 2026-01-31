<script>
    import Tooltip from '../elements/tooltip.svelte';
import Card from './card.svelte';
    export let experience;
</script>

<Card>
    <Tooltip html={experience.Description} placement="top">
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
    </Tooltip>
</Card>
