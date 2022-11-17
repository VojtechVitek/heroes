<button on:click><slot /></button>

<style>
	button {
		color: var(--button-fg-color);
		background-color: var(--button-bg-color);
		border-radius: 50%;
		height: var(--button-height);
		aspect-ratio: 1;
		border: var(--button-border);
		cursor: pointer;
	}
	:hover {
		background-color: var(--button-hovered);
	}
</style>
