<script lang="ts">
  import { onMount } from "svelte";
  import { useClerkContext } from "svelte-clerk";

  	onMount(async () => {
		const ctx = useClerkContext();
		if (ctx.session) {
			const token = await ctx.session.getToken();

			console.log(token)
		}
	});
</script>

<h1>Dashboard</h1>