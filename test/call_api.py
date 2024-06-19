import requests

BASE_URL = "http://localhost:8080"

def test_handle_init_api():
    response = requests.get(f"{BASE_URL}/init")
    print(response.text)

def test_handle_switch_api(world_name):
    response = requests.get(f"{BASE_URL}/switch", params={"worldName": world_name})
    print(response.text)

def test_handle_get_api(cid):
    response = requests.get(f"{BASE_URL}/get", params={"cid": cid})
    print(response.text)

def test_handle_put_api(args):
    response = requests.get(f"{BASE_URL}/put", params={"args": args})
    print(response.text)

def test_handle_set_password_api(password):
    response = requests.get(f"{BASE_URL}/setPassword", params={"password": password})
    print(response.text)

def test_handle_cat_api(file_hash):
    response = requests.get(f"{BASE_URL}/cat", params={"fileHash": file_hash})
    print(response.text)

def test_handle_get_world_cid_api():
    response = requests.get(f"{BASE_URL}/getWorldCID")
    print(response.text)

def test_handle_download_world_api(cid):
    response = requests.get(f"{BASE_URL}/downloadWorld", params={"cid": cid})
    print(response.text)

def test_handle_get_world_data_api():
    response = requests.get(f"{BASE_URL}/getWorldData")
    print(response.text)

if __name__ == "__main__":
    # Testing each endpoint
    test_handle_init_api()
    test_handle_switch_api("example_world")
    test_handle_get_api("example_cid")
    test_handle_put_api("arg1 arg2 arg3")
    test_handle_set_password_api("example_password")
    test_handle_cat_api("example_file_hash")
    test_handle_get_world_cid_api()
    test_handle_download_world_api("example_cid")
    test_handle_get_world_data_api()
