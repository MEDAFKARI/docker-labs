from fastapi import FastAPI, HTTPException
from fastapi.responses import RedirectResponse
from app.models import UrlRequest
from app.database import url_store
import uuid

app = FastAPI()


@app.get("/")
def home():
    return {"message": "URL Shortener API"}


@app.post("/shorten")
def shorten_url(request: UrlRequest):

    short_id = str(uuid.uuid4())[:8]

    url_store[short_id] = request.url

    return {
        "short_id": short_id,
        "original_url": request.url
    }


@app.get("/{short_id}")
def redirect_url(short_id: str):

    if short_id not in url_store:
        raise HTTPException(
            status_code=404,
            detail="URL not found"
        )

    return RedirectResponse(
        url=url_store[short_id]
    )


@app.get("/urls/all")
def get_urls():
    return url_store
