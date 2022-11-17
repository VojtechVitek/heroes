import { setClearGrid } from '$lib/board/algorithms/utils';
import { gridStore, statusStore } from '$lib/board/stores';
import { get } from 'svelte/store';

export default async function mazeAnimation(nodes) {
	statusStore.set('inProgress');
	let grid = get(gridStore);
	gridStore.set(setClearGrid(grid));
	for (let node of nodes) {
		//await wait();
		node.type = 'wall';
		gridStore.forceUpdate();
	}
	const columns = grid[0].length;
	const rows = grid.length;
	randomEmptyNode(grid, rows, columns).setType('start');
	randomEmptyNode(grid, rows, columns).setType('target');
	statusStore.set('done');
}

function randomEmptyNode(grid, rows, columns) {
	const random = (max) => Math.floor(Math.random() * max);
	const _rows = random(rows);
	const _columns = random(columns);
	if (grid[_rows][_columns].type != 'empty') return randomEmptyNode(grid, rows, columns);
	return grid[_rows][_columns];
}
