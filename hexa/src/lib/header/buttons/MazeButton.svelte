<script>
	import ButtonWdrop from './ButtonWdropdown.svelte';
	import depthFirst from '$lib/board/algorithms/maze/depthFirst.js';
	import randomMaze from '$lib/board/algorithms/maze/randomMaze.js';
	import recursiveDivision from '$lib/board/algorithms/maze/recursivedivision.js';
	import { statusStore } from '$lib/board/stores';

	let algoList = ['Depth First', 'Random', 'Recursive Division'];
	let activeAlgo = 'Depth First';
	function changeAlgo(event) {
		activeAlgo = event.detail.text;
	}
	function clickHandler() {
		if (activeAlgo == 'Depth First') depthFirst();
		if (activeAlgo == 'Random') randomMaze();
		if (activeAlgo == 'Recursive Division') recursiveDivision();
	}
	let status;
	function setStatus() {
		status = $statusStore;
		if (activeAlgo == 'Depth First')
			$statusStore = 'Draw a maze with a randomized Depth First Search';
		if (activeAlgo == 'Random') $statusStore = 'Draw a maze of random dots';
		if (activeAlgo == 'Recursive Division') $statusStore = 'Draw a maze using Recursive Division';
	}
	function resetStatus() {
		$statusStore = status;
	}
</script>

<ButtonWdrop
	style="width:200px"
	dropdownList={algoList}
	on:mouseenter={setStatus}
	on:mouseleave={resetStatus}
	on:click={clickHandler}
	on:algoritm={changeAlgo}>{activeAlgo} Maze</ButtonWdrop
>
