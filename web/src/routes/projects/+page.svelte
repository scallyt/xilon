<script lang="ts">
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";

  interface Project {
    project_name: string;
    chat_id: number;
    customer_name: string;
    id: number
  }

  let projects: Project[] = [];

  onMount(async () => {
    try {
      const response = await fetch("http://localhost:3000/projects");
      if (!response.ok) {
        throw new Error(`Response status: ${response.status}`);
      }

      projects = await response.json();
    } catch (error) {
      console.error(error.message);
    }
  })
</script>

<main class="pj-boxes">
  {#each projects as project}
      <button type="button" on:click={() => goto(`/project/${project.id}`)} class="pj-box">
        <p>{project.project_name}</p>
        <a href="/chats/{project.chat_id}">Chat</a>
        <p>{project.customer_name}</p>
      </button>
    {/each}
</main>

<style>
  .pj-boxes {
    margin-top: 10rem;
    display: flex;
    flex-wrap: wrap;
    gap: 1rem;
  }

  .pj-box {
    border: 1px solid black;
    padding: 1rem;
    width: 10rem;
    height: 10rem;
  }
</style>
