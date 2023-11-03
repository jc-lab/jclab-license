import { Buffer } from 'buffer';
import { sha256 } from '@noble/hashes/sha256'
import { toUrlBase64 } from './base64';

function makeKid(publicKey: Buffer): string {
  const h = sha256.create().update(publicKey).digest();
  return toUrlBase64(Buffer.from(h.slice(0, 8)).toString('base64'));
}

export interface PubKey {
  key: Buffer;
  notAfter: number;
}

const keys: Record<string, PubKey> = {};

function addKey(input: PubKey) {
  const kid = makeKid(input.key);
  keys[kid] = input;
}

export function getKey(kid: string): PubKey | undefined {
  return keys[kid];
}

addKey({
  key: Buffer.from([0xda, 0xc6, 0x67, 0xc3, 0x9d, 0x88, 0x73, 0xd1, 0x87, 0x77, 0xab, 0xdc, 0x34, 0x5f, 0x3b, 0xbb, 0x9b, 0x7e, 0x44, 0x50, 0x73, 0x59, 0x22, 0xea, 0x2b, 0x1d, 0x2f, 0x28, 0x48, 0xb4, 0x3a, 0xc1]),
  notAfter: -1,
});

