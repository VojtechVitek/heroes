<script>
	import { statusStore } from '$lib/board/stores.js';
</script>

<div>
	{#if $statusStore == ''}
		Draw a maze or generate one with one of the maze algoritms
	{:else if $statusStore == 'solved'}
		Maze solved
	{:else if $statusStore == 'unsolvable'}
		<span class="red">No path</span>
	{:else if $statusStore == 'inProgress'}
		Animation in progress
	{:else}
		{$statusStore}
	{/if}
</div>

<style>
	div {
		display: flex;
		align-items: center;
		justify-content: center;
		height: 30px;
	}
	.red {
		color: red;
	}
</style>
