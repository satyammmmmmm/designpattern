package main

import "fmt"

// Builder interface defines methods for building a package
type Builder interface {
	SetDependency(dependency []string)                // Method for setting dependencies
	SetConfiguration(configuration map[string]string) // Method for setting configurations
	SetInstaller(installer string)                    // Method for setting installer type
	GetPackage() Package                              // Method for getting the constructed package
}

// Package struct represents a software package
type Package struct {
	Dependency    []string          // Slice of dependencies required by the package
	Configuration map[string]string // Map representing configuration parameters of the package
	Installer     string            // Type of installer for the package (e.g., "msi", "deb")
}

// WindowInstaller struct represents a builder for Windows installation packages
type WindowInstaller struct {
	packageInfo Package // Holds package information being constructed
}

// Implementation of Builder interface methods for WindowInstaller
func (w *WindowInstaller) SetDependency(dependency []string) {
	w.packageInfo.Dependency = dependency
}
func (w *WindowInstaller) SetConfiguration(configuration map[string]string) {
	w.packageInfo.Configuration = configuration
}
func (w *WindowInstaller) SetInstaller(installer string) {
	w.packageInfo.Installer = installer
}
func (w *WindowInstaller) GetPackage() Package {
	return w.packageInfo
}

// LinuxInstaller struct represents a builder for Linux installation packages
type LinuxInstaller struct {
	packageInfo Package // Holds package information being constructed
}

// Implementation of Builder interface methods for LinuxInstaller
func (l *LinuxInstaller) SetDependency(dependency []string) {
	l.packageInfo.Dependency = dependency
}
func (l *LinuxInstaller) SetConfiguration(configuration map[string]string) {
	l.packageInfo.Configuration = configuration
}
func (l *LinuxInstaller) SetInstaller(installer string) {
	l.packageInfo.Installer = installer
}
func (l *LinuxInstaller) GetPackage() Package {
	return l.packageInfo
}

// ConstructOsPackage is a function to construct an OS installation package using a given builder
func ConstructOsPackage(build Builder, dependency []string, configuration map[string]string, installer string) Package {
	build.SetDependency(dependency)
	build.SetConfiguration(configuration)
	build.SetInstaller(installer)
	return build.GetPackage()
}

func main() {
	// Create a Windows installer builder
	windowsBuilder := &WindowInstaller{}

	// Construct a Windows installation package
	windowsPackage := ConstructOsPackage(
		windowsBuilder,
		[]string{"dependency1", "dependency2"},
		map[string]string{"config1": "value1", "config2": "value2"},
		"msi")

	// Print details of Windows installation package
	fmt.Println("Windows Installation Package:")
	fmt.Println("Dependencies:", windowsPackage.Dependency)
	fmt.Println("Configurations:", windowsPackage.Configuration)
	fmt.Println("Installer:", windowsPackage.Installer)

	// Create a Linux installer builder
	linuxBuilder := &LinuxInstaller{}

	// Construct a Linux installation package
	linuxPackage := ConstructOsPackage(
		linuxBuilder,
		[]string{"dependency3", "dependency4"},
		map[string]string{"config3": "value3", "config4": "value4"},
		"deb")

	// Print details of Linux installation package
	fmt.Println("\nLinux Installation Package:")
	fmt.Println("Dependencies:", linuxPackage.Dependency)
	fmt.Println("Configurations:", linuxPackage.Configuration)
	fmt.Println("Installer:", linuxPackage.Installer)
}
