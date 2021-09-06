import { fetch } from 'whatwg-fetch';
import { ArchiveAPI } from '../modules/Archive.ts';

export const api = new ArchiveAPI('http://localhost:2022', fetch);

// console.log(import.meta.env);
//export const api = new ArchiveAPI('http://192.168.41.51', fetch);
