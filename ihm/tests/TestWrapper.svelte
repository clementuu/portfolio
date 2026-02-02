<script>
  import CardGroup from '../elements/cardGroup.svelte';

  export let Component;
  export let props;
</script>

<CardGroup>
  <svelte:component this={Component} {...props} />
</CardGroup>
