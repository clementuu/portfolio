<script>
    import { tick } from 'svelte';

    export let html = "";
    export let placement = "top"; // top | bottom | left | right

    let visible = false;
    let tooltipEl;
    let wrapperEl;

    async function show() {
        visible = true;
        await tick();
        if (tooltipEl && wrapperEl) {
            positionTooltip();
        }
    }

    function hide() {
        visible = false;
    }

    function positionTooltip() {
        const wrapperRect = wrapperEl.getBoundingClientRect();

        if (placement === 'left' || placement === 'right') {
            return;
        }

        const tooltipRect = tooltipEl.getBoundingClientRect();
        
        let left = 0;

        if (wrapperRect.left + tooltipRect.width > window.innerWidth) {
            left = (window.innerWidth - tooltipRect.width) - wrapperRect.left - 5;
        }

        if (wrapperRect.left + left < 0) {
            left = -wrapperRect.left + 5;
        }

        tooltipEl.style.left = `${left}px`;
    }
</script>

<div class="wrapper" role="presentation" on:mouseenter={show} on:mouseleave={hide} bind:this={wrapperEl}>
    <slot></slot>

    {#if visible}
        <div class="tooltip {placement} visible" bind:this={tooltipEl}>
            {@html html}
        </div>
    {/if}
</div>

<style>

    :global(.tooltip) {
        position: absolute;
        min-width: 50vw;
        max-width: 80vw;
        margin: 0;
        background: #333;
        color: white;
        padding: 10px 10px;
        border-radius: 6px;
        font-size: 0.85rem;
        white-space: normal;
        text-wrap: wrap;
        z-index: 2000;
        opacity: 0;
        transition: opacity 0.15s ease;
        pointer-events: none;
    }

    :global(.tooltip.visible) {
        opacity: 1;
    }

    :global(.tooltip ul) {
        margin: 0;
        padding: 0;
    }

    :global(.tooltip li) {
        background-color: #333;
        max-width: fit-content;
        color: white;
        white-space: normal;
    }

    /* placements */
    .top    { bottom: 100%; left: 0; transform: translateY(-6px); }
    .bottom { top: 100%;    left: 0; transform: translateY(6px); }
    .left   { right: 100%;  top: 50%;  transform: translate(-6px, -50%); }
    .right  { left: 100%;   top: 50%;  transform: translate(6px, -50%); }
</style>
