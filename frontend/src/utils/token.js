import { USER_TOKEN } from '../constant/localStorage';
import { getLocal, setLocal } from './localStorage';

export const localSaveToken = (token) => {
  setLocal(USER_TOKEN, token);
};

export const localLoadToken = () => {
  return getLocal(USER_TOKEN);
};

export const localClearToken = () => {
  setLocal(USER_TOKEN, '');
};
