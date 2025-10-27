<script>
    import { fly } from 'svelte/transition';
    import { FolderOpen } from 'lucide-svelte';
    import { selectedDirectory, comicLibrary, selectDirectory } from './stores.js';
    import ComicCard from './ComicCard.svelte';
  </script>
  
  <div class="p-6" in:fly={{ y: 20, duration: 300 }}>
    {#if $selectedDirectory}
      <div class="mb-4">
        <h2 class="text-lg text-gray-400">Directory: {$selectedDirectory}</h2>
      </div>
      
      {#if $comicLibrary.length > 0}
        <div class="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-6 gap-6">
          {#each $comicLibrary as comic}
            <ComicCard {comic} />
          {/each}
        </div>
      {:else}
        <div class="text-center py-12">
          <p class="text-gray-400">No CBZ files found in the selected directory.</p>
        </div>
      {/if}
    {:else}
      <div class="text-center py-12">
        <p class="text-gray-400 mb-4">No directory selected. Please select a directory to scan for CBZ files.</p>
        <button 
          class="flex items-center space-x-2 bg-purple-800 hover:bg-purple-700 px-4 py-2 rounded-md transition-colors mx-auto"
          on:click={selectDirectory}
        >
          <FolderOpen size={16} />
          <span>Select Directory</span>
        </button>
      </div>
    {/if}
  </div>