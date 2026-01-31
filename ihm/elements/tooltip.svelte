<script>
    export let text = "";
    export let placement = "top"; // top | bottom | left | right

    let visible = false;

    function show() {
        visible = true;
    }

    function hide() {
        visible = false;
    }
</script>

<div class="wrapper" role="presentation" on:mouseenter={show} on:mouseleave={hide}>
    <slot></slot>

    {#if visible}
        <div class="tooltip {placement} visible">
            {@html text}
        </div>
    {/if}
</div>

<style>
    .wrapper {
        position: relative;
        display: inline-block;
    }

    :global(.tooltip) {
        position: absolute;
        min-width: 600px;
        max-width: 800px;
        background: #333;
        color: white;
        padding: 6px 10px;
        border-radius: 6px;
        font-size: 0.85rem;
        white-space: nowrap;
        text-wrap: wrap;
        z-index: 1000;
        opacity: 0;
        transform: translate(-50%, -6px);
        transition: opacity 0.15s ease;
        pointer-events: none;
    }

    :global(.tooltip.visible) {
        opacity: 1;
    }

    /* placements */
    .top    { bottom: 100%; left: 50%; transform: translate(-50%, -6px); }
    .bottom { top: 100%;    left: 50%; transform: translate(-50%, 6px); }
    .left   { right: 100%;  top: 50%;  transform: translate(-6px, -50%); }
    .right  { left: 100%;   top: 50%;  transform: translate(6px, -50%); }
</style>
