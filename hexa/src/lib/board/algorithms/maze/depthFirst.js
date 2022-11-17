import depthFirstAnimation from '$lib/board/animations/maze/depthFirstAnimation.js';
import { get } from 'svelte/store';
import { gridStore } from '../../stores.js';
import { setClearGrid } from '../utils.js';

export default function depthFirst() {
	let grid = setClearGrid(get(gridStore));

	let visitedInOrder = []; //All visited, for animation

	let visited = [];
	let current = grid[1][1];
	visitedInOrder.push(current);

	do {
		current.visited = true;
		let [neighbors, walls] = getNeighbors(current);
		if (neighbors.length) {
			visited.push(current);
			let rng = Math.floor(Math.random() * neighbors.length);
			let neighbor = neighbors[rng];
			visitedInOrder.push(walls[rng]);
			visitedInOrder.push(neighbor);
			current = neighbor;
		} else {
			current = visited.pop();
			visitedInOrder.push(current);
		}
	} while (visited.length);
	depthFirstAnimation(visitedInOrder);

	function getNeighbors(n) {
		let c = n.column;
		let r = n.row;
		let neighbors = [];
		let walls = [];
		if (validNode(r, c - 2)) {
			neighbors.push(grid[r][c - 2]);
			walls.push(grid[r][c - 1]);
		}
		if (validNode(r, c + 2)) {
			neighbors.push(grid[r][c + 2]);
			walls.push(grid[r][c + 1]);
		}
		if (validNode(r - 2, c)) {
			neighbors.push(grid[r - 2][c]);
			walls.push(grid[r - 1][c]);
		}
		if (validNode(r + 2, c)) {
			neighbors.push(grid[r + 2][c]);
			walls.push(grid[r + 1][c]);
		}
		return [neighbors, walls];
	}

	function validNode(row, column) {
		return (
			row > 0 &&
			row < grid.length - 1 &&
			column > 0 &&
			column < grid[0].length - 1 &&
			grid[row][column].visited == false
		);
	}
}
