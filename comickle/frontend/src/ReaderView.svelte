<script>
    import { fly } from 'svelte/transition';
    import { ChevronLeft, ChevronRight } from 'lucide-svelte';
    import { currentPage, totalPages, navigatePage } from './stores.js';
  </script>
  
  <div 
    class="h-full flex flex-col items-center justify-center relative bg-black"
    in:fly={{ y: 20, duration: 300 }}
  >
    <img 
      src={`/placeholder.svg?height=800&width=600&text=Page ${$currentPage + 1}`} 
      alt={`Page ${$currentPage + 1}`}
      class="max-h-[calc(100vh-8rem)] w-auto object-contain"
    />
    
    <!-- Navigation Controls -->
    <div class="absolute inset-x-0 bottom-4 flex justify-center space-x-4">
      <button 
        class="p-3 rounded-full bg-gray-800 bg-opacity-70 hover:bg-purple-900 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
        on:click={() => navigatePage('prev')}
        disabled={$currentPage === 0}
      >
        <ChevronLeft size={24} />
      </button>
      <button 
        class="p-3 rounded-full bg-gray-800 bg-opacity-70 hover:bg-purple-900 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
        on:click={() => navigatePage('next')}
        disabled={$currentPage === $totalPages - 1}
      >
        <ChevronRight size={24} />
      </button>
    </div>
    
    <!-- Page indicator -->
    <div class="absolute bottom-4 right-4 bg-gray-800 bg-opacity-70 px-3 py-1 rounded-md text-sm">
      {$currentPage + 1} / {$totalPages}
    </div>
  </div>