<script>
	export let type;
	export let active;
	export let ariaLabel;
</script>

<button aria-label={ariaLabel} on:mouseenter on:mouseleave on:click class={`${type} ${active}`}
	><slot /></button
>

<style>
	button {
		color: var(--button-fg-color);
		background-color: var(--button-bg-color);
		height: var(--button-height);
		border: none;
		cursor: pointer;
	}
	.left {
		border-radius: var(--button-border-radius) 0 0 var(--button-border-radius);
		border-right: var(--button-separator) solid 2px;
	}
	.right {
		border-radius: 0 var(--button-border-radius) var(--button-border-radius) 0;
		border-left: var(--button-separator) solid 2px;
	}
	.buttonSelected {
		background-color: var(--button-selected);
	}
	:hover {
		background-color: var(--button-hovered);
	}
</style>
