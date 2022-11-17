<script>
	import ButtonWdrop from './ButtonWdropdown.svelte';
	import dijkstra from '$lib/board/algorithms/pathfinders/dijkstra.js';
	import astar from '$lib/board/algorithms/pathfinders/astar.js';
	import depthFirstPath from '$lib/board/algorithms/pathfinders/depthFirst';
	import { statusStore } from '$lib/board/stores';
	let algoList = ['Dijkstra', 'A*', 'Depth First'];
	let activeAlgo = 'Dijkstra';
	function changeAlgo(event) {
		activeAlgo = event.detail.text;
	}
	function clickHandler() {
		if (activeAlgo == 'Dijkstra') dijkstra();
		if (activeAlgo == 'A*') astar();
		if (activeAlgo == 'Depth First') depthFirstPath();
	}
	let status;
	function setStatus() {
		status = $statusStore;
		let newStatus = `Solve with ${activeAlgo}`;
		if (activeAlgo == 'Depth First') newStatus += ' Search (shortest path not guaranteed)';
		$statusStore = newStatus;
	}
	function resetStatus() {
		$statusStore = status;
	}
</script>

<ButtonWdrop
	style="width:180px"
	dropdownList={algoList}
	on:mouseenter={setStatus}
	on:mouseleave={resetStatus}
	on:click={clickHandler}
	on:algoritm={changeAlgo}>Solve with {activeAlgo}</ButtonWdrop
>
