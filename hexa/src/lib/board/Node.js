import { startStore, endStore } from './stores.js';

export default class Node {
	constructor(row, column, type = 'empty') {
		this.row = row;
		this.column = column;
		this.setType(type);
		this.ini();
	}
	setType(type) {
		this.type = type;
		this.obstacle = 1;
		if (type == 'start') {
			this.distance = 0;
			startStore.set({ row: this.row, column: this.column });
		}
		if (type == 'target') {
			endStore.set({ row: this.row, column: this.column });
		}
		if (type == 'logs') this.obstacle = 2;
		if (type == 'mountain') this.obstacle = 3;
	}
	ini() {
		this.classes = '';
		this.h = Infinity;
		this.f = Infinity;
		this.distance = this.type == 'start' ? 0 : Infinity;
		this.visited = false;
		this.previousNode = null;
	}
}
