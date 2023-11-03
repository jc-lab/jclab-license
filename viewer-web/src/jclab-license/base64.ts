import {Buffer} from "buffer";

export function toUrlBase64(general: string): string {
  return general
      .replace(/\+/g, '-')
      .replace(/\//g, '_')
      .replace(/=/g, '');
}

export function fromUrlBase64(urlSafe: string): string {
  let output = urlSafe
      .replace(/-/g, '+')
      .replace(/_/g, '/');
  const tmp = output.length % 4;
  if (tmp) {
    output += '='.repeat(4 - tmp);
  }
  return output;
}

export function urlBase64Decode(urlSafe: string): Buffer {
  return Buffer.from(fromUrlBase64(urlSafe), 'base64');
}