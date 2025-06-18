from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from typing import List

app = FastAPI()

# Enable CORS
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  #TODO replace with frontend url
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

class Item(BaseModel):
    id: int
    name: str
    description: str

items = [
    Item(id=1, name="Item 1", description="Description 1"),
    Item(id=2, name="Item 2", description="Description 2"),
]

@app.get("/")
async def read_root():
    return {"message": "Welcome to the API"}

@app.get("/items", response_model=List[Item])
async def read_items():
    return items

@app.get("/items/{item_id}", response_model=Item)
async def read_item(item_id: int):
    for item in items:
        if item.id == item_id:
            return item
    return {"error": "Item not found"} 