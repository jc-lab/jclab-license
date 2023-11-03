import * as ed25519 from '@noble/ed25519';
import * as pako from 'pako';
import { Buffer } from 'buffer';
import { Claims } from './model';
import { urlBase64Decode } from './base64';
import { getKey } from './keys';

export * from './model';

export class LicenseError extends Error {
  constructor(message: string) {
    super(message);
  }
}

interface Header {
  alg: string;
  kid: string;
}

export async function verify(token: string): Promise<Claims> {
  const tokens = token.split('.');
  const signedPart = Buffer.from(tokens[0] + '.' + tokens[1]);
  const header: Header = JSON.parse(urlBase64Decode(tokens[0]).toString());
  const compressedBody = urlBase64Decode(tokens[1]);
  const signature = urlBase64Decode(tokens[2]);

  const pubKey = getKey(header.kid);
  if (!pubKey) {
    return Promise.reject(new LicenseError('invalid issuer'));
  }

  if (!await ed25519.verifyAsync(signature.toString('hex'), signedPart.toString('hex'), pubKey.key.toString('hex'))) {
    return Promise.reject(new LicenseError('invalid issuer'));
  }

  const body = pako.inflate(compressedBody);
  return JSON.parse(Buffer.from(body).toString());
}