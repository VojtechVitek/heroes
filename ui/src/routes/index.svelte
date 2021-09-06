<script lang="ts">
	import { onMount } from 'svelte';
	import { api } from '../modules/api.ts';

	let vaults: Array<any> = [];

	let cdrTree: any = {
		name: 'loading...',
		files: []
	};

	onMount(async () => {
		try {
			const resp = await api.getVaults();
			vaults = resp.vaults;
		} catch (e) {
			console.error(e);
		}
	});

	onMount(async () => {
		try {
			const resp = await api.getCdrTree();
			cdrTree = resp.tree;
		} catch (e) {
			console.error(e);
		}
	});
</script>

<svelte:head>
	<title>Home</title>
</svelte:head>

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
