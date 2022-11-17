export function setAllWalls(grid) {
	grid = grid.map((row) =>
		row.map((node) => {
			node.type = 'wall';
			node.ini();
			return node;
		})
	);
	return grid;
}

export function setClearGrid(grid) {
	grid = grid.map((row) =>
		row.map((node) => {
			node.type = 'empty';
			node.ini();
			return node;
		})
	);
	return grid;
}

export function setClearPath(grid) {
	grid = grid.map((row) =>
		row.map((node) => {
			if (!['start', 'target', 'wall', 'mountain', 'logs'].includes(node.type)) {
				node.type = 'empty';
			}
			node.ini();
			return node;
		})
	);
	return grid;
}
export function randomBetween(start, end) {
	return Math.floor(Math.random() * (start - end + 1)) + end;
}

export function getNeighbors(grid, node) {
	let c = node.column;
	let r = node.row;
	let neighbors = [];
	let validNode = (row, column) => {
		return (
			row >= 0 &&
			row < grid.length &&
			column >= 0 &&
			column < grid[0].length &&
			grid[row][column].visited == false &&
			grid[row][column].type != 'wall'
		);
	};
	if (validNode(r, c - 1)) neighbors.push(grid[r][c - 1]);
	if (validNode(r, c + 1)) neighbors.push(grid[r][c + 1]);
	if (validNode(r - 1, c)) neighbors.push(grid[r - 1][c]);
	if (validNode(r + 1, c)) neighbors.push(grid[r + 1][c]);
	return neighbors;
}
