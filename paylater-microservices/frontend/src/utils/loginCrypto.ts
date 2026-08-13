import { api } from '../services/api'

interface LoginPublicKeyResponse {
  public_key: string
}

let cachedPublicKey: string | null = null

function pemToArrayBuffer(pem: string): ArrayBuffer {
  const base64 = pem
    .replace(/-----BEGIN PUBLIC KEY-----/, '')
    .replace(/-----END PUBLIC KEY-----/, '')
    .replace(/\s/g, '')
  const binary = atob(base64)
  const bytes = new Uint8Array(binary.length)

  for (let index = 0; index < binary.length; index += 1) {
    bytes[index] = binary.charCodeAt(index)
  }

  return bytes.buffer
}

function uint8ArrayToBase64(bytes: Uint8Array): string {
  let binary = ''

  for (let index = 0; index < bytes.length; index += 1) {
    binary += String.fromCharCode(bytes[index]!)
  }

  return btoa(binary)
}

async function importPublicKey(pem: string): Promise<CryptoKey> {
  return crypto.subtle.importKey(
    'spki',
    pemToArrayBuffer(pem),
    {
      name: 'RSA-OAEP',
      hash: 'SHA-256',
    },
    false,
    ['encrypt'],
  )
}

async function fetchLoginPublicKey(): Promise<string> {
  if (cachedPublicKey) {
    return cachedPublicKey
  }

  const response = await api.get<LoginPublicKeyResponse>('/login/public-key')
  const publicKey = response.data.public_key
  cachedPublicKey = publicKey
  return publicKey
}

export async function encryptLoginPassword(password: string): Promise<string> {
  const publicKeyPem = await fetchLoginPublicKey()
  const key = await importPublicKey(publicKeyPem)
  const payload = JSON.stringify({
    password,
    ts: Math.floor(Date.now() / 1000),
  })
  const encrypted = await crypto.subtle.encrypt(
    { name: 'RSA-OAEP' },
    key,
    new TextEncoder().encode(payload),
  )

  return uint8ArrayToBase64(new Uint8Array(encrypted))
}

export async function buildEncryptedLoginBody(
  email: string,
  password: string,
): Promise<{ email: string; encrypted_password: string }> {
  return {
    email,
    encrypted_password: await encryptLoginPassword(password),
  }
}

/** Clears cached public key (useful for tests or key rotation). */
export function clearLoginPublicKeyCache(): void {
  cachedPublicKey = null
}
