<script>
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function close() {
		dispatch('close');
	}
</script>

<div class="modal">
	<div class="backdrop" on:click={close} />
	<div class="wrapper">
		<h1>Welcome to AlgoPath!</h1>
		<div class="content">
			<p>
				AlgoPath is a visualizer for common maze solving and generating algorithms.
				<br />
				To start draw a maze or select one of the maze generating algorithms.
				<br />
				Then pick one of the solving algorithms to solve the maze
			</p>
			<!-- svelte-ignore a11y-media-has-caption -->
			<video src="/help-video.webm" autoplay loop>
				Sorry, your browser doesn't support embedded videos.
			</video>
		</div>
	</div>
</div>

<style>
	.modal {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100vh;
		display: flex;
		justify-content: center;
		align-items: center;
		z-index: 1;
	}
	.backdrop {
		position: absolute;
		width: 100%;
		height: 100%;
		background-color: rgba(0, 0, 0, 0.6);
	}
	.wrapper {
		width: 700px;
		height: 500px;
		border-radius: 0.2rem;
		background-color: white;
		z-index: 2;
		color: black;
		box-shadow: 4px 4px 4px rgba(0, 0, 0, 0.25);
	}
	.content {
		text-align: center;
	}
	h1 {
		text-align: center;
	}
</style>
