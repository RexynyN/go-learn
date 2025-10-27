<script>
  import { onMount } from 'svelte';
  import { GetPages, ExtractPage } from '../wailsjs/go/main/App';
  
  export let comic;
  export let onBack;
  
  let pages = [];
  let currentPageIndex = 0;
  let currentPagePath = '';
  let loading = false;
  let fullTitle = comic ? comic.title : '';
  
  onMount(async () => {
    if (comic) {
      await loadPages();
    }
  });
  
  async function loadPages() {
    try {
      loading = true;
      pages = await GetPages(comic.path);
      if (pages.length > 0) {
        await loadPage(0);
      }
    } catch (error) {
      console.error('Erro ao carregar páginas:', error);
      alert('Erro ao carregar páginas: ' + error.message);
    } finally {
      loading = false;
    }
  }
  
  async function loadPage(index) {
    if (!pages || index < 0 || index >= pages.length) {
      return;
    }
    
    try {
      loading = true;
      currentPageIndex = index;
      const pageName = pages[index];
      currentPagePath = await ExtractPage(comic.path, pageName);
    } catch (error) {
      console.error('Erro ao carregar página:', error);
      alert('Erro ao carregar página: ' + error.message);
    } finally {
      loading = false;
    }
  }
  
  function prevPage() {
    if (currentPageIndex > 0) {
      loadPage(currentPageIndex - 1);
    }
  }
  
  function nextPage() {
    if (currentPageIndex < pages.length - 1) {
      loadPage(currentPageIndex + 1);
    }
  }
  
  function handleKeydown(event) {
    if (event.key === 'ArrowLeft') {
      prevPage();
    } else if (event.key === 'ArrowRight') {
      nextPage();
    } else if (event.key === 'Escape') {
      if (onBack) onBack();
    }
  }
  
  function handleClick(event) {
    // Dividir a tela em duas partes para navegação
    const x = event.clientX;
    const width = window.innerWidth;
    
    if (x < width / 2) {
      prevPage();
    } else {
      nextPage();
    }
  }
</script>

<svelte:window on:keydown={handleKeydown} />

<div class="reader">
  <div class="reader-toolbar">
    <div class="reader-title">{fullTitle}</div>
    <div class="page-info">
      Página {currentPageIndex + 1} de {pages.length}
    </div>
  </div>
  
  <div class="reader-content" on:click={handleClick}>
    {#if loading && !currentPagePath}
      <div class="loading">Carregando página</div>
    {:else if currentPagePath}
      <img 
        src={`file://${currentPagePath}`} 
        alt={`Página ${currentPageIndex + 1}`} 
        class="reader-image" 
      />
    {/if}
  </div>
  
  <div class="reader-navigation">
    <button on:click={prevPage} disabled={currentPageIndex <= 0 || loading}>
      Anterior
    </button>
    <button on:click={nextPage} disabled={currentPageIndex >= pages.length - 1 || loading}>
      Próxima
    </button>
  </div>
</div>

<style>
  .reader-image {
    width: auto;
    height: auto;
    max-width: 100%;
    max-height: 100%;
    object-fit: contain;
  }
</style>