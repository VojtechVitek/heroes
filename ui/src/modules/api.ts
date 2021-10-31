import { fetch } from 'whatwg-fetch';
import { API } from './rpc.gen';

const apiUrl = 'http://localhost:7777';

export const api: API = new API(apiUrl, fetch);
