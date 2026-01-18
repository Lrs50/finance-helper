import time
from pprint import pprint
from typing import List

import requests

API_URL = "http://api:8080"


def create_user(
    username: str, name: str, email: str, password: str, phone_number: str
) -> dict | None:
    payload = {
        "username": username,
        "name": name,
        "password": password,
        "email": email,
        "phone_number": phone_number,
    }

    try:
        response = requests.post(
            f"{API_URL}/users",
            json=payload,
            headers={"Content-Type": "application/json"},
        )

        if response.status_code == 201:
            user_data = response.json()
            return user_data
        else:
            print(f"Failed to create user: {response.text}")
            return None

    except Exception as e:
        print(f"Error: {e}")
        return None


def get_all_users() -> List | None:
    try:
        response = requests.get(f"{API_URL}/users")

        if response.status_code == 200:
            users = response.json()
            return users
        else:
            print(f"Failed to get users: {response.text}")
            return None

    except Exception as e:
        print(f"Error: {e}")
        return None


def get_user_by_id(user_id: str) -> dict | None:
    try:
        response = requests.get(f"{API_URL}/users/by-id", params={"id": user_id})

        print(f"Status Code: {response.status_code}")

        if response.status_code == 200:
            user_data = response.json()
            return user_data
        else:
            print(f"Failed to get user: {response.text}")
            return None

    except Exception as e:
        print(f"Error: {e}")
        return None


def main():

    print("Testing API")
    while True:
        try:
            response = requests.get(API_URL)
            pprint(response.text)
        except Exception as e:
            print(f"Error connecting to API: {e}")

        pprint(get_all_users())
        time.sleep(10)


if __name__ == "__main__":
    main()
