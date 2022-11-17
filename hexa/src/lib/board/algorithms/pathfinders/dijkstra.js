import { get } from 'svelte/store';
import pathAnimation from '../../animations/pathfinders/pathAnimation.js';
import { gridStore } from '../../stores.js';
import { getNeighbors, setClearPath } from '../utils.js';

export default function dijkstra() {
	let grid = get(gridStore);
	grid = setClearPath(grid);

	let unvisited = grid.flat();
	let current;
	let visitedInOrder = [];
	let shortestPath = [];
	let unsolvable = false;

	while (unvisited.length !== 0) {
		unvisited.sort((a, b) => a.distance - b.distance);
		if (unvisited[0].distance == Infinity) {
			unsolvable = true;
			break;
		}
		current = unvisited.shift();
		if (current.type == 'wall') continue;
		visitedInOrder.push(current);
		let neighbors = getNeighbors(grid, current);
		neighbors.forEach((neighbor) => {
			let distance = current.distance + neighbor.obstacle;
			if (neighbor.distance > distance) {
				neighbor.distance = distance;
				neighbor.previousNode = current;
			}
			visitedInOrder.push(neighbor);
		});
		current.visited = true;
		if (current.type == 'target') {
			while (current !== null) {
				shortestPath.unshift(current);
				current = current.previousNode;
			}
			break;
		}
	}
	pathAnimation(visitedInOrder, shortestPath, unsolvable);
}
