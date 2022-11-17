import { get } from 'svelte/store';
import { animationSpeedStore } from '../stores';

export async function wait(factor = 1) {
	let speed = get(animationSpeedStore);
	await new Promise((resolve) => setTimeout(resolve, speed * factor));
}
