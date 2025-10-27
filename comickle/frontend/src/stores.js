import { writable } from 'svelte/store';

// Library state
export const selectedDirectory = writable("");
export const comicLibrary = writable([]);
export const isLoading = writable(false);

// Reader state
export const isReaderView = writable(false);
export const currentComic = writable(null);
export const currentPage = writable(0);
export const totalPages = writable(0);


export function navigatePage(direction) {
  currentPage.update(page => {
    let newPage = page;
    let total;
    totalPages.subscribe(value => { total = value; })();
    
    if (direction === 'next' && page < total - 1) {
      newPage = page + 1;
    } else if (direction === 'prev' && page > 0) {
      newPage = page - 1;
    }
    
    return newPage;
  });
}

export function returnToLibrary() {
  isReaderView.set(false);
  currentComic.set(null);
  currentPage.set(0);
}


// In a real application, replace the mock functions with actual Wails calls:

export async function selectDirectory() {
    isLoading.set(true);
    try {
      // Call the Go function exposed by Wails
      const directory = await window.go.main.App.SelectDirectory();
      
      if (directory) {
        selectedDirectory.set(directory);
        const comics = await window.go.main.App.ScanDirectory(directory);
        comicLibrary.set(comics);
      }
    } catch (error) {
      console.error("Error selecting directory:", error);
      // You might want to add error handling here
    } finally {
      isLoading.set(false);
    }
  }
  
  export async function openComic(comic) {
    isLoading.set(true);
    try {
      // Call the Go function exposed by Wails
      const result = await window.go.main.App.OpenComic(comic.path);
      
      currentComic.set(comic);
      currentPage.set(0);
      totalPages.set(result.totalPages);
      isReaderView.set(true);
    } catch (error) {
      console.error("Error opening comic:", error);
      // You might want to add error handling here
    } finally {
      isLoading.set(false);
    }
  }
  
  // For getting a specific page, you would add:
  export async function getPage(comicPath, pageIndex) {
    isLoading.set(true);
    try {
      return await window.go.main.App.GetPage(comicPath, pageIndex);
    } catch (error) {
      console.error("Error getting page:", error);
      return null;
    } finally {
      isLoading.set(false);
    }
  }

