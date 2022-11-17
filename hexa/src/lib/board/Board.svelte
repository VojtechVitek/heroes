<script>
	import { setClearPath } from '$lib/board/algorithms/utils';
	import Cell from '$lib/board/Cell.svelte';
	import { onMount } from 'svelte';
	import recursiveDivision from './algorithms/maze/recursivedivision';
	import dijkstra from './algorithms/pathfinders/dijkstra.js';
	import { algoStore, endStore, gridStore, startStore } from './stores.js';

	let width = 200;
	let height = 200;
	onMount(() => {
		gridStore.init(40, 30);
		recursiveDivision();
		dijkstra();
	});

	const setNodeType = (row, column, type) => {
		if (row && column) {
			$gridStore[row][column].setType(type);
			$gridStore[row][column] = $gridStore[row][column];
		}
	};

	let dragNode = null;
	const mouseHandler = (event, row, column) => {
		if (event.buttons != 1 || dragNode !== null) return;
		let currentValue = $gridStore[row][column].type;
		if (currentValue == 'wall' || currentValue == 'start') {
			return;
		}

		if (currentValue == 'target') {
			setNodeType($startStore.row, $startStore.column, 'empty');

			$startStore = $endStore;
			setNodeType($startStore.row, $startStore.column, 'start');

			$endStore = { row: undefined, column: undefined };

			gridStore.set(setClearPath($gridStore));
			gridStore.forceUpdate();
		} else {
			setNodeType($endStore.row, $endStore.column, 'empty');
			setNodeType(row, column, 'target');
			$algoStore();
		}
	};
	const mouseLeaveHandler = (event, row, column) => {
		if (event.buttons != 1) return;
		let currentValue = $gridStore[row][column].type;
		if (['start', 'target'].includes(currentValue)) {
			dragNode = { row, column, type: currentValue };
		}
	};
	const mouseEnterHandler = (event, row, column) => {
		if (event.buttons != 1 || $gridStore[row][column].type !== 'empty' || dragNode == null) return;
		setNodeType(dragNode.row, dragNode.column, 'empty');
		setNodeType(row, column, dragNode.type);
		dragNode = null;
	};
</script>

<div class="container" bind:clientWidth={width} bind:clientHeight={height}>
	<div class="board" on:mouseleave={(dragNode = null)} on:mouseup={(dragNode = null)}>
		{#each $gridStore as row, r}
			<div>
				{#each row as node, c}
					<Cell
						{node}
						on:mousedown={(event) => mouseHandler(event, r, c)}
						on:mouseover={(event) => mouseHandler(event, r, c)}
						on:mouseleave={(event) => mouseLeaveHandler(event, r, c)}
						on:mouseenter={(event) => mouseEnterHandler(event, r, c)}
					/>
				{/each}
			</div>
		{/each}
	</div>
</div>

<style>
	.container {
		flex: 1 1 auto;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	.board {
		display: flex;
		border-left: 1px solid;
		border-top: 1px solid;
		border-color: lightgrey;
	}
</style>
