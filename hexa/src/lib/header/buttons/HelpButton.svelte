<script>
	import Help from '$lib/help/Help.svelte';
	import RoundButton from './RoundButton.svelte';
	let show = false;
</script>

<RoundButton on:click={() => (show = true)}>?</RoundButton>

{#if show}
	<Help on:close={() => (show = false)} />
{/if}
