<script>
  export let comics = [];
  export let onSelectComic;
  
  function handleSelectComic(comic) {
    if (onSelectComic) {
      onSelectComic(comic);
    }
  }
</script>

<div class="library">
  {#if comics.length === 0}
    <div class="empty-library">
      <p>Nenhum arquivo CBZ ou ZIP encontrado na pasta selecionada.</p>
    </div>
  {:else}
    {#each comics as comic}
      <div class="comic-card" on:click={() => handleSelectComic(comic)}>
        <div class="comic-cover-container">
          {#if comic.coverPath}
            <img src={`file://${comic.coverPath}`} alt={comic.title} class="comic-cover" />
          {:else}
            <div class="comic-cover-placeholder">
              <span>Sem capa</span>
            </div>
          {/if}
        </div>
        <div class="comic-info">
          <h3 class="comic-title">{comic.title}</h3>
          <p class="comic-pages">{comic.pageCount} páginas</p>
        </div>
      </div>
    {/each}
  {/if}
</div>

<style>
  .empty-library {
    grid-column: 1 / -1;
    text-align: center;
    padding: 40px;
    color: #888;
  }
  
  .comic-cover-placeholder {
    width: 100%;
    aspect-ratio: 2/3;
    display: flex;
    justify-content: center;
    align-items: center;
    background-color: var(--secondary-color);
    color: #888;
    font-size: 14px;
  }
</style>