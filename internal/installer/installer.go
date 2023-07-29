package installer

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/thomasoca/cv-generator/pkg/image"
	"github.com/thomasoca/cv-generator/pkg/utils"
)

func downloadFile(url, filePath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func latexClassMod(class string) error {
	switch class {
	case "altacv.cls":
		err := utils.AltaCvMod()
		return err
	}
	return nil
}

func InstallPrerequisite() error {
	// Check if latex class is already in the machine
	listOfClasses := GetLatexClass()

	fmt.Println("Checking Latex class availability....")
	for _, m := range listOfClasses {
		class := m["class"]
		url := m["url"]
		if !image.IsImageFileExist(class) {
			// Download latex class from the author Git repository
			if err := downloadFile(url, class); err != nil {
				fmt.Println("Failed to download latex class:", err)
				return err
			}
			// Apply mod to the Latex class
			if err := latexClassMod(class); err != nil {
				fmt.Println("Failed to apply mod to the Latex class:", err)
				return err
			}
			fmt.Printf("%s downloaded successfully.\n", class)
		}
	}
	fmt.Println("All of the Latex classes are available in the system")

	// Check if pdflatex is installed
	fmt.Println("Checking Latex executables....")
	if err := utils.RunCommand("pdflatex", nil, nil, "--version"); err != nil {
		if runtime.GOOS == "linux" {
			fmt.Println("pdflatex is not installed, starting TinyTex installation process....")

			// Download TinyTex for minimal latex installation
			tinyTexURL := "https://yihui.org/tinytex/install-bin-unix.sh"
			if !image.IsImageFileExist("install-bin-unix.sh") {
				fmt.Println("Downloading TinyTex installer...")
				if err := downloadFile(tinyTexURL, "install-bin-unix.sh"); err != nil {
					fmt.Println("Failed to download install-bin-unix.sh:", err)
					return err
				}
			}
			fmt.Println("install-bin-unix.sh downloaded successfully.")

			// Execute the TinyTex installation script
			if err := utils.RunCommand("sh", nil, nil, "install-bin-unix.sh"); err != nil {
				fmt.Println("Failed to install TinyTex:", err)
				return err
			}
			fmt.Println("TinyTex installed successfully, cleaning up files.")
			err := os.Remove("install-bin-unix.sh")
			if err != nil {
				fmt.Println("Error removing the file:", err)
				return err
			}

			// Set PATH variable
			os.Setenv("PATH", os.Getenv("PATH")+":"+os.Getenv("HOME")+"/bin")

			// Install the necessary packages
			fmt.Println("Installing necessary Latex packages....")
			packages := GetPackageList()
			if err := utils.RunCommand("tlmgr", nil, nil, append([]string{"install"}, packages...)...); err != nil {
				fmt.Println("Failed to install packages:", err)
				return err
			}
			fmt.Println("Packages installed successfully.")

			// Suggest the user to add the export statement to their .bashrc
			fmt.Println("######################## IMPORTANT NOTICE ########################")
			fmt.Println("The PATH has been updated for the current session. If you want to persist the changes, add the following line to your ~/.bashrc file:")
			fmt.Printf("export PATH=%s\n\n", os.Getenv("PATH")+":"+os.Getenv("HOME")+"/bin")
			fmt.Println("Remember to restart your terminal or run 'source ~/.bashrc' to apply the changes.")
			return nil
		}
		fmt.Println("The installer is not configured for this operating system, please refer to https://yihui.org/tinytex/ for more information")
		return nil
	}
	fmt.Println("All of the dependencies already installed, you are good to go!")
	return nil
}
