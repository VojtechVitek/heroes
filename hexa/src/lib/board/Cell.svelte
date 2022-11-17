<!-- svelte-ignore a11y-mouse-events-have-key-events -->
<script>
	import Logs from '$lib/icons/Logs.svelte';
	import Mountain from '$lib/icons/Mountain.svelte';
	import Start from '$lib/icons/Start.svelte';
	import Target from '$lib/icons/Target.svelte';

	export let node;
</script>

<div class={`${node.type} ${node.classes}`} on:mouseover on:mousedown on:mouseleave on:mouseenter>
	{#if node.type == 'start'}
		<Start id="start" />
	{:else if node.type == 'mountain'}
		<Mountain width="20px" height="20px" fill="black" />
	{:else if node.type == 'logs'}
		<Logs width="20px" height="20px" fill="black" />
	{:else if node.type == 'target'}
		<Target />
	{/if}
</div>

<style>
	div {
		border-right: solid 1px;
		border-bottom: solid 1px;
		border-color: lightgrey;
		height: 30px;
		aspect-ratio: 1;
	}

	.start,
	.target {
		cursor: grab;
	}
	.start:active,
	.target:active {
		cursor: grabbing;
	}

	.wall {
		background-color: grey;
		/* transition: background-color 1s; */
	}
	.digger {
		background: red;
	}

	.checked {
		/* background: rgb(57, 99, 57); */
		background: rgb(194, 247, 194);
		/* animation: color-me-in 0.5s; */
	}
	.checked2 {
		/* background: rgb(45, 79, 45); */
		background: lightgreen;
	}
	.shortestPath {
		/* background-color: rgb(162, 231, 162); */
		background-color: green;
		/* transition: background-color 2s; */
	}
</style>
