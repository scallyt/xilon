<script lang="ts">
	import { i18n } from '$lib/i18n';
	import { ParaglideJS } from '@inlang/paraglide-sveltekit';
	import type { Snippet } from 'svelte';
	import { ClerkProvider } from 'svelte-clerk';
	import Navbar from '$lib/components/Navbar.svelte';

	const { children }: { children: Snippet } = $props();
</script>

<ClerkProvider>
	<ParaglideJS {i18n}>
		<Navbar />
		{@render children()}
	</ParaglideJS>
</ClerkProvider>