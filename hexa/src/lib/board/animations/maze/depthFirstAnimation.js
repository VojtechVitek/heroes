import { setAllWalls } from '$lib/board/algorithms/utils';
import { gridStore, statusStore } from '$lib/board/stores';
import { get } from 'svelte/store';
import { wait } from '../utils';
export default async function depthFirstAnimation(nodes) {
	statusStore.set('inProgress');
	gridStore.set(setAllWalls(get(gridStore)));
	for (let node of nodes) {
		await wait();
		node.type = 'digger';
		gridStore.forceUpdate();
		await wait();
		node.type = 'empty';
		gridStore.forceUpdate();
	}
	nodes[Math.floor(Math.random() * nodes.length)].setType('start');
	nodes[Math.floor(Math.random() * nodes.length)].setType('target');
	statusStore.set('done');
}
