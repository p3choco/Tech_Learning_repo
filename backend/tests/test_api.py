from fastapi.testclient import TestClient
from main import app

client = TestClient(app)


def test_read_root():
    r = client.get("/")
    assert r.status_code == 200
    assert r.json() == {"message": "Welcome to the API"}


def test_get_all_items():
    r = client.get("/items")
    assert r.status_code == 200
    data = r.json()
    assert isinstance(data, list) and len(data) >= 2
    assert data[0]["id"] == 1


def test_get_single_item():
    r = client.get("/items/1")
    assert r.status_code == 200
    assert r.json()["name"] == "Item 1"


def test_item_not_found():
    r = client.get("/items/999")
    assert r.status_code == 200
    assert r.json() == {"error": "Item not found"}
