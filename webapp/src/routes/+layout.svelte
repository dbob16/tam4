<script>
	import favicon from '$lib/assets/favicon.svg';
	import '$lib/assets/css/main.css'
	import '$lib/assets/css/schemeofcolors.css'
	import { browser } from '$app/environment';

	let { children, data } = $props();

	const copyTheme = () => {return data.settings.display_theme}

	const displayTheme = $state(copyTheme())

	if (browser) {
	  document.body.classList.add(displayTheme)
	}
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

{@render children()}
