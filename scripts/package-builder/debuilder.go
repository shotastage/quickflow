package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Define package details
	packageName := "quickflow-server"
	version := "1.0.0"
	architecture := "amd64"
	maintainer := "Shota Shimazu <shota.shimazu@magique.app>"
	description := "High-performance CMS and API platform written in Go"

	// Create package structure
	os.MkdirAll("debian/DEBIAN", 0755)
	os.MkdirAll("debian/usr/bin", 0755)

	// Create control file
	control := fmt.Sprintf(`Package: %s
Version: %s
Section: web
Priority: optional
Architecture: %s
Maintainer: %s
Description: %s
`, packageName, version, architecture, maintainer, description)

	err := os.WriteFile("debian/DEBIAN/control", []byte(control), 0644)
	if err != nil {
		fmt.Println("Error writing control file:", err)
		return
	}

	// Build the binary
	cmd := exec.Command("go", "build", "-o", "debian/usr/bin/quickflow")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error building binary:", err)
		return
	}

	// Build the deb package
	cmd = exec.Command("dpkg-deb", "--build", "debian")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error building deb package:", err)
		return
	}

	fmt.Println("Deb package built successfully")
}
