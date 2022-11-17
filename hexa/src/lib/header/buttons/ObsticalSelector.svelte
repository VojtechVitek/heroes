<script>
	import GroupButton from './GroupButton.svelte';
	import Mountain from '$lib/icons/Mountain.svelte';
	import Wall from '$lib/icons/Wall.svelte';
	import Logs from '$lib/icons/Logs.svelte';
	import { activeObstacleStore, statusStore } from '$lib/board/stores.js';
	let status;
	function setStatus(type) {
		status = $statusStore;
		$statusStore = `Draw ${type}`;
	}
	function resetStatus() {
		$statusStore = status;
	}
</script>

<div>
	<GroupButton
		type="left"
		ariaLabel="Select wall"
		on:click={() => ($activeObstacleStore = 'wall')}
		on:mouseenter={() => setStatus('walls')}
		on:mouseleave={resetStatus}
		active={$activeObstacleStore == 'wall' ? 'buttonSelected' : ''}
		><Wall fill="var(--button-fg-color)" /></GroupButton
	>
	<GroupButton
		type="middle"
		ariaLabel="Select logs"
		on:click={() => ($activeObstacleStore = 'logs')}
		on:mouseenter={() => setStatus('logs')}
		on:mouseleave={resetStatus}
		active={$activeObstacleStore == 'logs' ? 'buttonSelected' : ''}
		><Logs fill="var(--button-fg-color)" /></GroupButton
	>
	<GroupButton
		type="right"
		ariaLabel="Select mountain"
		on:click={() => ($activeObstacleStore = 'mountain')}
		on:mouseenter={() => setStatus('mountains')}
		on:mouseleave={resetStatus}
		active={$activeObstacleStore == 'mountain' ? 'buttonSelected' : ''}
		><Mountain fill="var(--button-fg-color)" /></GroupButton
	>
</div>

<style>
	div {
		display: flex;
		gap: 0;
	}
</style>
