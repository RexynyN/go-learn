<script>
    import { openComic } from './stores.js';
    
    export let comic;
  </script>
  
  <div 
    class="bg-gray-800 rounded-lg overflow-hidden hover:ring-2 hover:ring-purple-600 transition-all cursor-pointer transform hover:scale-105"
    on:click={() => openComic(comic)}
  >
    <div class="aspect-[2/3] relative">
      <img 
        src={comic.coverImage || "/placeholder.svg"} 
        alt={comic.title}
        class="w-full h-full object-cover"
      />
    </div>
    <div class="p-3">
      <h3 class="text-sm font-medium truncate">{comic.title}</h3>
    </div>
  </div>