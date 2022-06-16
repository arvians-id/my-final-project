import Axios from 'axios';
import { URL_API_GET_CHECK_LOGIN_USER, URL_API_POST_LOGIN_USER } from '../constant/api';

export const API_LOGIN = async ({ email, password }) => {
  const response = await Axios.post(URL_API_POST_LOGIN_USER, {
    email,
    password
  })
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_CHECK_LOGIN = async (token) => {
  const response = await Axios.get(`${URL_API_GET_CHECK_LOGIN_USER}/${token}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};


