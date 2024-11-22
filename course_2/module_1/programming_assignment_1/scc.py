import os
import urllib.request
import shutil

def main():
    url = "https://d3c33hcgiwev3.cloudfront.net/_410e934e6553ac56409b2cb7096a44aa_SCC.txt?Expires=1732320000&Signature=KJWsQopdph4scWJbt6-3ShgYA9sYtMFT2JYiGfjILBOVBAonDwSymetW8Oghe1swytfWEfKARNKQzXM8jRBmozVKyl~7ASj~7sURYXLPj9AKsktY-s169h-xtjoKYfpzRFb9ewcY~h98IMH8z44savQevRxj35lEslWsbvPSoNE_&Key-Pair-Id=APKAJLTNE6QMUY6HBC5A"
    output_file = "SCC.txt"

    # Get the current working directory to ensure the file is created in the same location as the program.
    current_dir = os.getcwd()
    file_path = os.path.join(current_dir, output_file)

    try:
        # Send an HTTP GET request to download the file.
        print(f"Downloading file from '{url}'...")
        with urllib.request.urlopen(url) as response:
            # Check if the HTTP response status code indicates success.
            if response.status < 200 or response.status >= 300:
                print(f"Error: HTTP status {response.status} received.")
                return

            # Save the content of the response to the output file.
            with open(file_path, "wb") as file:
                shutil.copyfileobj(response, file)

        print(f"File downloaded and saved successfully at: {file_path}")
    except urllib.error.URLError as e:
        print(f"Error downloading file: {e}")
    except OSError as e:
        print(f"Error saving file: {e}")

if __name__ == "__main__":
    main()
