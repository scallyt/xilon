<script lang="ts">
  import { SignInButton, SignUpButton, SignOutButton, useClerkContext, SignedOut } from "svelte-clerk";

  const ctx = useClerkContext();
	const userId = $derived(ctx.auth.userId);
  import * as m from '$lib/paraglide/messages.js';

  
</script>

<nav class="navbar">
  <div class="container">
    <a href="/" class="brand">
      <span class="neko-text">Neko</span>
    </a>
    
    <div class="nav-links">
      {#if userId}
        <a href="/dashboard" class="nav-link">Dashboard</a>
        <a href="/chats" class="nav-link">Chats</a>
        <a href="/projects" class="nav-link">Projects</a>
        <div>
          <SignOutButton redirectUrl="/" />
        </div>
      {:else}
        <button class="auth-btn">
          <a href="/auth/register">{m.signUp()}</a>
        </button>
        <button class="auth-btn">
          <a href="/auth/login">{m.signIn()}</a>
        </button>
      {/if}
    </div>
  </div>
</nav>

<style>
  .navbar {
    background-color: #ffffff;
    box-shadow: 0 2px 8px rgba(0,0,0,0.1);
    padding: 1rem 0;
    position: fixed;
    width: 100%;
    top: 0;
    z-index: 1000;
  }

  .container {
    max-width: 1200px;
    margin: 0 auto;
    padding: 0 2rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .brand {
    text-decoration: none;
  }

  .neko-text {
    font-size: 1.8rem;
    font-weight: 700;
    background: linear-gradient(45deg, #FF6B6B, #4ECDC4);
    -webkit-background-clip: text;
    background-clip: text;
    -webkit-text-fill-color: transparent;
    text-decoration: none;
  }

  .nav-links {
    display: flex;
    gap: 1.5rem;
    align-items: center;
  }

  .nav-link {
    color: #333;
    text-decoration: none;
    font-weight: 500;
    transition: color 0.2s ease;
  }

  .nav-link:hover {
    color: #4ECDC4;
  }

  .auth-btn, .sign-out-btn {
    padding: 0.5rem 1.5rem;
    border-radius: 25px;
    border: none;
    font-weight: 600;
    cursor: pointer;
    transition: transform 0.2s ease;
  }

  .auth-btn {
    background: linear-gradient(45deg, #FF6B6B, #4ECDC4);
    color: white;
  }

  .sign-out-btn {
    background-color: #f8f9fa;
    color: #333;
    border: 1px solid #ddd;
  }

  .auth-btn:hover, .sign-out-btn:hover {
    transform: translateY(-2px);
  }
</style>