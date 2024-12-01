import os
import urllib.request
import shutil

def main():
    url = "https://d3c33hcgiwev3.cloudfront.net/_fe8d0202cd20a808db6a4d5d06be62f4_clustering_big.txt?Expires=1733184000&Signature=bF9fl-vxZuVopJdZfYVaKXV84nLQj4usxkcI67aIdfQRRl5L0j4JCqNW39~R0qTiInE4BI75mNIJ2VZedbsqJc3vCgjsaf5uxTKLixzjrYWWSWrzTJ~GOV3x0roVplwxbT0TybFcGXEP0z9k2bmNA9Sn0hPuvtHkHGyXFTC33Oo_&Key-Pair-Id=APKAJLTNE6QMUY6HBC5A"
    output_file = "clustering_big.txt"

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
