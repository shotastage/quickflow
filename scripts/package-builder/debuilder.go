package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Define package details
	packageName := "quickflow-server"
	version := "1.0.0"
	architecture := "amd64"

	// Define paths
	resourceDir := "resources/debian"
	debianDir := "debian"
	binDir := filepath.Join(debianDir, "usr/bin")
	binSrc := "bin/quickflow"

	// Create package structure
	err := os.MkdirAll(binDir, 0755)
	if err != nil {
		fmt.Println("Error creating package structure:", err)
		return
	}

	// Copy control files
	err = copyDir(resourceDir, debianDir)
	if err != nil {
		fmt.Println("Error copying control files:", err)
		return
	}

	// Build the binary using make
	cmd := exec.Command("make", "build")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error building binary:", err)
		return
	}

	// Copy the built binary to the package structure
	err = copyFile(binSrc, filepath.Join(binDir, "quickflow"))
	if err != nil {
		fmt.Println("Error copying binary:", err)
		return
	}

	// Set permissions for the binary
	err = os.Chmod(filepath.Join(binDir, "quickflow"), 0755)
	if err != nil {
		fmt.Println("Error setting permissions for binary:", err)
		return
	}

	// Build the deb package
	cmd = exec.Command("dpkg-deb", "--build", debianDir)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error building deb package:", err)
		return
	}

	// Move the deb package to the current directory
	err = os.Rename(debianDir+".deb", fmt.Sprintf("%s_%s_%s.deb", packageName, version, architecture))
	if err != nil {
		fmt.Println("Error renaming deb package:", err)
		return
	}

	// Clean up
	err = os.RemoveAll(debianDir)
	if err != nil {
		fmt.Println("Error cleaning up:", err)
		return
	}

	fmt.Println("Deb package built successfully")
}

// copyDir copies a whole directory recursively
func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		return os.WriteFile(dstPath, data, info.Mode())
	})
}

// copyFile copies a single file from src to dst
func copyFile(src, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	return os.WriteFile(dst, data, 0755)
}
