<script>
    import { onMount } from 'svelte';

    export let experience;
    export let index;

    let visible = false;
    let cardNode;

    onMount(() => {
        const observer = new IntersectionObserver(
            (entries) => {
                entries.forEach((entry) => {
                    if (entry.isIntersecting) {
                        visible = true;
                        observer.unobserve(cardNode);
                    }
                });
            },
            {
                threshold: 0.1,
            }
        );

        observer.observe(cardNode);

        return () => {
            if (cardNode) {
                observer.unobserve(cardNode);
            }
        };
    });
</script>

<article bind:this={cardNode} class:visible style="transition-delay: {index * 150}ms;">
    <h3>{experience.Intitule} [{experience.Type}]</h3>
    <p><strong>{experience.Structure} | {experience.Periode}</strong></p>
    <ul>
        {#each experience.Taches as tache}
            <li>{tache}</li>
        {/each}
    </ul>
</article>

<style>
    article {
        background-color: #f9f9f9;
        padding: 1.5rem;
        margin-bottom: 1.5rem;
        border-radius: 8px;
        box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
        opacity: 0;
        transform: translateY(20px);
        transition: opacity 0.5s ease-out, transform 0.5s ease-out, box-shadow 0.2s ease-in-out;
    }

    article.visible {
        opacity: 1;
        transform: translateY(0);
    }
    
    h3 {
        font-size: 1.3rem;
        color: #333;
        margin-top: 0;
        margin-bottom: 0.5rem;
        border-bottom: none;
    }

    p {
        font-size: 1rem;
        color: #555;
        line-height: 1.5;
    }

    strong {
        color: #007bff;
    }

    ul {
        padding-left: 20px;
        margin-top: 1rem;
        color: #666;
    }

    li {
        margin-bottom: 0.5rem;
    }
</style>
