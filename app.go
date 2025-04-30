package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp cria uma nova estrutura de aplicação
func NewApp() *App {
	return &App{}
}

// startup é chamado quando o app inicia
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// ComicBook representa um arquivo CBZ/ZIP
type ComicBook struct {
	Path      string `json:"path"`
	Title     string `json:"title"`
	CoverPath string `json:"coverPath"`
	PageCount int    `json:"pageCount"`
}

// SelectFolder abre um diálogo para selecionar uma pasta e retorna o caminho
func (a *App) SelectFolder() (string, error) {
	selectedFolder, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Selecione uma pasta com arquivos CBZ",
	})
	
	if err != nil {
		return "", err
	}
	
	return selectedFolder, nil
}

// ScanFolder escaneia uma pasta em busca de arquivos CBZ/ZIP e retorna informações sobre eles
func (a *App) ScanFolder(folderPath string) ([]ComicBook, error) {
	var comics []ComicBook
	
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		if !info.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".cbz" || ext == ".zip" {
				// Extrair a primeira página para usar como capa
				coverPath, pageCount, err := a.extractCover(path)
				if err != nil {
					// Se não conseguir extrair a capa, ainda adiciona o comic, mas sem capa
					comics = append(comics, ComicBook{
						Path:      path,
						Title:     filepath.Base(path),
						CoverPath: "",
						PageCount: 0,
					})
				} else {
					comics = append(comics, ComicBook{
						Path:      path,
						Title:     filepath.Base(path),
						CoverPath: coverPath,
						PageCount: pageCount,
					})
				}
			}
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// Ordenar por título
	sort.Slice(comics, func(i, j int) bool {
		return comics[i].Title < comics[j].Title
	})
	
	return comics, nil
}

// extractCover extrai a primeira página de um arquivo CBZ/ZIP para usar como capa
func (a *App) extractCover(cbzPath string) (string, int, error) {
	reader, err := zip.OpenReader(cbzPath)
	if err != nil {
		return "", 0, err
	}
	defer reader.Close()
	
	// Encontrar arquivos de imagem
	var imageFiles []string
	for _, file := range reader.File {
		lowerName := strings.ToLower(file.Name)
		if strings.HasSuffix(lowerName, ".jpg") || 
		   strings.HasSuffix(lowerName, ".jpeg") || 
		   strings.HasSuffix(lowerName, ".png") || 
		   strings.HasSuffix(lowerName, ".gif") || 
		   strings.HasSuffix(lowerName, ".webp") {
			imageFiles = append(imageFiles, file.Name)
		}
	}
	
	if len(imageFiles) == 0 {
		return "", 0, fmt.Errorf("nenhuma imagem encontrada no arquivo")
	}
	
	// Ordenar arquivos de imagem
	sort.Strings(imageFiles)
	
	// Extrair a primeira imagem como capa
	coverFileName := imageFiles[0]
	
	// Encontrar o arquivo no zip
	var coverFile *zip.File
	for _, f := range reader.File {
		if f.Name == coverFileName {
			coverFile = f
			break
		}
	}
	
	if coverFile == nil {
		return "", len(imageFiles), fmt.Errorf("arquivo de capa não encontrado")
	}
	
	// Criar diretório temporário se não existir
	tempDir := filepath.Join(os.TempDir(), "cbzreader", "covers")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", len(imageFiles), err
	}
	
	// Criar um nome de arquivo único para a capa extraída
	safeTitle := strings.ReplaceAll(filepath.Base(cbzPath), " ", "_")
	coverPath := filepath.Join(tempDir, safeTitle+".jpg")
	
	// Verificar se o arquivo já existe
	if _, err := os.Stat(coverPath); err == nil {
		return coverPath, len(imageFiles), nil
	}
	
	// Extrair o arquivo
	src, err := coverFile.Open()
	if err != nil {
		return "", len(imageFiles), err
	}
	defer src.Close()
	
	dst, err := os.Create(coverPath)
	if err != nil {
		return "", len(imageFiles), err
	}
	defer dst.Close()
	
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", len(imageFiles), err
	}
	
	return coverPath, len(imageFiles), nil
}

// GetPages retorna uma lista de páginas de um arquivo CBZ/ZIP
func (a *App) GetPages(cbzPath string) ([]string, error) {
	reader, err := zip.OpenReader(cbzPath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	
	// Encontrar arquivos de imagem
	var imageFiles []string
	for _, file := range reader.File {
		lowerName := strings.ToLower(file.Name)
		if strings.HasSuffix(lowerName, ".jpg") || 
		   strings.HasSuffix(lowerName, ".jpeg") || 
		   strings.HasSuffix(lowerName, ".png") || 
		   strings.HasSuffix(lowerName, ".gif") || 
		   strings.HasSuffix(lowerName, ".webp") {
			imageFiles = append(imageFiles, file.Name)
		}
	}
	
	// Ordenar arquivos de imagem
	sort.Strings(imageFiles)
	
	return imageFiles, nil
}

// ExtractPage extrai uma página de um arquivo CBZ/ZIP para um local temporário e retorna o caminho
func (a *App) ExtractPage(cbzPath string, pageName string) (string, error) {
	reader, err := zip.OpenReader(cbzPath)
	if err != nil {
		return "", err
	}
	defer reader.Close()
	
	// Criar diretório temporário se não existir
	tempDir := filepath.Join(os.TempDir(), "cbzreader", "pages")
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", err
	}
	
	// Criar um nome de arquivo único para a página extraída
	safeTitle := strings.ReplaceAll(filepath.Base(cbzPath), " ", "_")
	safePage := strings.ReplaceAll(filepath.Base(pageName), " ", "_")
	outputPath := filepath.Join(tempDir, safeTitle+"_"+safePage)
	
	// Verificar se o arquivo já existe
	if _, err := os.Stat(outputPath); err == nil {
		return outputPath, nil
	}
	
	// Encontrar o arquivo no zip
	var file *zip.File
	for _, f := range reader.File {
		if f.Name == pageName {
			file = f
			break
		}
	}
	
	if file == nil {
		return "", fmt.Errorf("página não encontrada no arquivo: %s", pageName)
	}
	
	// Extrair o arquivo
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	
	dst, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	
	return outputPath, nil
}

// CleanupTempFiles remove arquivos temporários
func (a *App) CleanupTempFiles() error {
	tempDir := filepath.Join(os.TempDir(), "cbzreader")
	return os.RemoveAll(tempDir)
}