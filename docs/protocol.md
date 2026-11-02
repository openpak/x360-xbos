# Protocol: x360-xbos

Live service 21 (`xbos`), XAM 17559.

After the user logon the console asks for the certificate revocation list with 16 bytes of its key 0x42. Reply:
256-byte signature, 16 bytes, `u32` count of 20-byte revoked hashes. OpenPak sends 276 zero bytes. The other
service-21 requests (0x915, 0x916, 0x91A, 0x92A) are not served. Services 3, 4 and 9 also use `/xbos`; 9
(FindMediaInstanceUrls) is x360-store.

## Wharf

Behaviour is carried over unchanged from x360-xuacs v0.9.7, where the post-logon request is answered in Wharf.
