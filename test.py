import os
import json
import base64
import win32crypt
from Crypto.Cipher import AES

path = r'%LocalAppData%\Microsoft\Edge\User Data\Local State'
path = os.path.expandvars(path)
with open(path, 'r') as file:
    encrypted_key = json.loads(file.read())['os_crypt']['encrypted_key']
    # Base64 decoding
    encrypted_key = base64.b64decode(encrypted_key)
    # Remove DPAPI
    encrypted_key = encrypted_key[5:]
    decrypted_key = win32crypt.CryptUnprotectData(
        encrypted_key, None, None, None, 0)[1]  # Decrypt key

print([b for b in decrypted_key])

encrypted_value = bytes(
    [118, 49, 48, 34, 21, 83, 112, 110, 89, 134, 192, 229, 153, 132, 94, 224, 2, 249, 6, 6, 158, 96, 196, 163, 212, 180, 225, 157, 61, 122, 181, 232, 143, 82, 58, 167, 145])

cipher = AES.new(decrypted_key, AES.MODE_GCM, nonce=encrypted_value[3:3+12])
decrypted_value = cipher.decrypt_and_verify(
    encrypted_value[3+12:-16], encrypted_value[-16:])

print(decrypted_value)
