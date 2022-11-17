import { get } from 'svelte/store';
import pathAnimation from '../../animations/pathfinders/pathAnimation.js';
import { endStore, gridStore, startStore } from '../../stores.js';
import { getNeighbors, setClearPath } from '../utils.js';

const distance = (x1, y1, x2, y2) => Math.hypot(x2 - x1, y2 - y1);

export default function astar() {
	let grid = get(gridStore);
	grid = setClearPath(grid);
	let startValue = get(startStore);
	let endValue = get(endStore);

	let visitedInOrder = [];
	let shortestPath = [];
	let unsolvable = false;

	let startNode = grid[startValue.row][startValue.column];
	let endNode = grid[endValue.row][endValue.column];

	let openSet = [];

	openSet.push(startNode);
	visitedInOrder.push(startNode);

	while (openSet.length) {
		let winner = openSet[0];
		for (let node of openSet) {
			if (node.f < winner.f) winner = node;
		}
		let current = winner;
		if (current == endNode) {
			while (current !== null) {
				shortestPath.unshift(current);
				current = current.previousNode;
			}
			break;
		}
		openSet = openSet.filter((n) => n != current);
		current.visited = true;
		let neighbors = getNeighbors(grid, current);
		for (let neighbor of neighbors) {
			visitedInOrder.push(neighbor);
			let tempG = current.distance + neighbor.obstacle;
			if (openSet.includes(neighbor)) {
				if (tempG < neighbor.distance) {
					neighbor.distance = tempG;
					neighbor.previousNode = current;
				}
			} else {
				neighbor.previousNode = current;
				neighbor.distance = tempG;
				openSet.push(neighbor);
			}
			neighbor.h = distance(neighbor.column, neighbor.row, endNode.column, endNode.row);
			neighbor.f = neighbor.distance + neighbor.h;
		}
	}
	pathAnimation(visitedInOrder, shortestPath, unsolvable);
}
