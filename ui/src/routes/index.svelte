<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '../modules/api';
	import type { Map } from '../modules/rpc.gen';

	let maps: string[] = [];
	let map: Map = { Tiles: [] };

	const tileTypeColor = {
		0: '#0F3F50', // Dirt
		1: '#8FCFDF', // Sand
		2: '#004000', // Grass
		3: '#C0C0B0', // Snow
		4: '#6F804F', // Swamp
		5: '#307080', // Rough
		6: '#308000', // Subterranean
		7: '#4F4F4F', // Lava
		8: '#90500F', // Water
		9: '#000000' // Rock
	};

	onMount(async () => {
		try {
			const resp = await api.listMaps();
			maps = resp.maps;
		} catch (e) {
			console.error(e);
		}
	});
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>

<select>
	{#each maps as name}
		<option value={name}>{name}</option>
	{/each}
</select>

<table>
	{#each Array(map.MapSize) as _, i}
		<tr style="margin: 0px; padding: 0px;">
			{#each map.Tiles.slice(i * map.MapSize, (i + 1) * map.MapSize) as tile}
				<td
					style="background-color: {tileTypeColor[tile.TerrainType]}; margin: 0px; padding: 5px;"
				/>
			{/each}
		</tr>
	{/each}
</table>

<section />

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 1;
	}

	td {
		vertical-align: top;
		text-align: left;
	}

	.left {
		width: 50%;
		height: 1000px;
		width: 800px;
	}

	.right {
		width: 50%;
		height: 1000px;
		width: 400px;
	}
</style>
