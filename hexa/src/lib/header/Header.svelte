<script>
	import astar from '$lib/board/algorithms/pathfinders/astar';
	import dijkstra from '$lib/board/algorithms/pathfinders/dijkstra';
	import { algoStore, statusStore } from '$lib/board/stores.js';
	import HelpButton from './buttons/HelpButton.svelte';
	import ObsticalSelector from './buttons/ObsticalSelector.svelte';
	import Status from './Status.svelte';
	let status;
	function setStatus(action) {
		status = $statusStore;
		$statusStore = `${action}`;
	}
	function resetStatus() {
		$statusStore = status;
	}
</script>

<header>
	<input
		type="radio"
		id="dijkstra"
		value="Dijsktra"
		checked={$algoStore == dijkstra}
		on:click={() => {
			$algoStore = dijkstra;
			dijkstra();
		}}
	/>
	<label for="dijkstra">Dijkstra</label>
	<input
		type="radio"
		id="astar"
		value="A*"
		checked={$algoStore == astar}
		on:click={() => {
			$algoStore = astar;
			astar();
		}}
	/>
	<label for="astar">A*</label>

	<ObsticalSelector />
	<HelpButton />
</header>
<Status />

<style>
	header {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 10px;
		height: 50px;
		background-color: var(--header-bg-color);
		/* color: var(--button-fg-color); */
		color: black;
		box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
	}
</style>
