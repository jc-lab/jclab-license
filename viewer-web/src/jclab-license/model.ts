export enum LicenseType {
  OpenSource = 'opensource',
  Commercial = 'commercial',
}

export interface Claims {
  iat: number;
  iss: string;
  product: string;
  lictype: string;
  lickeyh: string;
  licmaxv: number;
  lisname: string;
  lisemail: string;
}
