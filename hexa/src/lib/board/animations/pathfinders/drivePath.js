import { animationSpeedStore } from '$lib/board/stores';
import { get } from 'svelte/store';

export default function drivePath(nodes) {
	const start = document.getElementById('start');
	let road = path(nodes);
	//current implementation do not allow to change speed during animation
	let time = nodes.length * 15 * get(animationSpeedStore);
	let oldStyle = start.style.cssText;
	start.style.cssText += `offset-path: path('${road}'); animation: move ${time}ms forwards linear;`;
	start.onanimationend = () => {
		start.style.cssText = oldStyle;
	};
}

function path(nodes) {
	let row = nodes[0].row;
	let column = nodes[0].column;
	nodes.shift();
	let arr = [];
	let hor = 15;
	let vert = 15;
	for (let node of nodes) {
		hor += (node.row - row) * 30;
		vert += (node.column - column) * 30;
		arr.push(`${hor},${vert}`);
		row = node.row;
		column = node.column;
	}
	return `M15,15 ${arr.join(' ')}`;
}
