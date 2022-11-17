import { setClearPath } from '$lib/board/algorithms/utils';

import { gridStore, shortestPathStore, statusStore } from '$lib/board/stores';
import { get } from 'svelte/store';

export default async function pathAnimation(checkedNodes, shortestPath, unsolvable) {
	let grid = get(gridStore);
	statusStore.set('inProgress');
	gridStore.set(setClearPath(grid));
	for (let node of checkedNodes) {
		node.classes = 'checked';
		gridStore.forceUpdate();
	}
	if (unsolvable) {
		statusStore.set('unsolvable');
		return;
	}
	for (let node of shortestPath) {
		node.classes = 'shortestPath';
		gridStore.forceUpdate();
	}
	statusStore.set(`solved in ${shortestPath.length}`);
	shortestPathStore.set(shortestPath);
}
