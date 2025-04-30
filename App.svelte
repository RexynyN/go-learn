<script>
  import { onMount, onDestroy } from 'svelte';
  import Library from './Library.svelte';
  import Reader from './Reader.svelte';
  import { SelectFolder, ScanFolder, CleanupTempFiles } from '../wailsjs/go/main/App';

  let folderPath = '';
  let comics = [];
  let loading = false;
  let selectedComic = null;
  let view = 'library'; // 'library' ou 'reader'

  onMount(async () => {
    // Verificar se temos um caminho de pasta salvo
    const savedPath = localStorage.getItem('folderPath');
    if (savedPath) {
      folderPath = savedPath;
      await loadComics();
    }
  });

  onDestroy(async () => {
    // Limpar arquivos temporários quando o componente for destruído
    await CleanupTempFiles();
  });

  async function selectFolder() {
    try {
      loading = true;
      const selected = await SelectFolder();
      if (selected) {
        folderPath = selected;
        localStorage.setItem('folderPath', folderPath);
        await loadComics();
      }
    } catch (error) {
      console.error('Erro ao selecionar pasta:', error);
      alert('Erro ao selecionar pasta: ' + error.message);
    } finally {
      loading = false;
    }
  }

  async function loadComics() {
    try {
      loading = true;
      comics = await ScanFolder(folderPath);
    } catch (error) {
      console.error('Erro ao carregar quadrinhos:', error);
      alert('Erro ao carregar quadrinhos: ' + error.message);
    } finally {
      loading = false;
    }
  }

  function openComic(comic) {
    selectedComic = comic;
    view = 'reader';
  }

  function backToLibrary() {
    selectedComic = null;
    view = 'library';
  }
</script>

<div class="app-container">
  <header class="header">
    <h1>CBZ Reader</h1>
    {#if view === 'library'}
      <button on:click={selectFolder} disabled={loading}>
        {loading ? 'Carregando...' : folderPath ? 'Mudar Pasta' : 'Selecionar Pasta'}
      </button>
    {:else}
      <button on:click={backToLibrary}>Voltar à Biblioteca</button>
    {/if}
  </header>

  <div class="content">
    {#if view === 'library'}
      {#if !folderPath}
        <div class="empty-state">
          <h2>Bem-vindo ao CBZ Reader</h2>
          <p>Selecione uma pasta contendo arquivos CBZ para começar.</p>
          <button on:click={selectFolder} disabled={loading}>
            {loading ? 'Carregando...' : 'Selecionar Pasta'}
          </button>
        </div>
      {:else if loading}
        <div class="loading">Carregando quadrinhos</div>
      {:else}
        <Library {comics} onSelectComic={openComic} />
      {/if}
    {:else}
      <Reader comic={selectedComic} onBack={backToLibrary} />
    {/if}
  </div>
</div>

<style>
  .app-container {
    height: 100%;
    display: flex;
    flex-direction: column;
  }
</style>