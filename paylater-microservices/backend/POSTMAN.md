# Postman RSA Login Setup

Frontend login works because it encrypts correctly. Postman failed because:
1. **`login_password` must be a collection variable** — not typed in the body as `"password"`.
2. **Async pre-request scripts** can send the request before encryption finishes.

This collection uses **synchronous** encryption with a cached public key.

## Setup (every session)

1. **Delete** any old **PayLater API** collection in Postman.

2. **Import** `backend/PayLater API.postman_collection.json`  
   → Collection name: **PayLater API (RSA Login)**

3. Open collection **Variables** (Current value column):

   | Variable | Set to |
   |----------|--------|
   | `baseUrl` | `http://localhost:8080` |
   | `login_email` | `admin@gmail.com` |
   | `login_password` | your real admin password (same as frontend login) |

4. Run **Get Login Public Key** → must return `200`  
   (caches `login_public_key` and sets `encrypted_password` if `login_password` is set)

5. Run **Admin login** → expect `200` + JWT

## Request body (must look like this)

```json
{
  "email": "{{login_email}}",
  "encrypted_password": "{{encrypted_password}}"
}
```

**Never** send `"password"` in the body.

## Response codes

| Code | Meaning |
|------|---------|
| **400** `invalid login credentials` | Wrong body format, empty `encrypted_password`, or decrypt failed |
| **401** `invalid password` / `invalid email` | Encryption OK, wrong email/password in DB |
| **200** | Success |

## Postman Console errors

| Error | Fix |
|-------|-----|
| `login_password is empty` | Set `login_password` in collection Variables |
| `login_public_key is missing` | Run **Get Login Public Key** first |
