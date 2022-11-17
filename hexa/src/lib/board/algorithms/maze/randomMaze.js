import mazeAnimation from '$lib/board/animations/maze/mazeAnimation.js';
import { get } from 'svelte/store';
import { gridStore } from '../../stores.js';

export default function randomMaze() {
	let grid = get(gridStore);
	let visitedInOrder = [];
	grid.forEach((row) =>
		row.forEach((node) => {
			if (Math.random() > 0.7 && !['start', 'target'].includes(node.type))
				visitedInOrder.push(node);
		})
	);
	mazeAnimation(visitedInOrder);
}
