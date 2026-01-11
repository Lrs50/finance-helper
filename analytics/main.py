import time
from pprint import pprint

import requests


def main():
    while True:
        try:
            response = requests.get("http://api:8080")
            pprint(response.text)
        except Exception as e:
            print(f"Error connecting to API: {e}")

        time.sleep(10)


if __name__ == "__main__":
    main()
