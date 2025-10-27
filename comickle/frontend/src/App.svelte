<script>
  import { onMount } from 'svelte';
  import { isReaderView, isLoading, currentPage, navigatePage, returnToLibrary } from './stores.js';
  import { get } from 'svelte/store';
  
  // Import components
  import Header from './Header.svelte';
  import Footer from './Footer.svelte';
  import LibraryView from './LibraryView.svelte';
  import ReaderView from './ReaderView.svelte';
  import LoadingSpinner from './LoadingSpinner.svelte';

  // Keyboard navigation
  function handleKeydown(event) {
    if (!get(isReaderView)) return;
    
    if (event.key === 'ArrowRight') {
      navigatePage('next');
    } else if (event.key === 'ArrowLeft') {
      navigatePage('prev');
    } else if (event.key === 'Escape') {
      returnToLibrary();
    }
  }

  onMount(() => {
    window.addEventListener('keydown', handleKeydown);
    return () => {
      window.removeEventListener('keydown', handleKeydown);
    };
  });
</script>

<svelte:head>
  <title>CBZ Reader</title>
</svelte:head>

<div class="min-h-screen bg-gray-900 text-white flex flex-col">
  <!-- Header -->
  <Header />

  <!-- Main Content -->
  <main class="flex-1 overflow-hidden">
    {#if $isLoading}
      <LoadingSpinner />
    {:else if get(isReaderView)}
      <ReaderView />
    {:else}
      <LibraryView />
    {/if}
  </main>

  <!-- Footer -->
  <Footer />
</div>

<style>
  :global(body) {
    margin: 0;
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen,
      Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
    background-color: #111827; /* bg-gray-900 */
    color: white;
  }
</style>