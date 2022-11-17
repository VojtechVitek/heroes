<script lang="ts">
	import Board from '$lib/board/Board.svelte';
	import Header from '$lib/header/Header.svelte';
	import './app.css';
</script>

<svelte:head>
	<title>Map</title>
	<meta name="description" content="Map" />
</svelte:head>

<Header />
<div class="text-column">
	<Board />
</div>
