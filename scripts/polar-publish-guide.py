"""Creates the 'Signal Loss Endings Guide' product on Polar (pay what you want, min 3 EUR) with both PDFs as downloads.
Usage: python3 scripts/polar-publish-guide.py releases/signal-loss-endings-guide-en.pdf releases/signal-loss-endings-guide-de.pdf
Token: ~/.config/polar/token (never printed)."""
import base64, hashlib, json, os, sys, urllib.request

TOKEN = os.environ.get("POLAR_TOKEN") or open(os.path.expanduser("~/.config/polar/token")).read().strip()

def call(method, path, body=None):
    req = urllib.request.Request("https://api.polar.sh/v1" + path, method=method,
        data=json.dumps(body).encode() if body is not None else None,
        headers={"Authorization": f"Bearer {TOKEN}", "Content-Type": "application/json", "User-Agent": "curl/8.7.1", "Accept": "application/json"})
    try:
        with urllib.request.urlopen(req) as r:
            return json.loads(r.read() or b"{}")
    except urllib.error.HTTPError as e:
        sys.exit(f"HTTP {e.code} {e.read().decode()[:800]}")

def upload(path):
    data = open(path, "rb").read()
    sha = base64.b64encode(hashlib.sha256(data).digest()).decode()
    f = call("POST", "/files/", {"name": os.path.basename(path), "mime_type": "application/pdf", "size": len(data),
        "checksum_sha256_base64": sha, "service": "downloadable",
        "upload": {"parts": [{"number": 1, "chunk_start": 0, "chunk_end": len(data), "checksum_sha256_base64": sha}]}})
    part = f["upload"]["parts"][0]
    req = urllib.request.Request(part["url"], method="PUT", data=data, headers={**(part.get("headers") or {}), "User-Agent": "curl/8.7.1"})
    for attempt in range(4):  # the storage endpoint sometimes resets the connection
        try:
            with urllib.request.urlopen(req) as r:
                etag = r.headers["ETag"]
            break
        except (urllib.error.URLError, ConnectionError) as e:
            print("retry", attempt + 1, e); __import__("time").sleep(2 * (attempt + 1))
    else:
        sys.exit("upload failed")
    done = call("POST", f"/files/{f['id']}/uploaded", {"id": f["upload"]["id"], "path": f["upload"]["path"],
        "parts": [{"number": 1, "checksum_etag": etag, "checksum_sha256_base64": sha}]})
    print("uploaded", os.path.basename(path), done.get("is_uploaded"))
    return done["id"]

DESC = """**Stuck on Gryps-4? Curious about every ending?** The Signal Loss Endings Guide is a short PDF with the winning line for all eight endings (one of them hidden), the shared road through chapters 1-7, every item and where to find it, and the rules behind the satellite route, luck and New Game+.

English and German edition included. **Contains spoilers.** The game itself stays free: https://aeon022.github.io/signal_loss/"""

files = [upload(p) for p in sys.argv[1:]]
benefit = call("POST", "/benefits/", {"type": "downloadables", "description": "Signal Loss Endings Guide (EN + DE)", "properties": {"files": files}})
product = call("POST", "/products/", {"name": "Signal Loss: Endings Guide (PDF)", "description": DESC, "recurring_interval": None,
    "prices": [{"amount_type": "custom", "price_currency": "eur", "minimum_amount": 300, "preset_amount": 500}]})
call("POST", f"/products/{product['id']}/benefits", {"benefits": [benefit["id"]]})
link = call("POST", "/checkout-links/", {"products": [product["id"]], "payment_processor": "stripe"})
print("product", product["id"]); print("checkout", link["url"])
