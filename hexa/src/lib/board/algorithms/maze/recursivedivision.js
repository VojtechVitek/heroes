import mazeAnimation from '$lib/board/animations/maze/mazeAnimation.js';
import { get } from 'svelte/store';
import { gridStore } from '../../stores.js';
import { setClearGrid } from '../utils.js';

export default function recursiveDivision() {
	let grid = get(gridStore);
	grid = setClearGrid(grid);
	let visitedInOrder = [];
	let columns = grid.length;
	let rows = grid[0].length;

	if (!(rows % 2)) {
		for (let c = 0; c < columns; c++) visitedInOrder.push(grid[c][rows - 1]);
		rows--;
	}
	if (!(columns % 2)) {
		for (let r = rows - 1; r >= 0; r--) visitedInOrder.push(grid[columns - 1][r]);
		columns--;
	}
	(function addOuterWalls() {
		for (let r = 0; r < rows; r++) visitedInOrder.push(grid[0][r]);
		for (let c = 0; c < columns; c++) visitedInOrder.push(grid[c][rows - 1]);
		for (let r = rows - 1; r >= 0; r--) visitedInOrder.push(grid[columns - 1][r]);
		for (let c = columns - 1; c >= 0; c--) visitedInOrder.push(grid[c][0]);
	})();
	division(1, columns - 1, 1, rows - 1);
	function division(columnStart, columnEnd, rowStart, rowEnd) {
		let columns = Math.abs(columnStart - columnEnd);
		let rows = Math.abs(rowStart - rowEnd);
		let divideColumns = getDivideColumns(columns, rows);
		if ((divideColumns && rows < 3) || (!divideColumns && columns < 3)) return;
		let wall = divideColumns
			? randomEvenBetween(columnStart, columnEnd)
			: randomEvenBetween(rowStart, rowEnd);
		let open = divideColumns
			? randomOddBetween(rowStart, rowEnd)
			: randomOddBetween(columnStart, columnEnd);
		if (divideColumns) {
			for (let r = rowStart; r < rowEnd; r++) {
				if (r != open) visitedInOrder.push(grid[wall][r]);
			}
			division(wall, columnEnd, rowStart, rowEnd);
			division(columnStart, wall, rowStart, rowEnd);
		} else {
			for (let c = columnStart; c < columnEnd; c++) {
				if (c != open) visitedInOrder.push(grid[c][wall]);
			}
			division(columnStart, columnEnd, rowStart, wall);
			division(columnStart, columnEnd, wall, rowEnd);
		}
	}
	mazeAnimation(visitedInOrder);
}

function getDivideColumns(columns, rows) {
	if (rows > columns) return false;
	else if (rows < columns) return true;
	else if (rows == columns) return Math.random() < 0.5;
}

function randomOddBetween(start, end) {
	let random = Math.floor(Math.random() * (start - end + 1)) + end;
	if (!(random % 2)) return randomOddBetween(start, end);
	return random;
}
function randomEvenBetween(start, end) {
	let random = Math.floor(Math.random() * (start - end + 1)) + end;
	if (random % 2) return randomEvenBetween(start, end);
	return random;
}
