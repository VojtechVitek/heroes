import { writable } from 'svelte/store';
import dijkstra from './algorithms/pathfinders/dijkstra.js';
import Node from './Node.js';
let gridColumns = 3000;
let gridRows = 3000;

function ini(columns, rows) {
	gridColumns = columns;
	gridRows = rows;
	return newGrid(columns, rows);
}

function newGrid(columns, rows) {
	let grid = [...Array(columns)].map((_, row) =>
		[...Array(rows)].map((_, column) => new Node(row, column))
	);
	let centerX = Math.floor(grid.length / 2);
	let centerY = Math.floor(grid[0].length / 2);
	grid[centerX - 5][centerY].setType('start');
	grid[centerX + 5][centerY].setType('target');
	return grid;
}

function createGrid() {
	const { subscribe, set, update } = writable([]);
	return {
		subscribe,
		init: (columns, rows) => set(ini(columns, rows)),
		set,
		reset: () => set(newGrid(gridColumns, gridRows)),
		forceUpdate: () => update((n) => n)
	};
}

export const reset = () => {
	gridStore.reset();
	statusStore.set('');
};

export const gridStore = createGrid();
export const statusStore = writable('');
export const startStore = writable({ row: undefined, column: undefined });
export const algoStore = writable(dijkstra);
export const endStore = writable({ row: undefined, column: undefined });
export const activeObstacleStore = writable('wall');
export const animationSpeedStore = writable('20');
export const shortestPathStore = writable([]);
