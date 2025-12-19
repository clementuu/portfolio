<script>
    import { onMount } from 'svelte';
    export let index = 0;
    export let delay = 150;
    export let threshold = 0.1;

    let visible = false;
    let element;

    onMount(() => {
        const observer = new IntersectionObserver(
            (entries) => {
                if (entries[0].isIntersecting) {
                    visible = true;
                    observer.unobserve(element);
                }
            },
            { threshold }
        );

        observer.observe(element);
        return () => { if (element) observer.unobserve(element) };
    });
</script>

<div bind:this={element} class:visible style="transition-delay: {index * delay}ms;">
    <slot></slot>
</div>

<style>
    div {
        opacity: 0;
        transform: translateY(20px);
        transition: opacity 0.5s ease-out, transform 0.5s ease-out;
        /* Make the div take up the space of its content */
        display: block; 
        width: 100%;
    }
    div.visible {
        opacity: 1;
        transform: translateY(0);
    }
</style>
