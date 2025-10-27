<script>
    import { Home, FolderOpen, Settings } from 'lucide-svelte';
    import { isReaderView, currentComic, currentPage, totalPages, returnToLibrary, selectDirectory } from './stores.js';
  </script>
  
  <header class="bg-gray-800 border-b border-purple-900 p-4 flex justify-between items-center">
    <div class="flex items-center space-x-2">
      <button 
        class="p-2 rounded-full hover:bg-purple-900 transition-colors"
        on:click={returnToLibrary}
        disabled={!$isReaderView}
      >
        <Home size={20} class={$isReaderView ? "text-white" : "text-gray-500"} />
      </button>
      <h1 class="text-xl font-bold">
        {#if $isReaderView}
          {$currentComic?.title} <span class="text-sm text-gray-400">({$currentPage + 1}/{$totalPages})</span>
        {:else}
          CBZ Reader
        {/if}
      </h1>
    </div>
    <div class="flex items-center space-x-2">
      {#if !$isReaderView}
        <button 
          class="flex items-center space-x-2 bg-purple-800 hover:bg-purple-700 px-4 py-2 rounded-md transition-colors"
          on:click={selectDirectory}
        >
          <FolderOpen size={16} />
          <span>Select Directory</span>
        </button>
      {/if}
      <button class="p-2 rounded-full hover:bg-purple-900 transition-colors">
        <Settings size={20} />
      </button>
    </div>
  </header>