<script>
	import { createEventDispatcher, onMount } from 'svelte';
	import { scale } from 'svelte/transition';
	import Arrow from '$lib/icons/Arrow.svelte';
	import { statusStore } from '$lib/board/stores';

	export let dropdownList;
	let style = $$props.style;

	const dispatch = createEventDispatcher();
	let show = false;
	let dropdown = null;

	onMount(() => {
		const handleOutsideClick = (event) => {
			if (show && !dropdown.contains(event.target)) {
				show = false;
			}
		};
		document.addEventListener('click', handleOutsideClick, false);

		return () => {
			document.removeEventListener('click', handleOutsideClick, false);
		};
	});

	function changeAlgoritm(algoritm) {
		dispatch('algoritm', {
			text: algoritm
		});
	}
	let status;
	function setStatus() {
		status = $statusStore;
		$statusStore = 'Change algorithm';
	}
	function resetStatus() {
		$statusStore = status;
	}
</script>

<div class="split-button" {style} bind:this={dropdown}>
	<button on:click on:mouseenter on:mouseleave>
		<slot />
	</button>
	<span
		class="dropdown-button"
		on:click={() => (show = !show)}
		on:mouseenter={setStatus}
		on:mouseleave={resetStatus}
	>
		<Arrow fill="var(--button-fg-color)" />
	</span>
	{#if show}
		<div
			class="dropdown-content"
			in:scale={{ duration: 100, start: 0.95 }}
			out:scale={{ duration: 75, start: 0.95 }}
		>
			<ul>
				{#each dropdownList as item}
					<li
						on:click={() => {
							show = !show;
							changeAlgoritm(item);
						}}
					>
						{item}
					</li>
				{/each}
			</ul>
		</div>
	{/if}
</div>

<style>
	.split-button {
		background-color: var(--button-bg-color);
		border-radius: var(--button-border-radius);
		cursor: pointer;
		position: relative;
		display: flex;
		justify-content: space-between;
		height: var(--button-height);
	}
	button {
		border-radius: var(--button-border-radius) 0 0 var(--button-border-radius);
		background-color: var(--button-bg-color);
		color: var(--button-fg-color);
		width: 100%;
		text-align: start;
		border: none;
		cursor: pointer;
	}
	.dropdown-button {
		display: flex;
		align-items: center;
		justify-content: end;
		padding-left: 8px;
		padding-right: 8px;
		border-left: var(--button-separator) solid 2px;
		border-radius: 0 var(--button-border-radius) var(--button-border-radius) 0;
	}
	button:hover,
	.dropdown-button:hover {
		background-color: var(--button-hovered);
	}
	.dropdown-content {
		transform: translateY(var(--button-height));
		position: absolute;
		background-color: white;
		color: black;
		width: inherit;
		z-index: 1;
		box-shadow: rgba(100, 100, 111, 0.2) 0px 7px 29px 0px;
	}
	ul {
		padding: 0;
		width: inherit;
	}
	li {
		list-style-type: none;
		padding: 8px;
		padding-left: 30px;
	}
	li:hover {
		background-color: rgb(138, 138, 138);
	}
</style>
