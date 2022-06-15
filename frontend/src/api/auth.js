import Axios from 'axios';
import { ENDPOINT_API_GET_CHECK_LOGIN_USER, ENDPOINT_API_GET_REGISTER_USER, ENDPOINT_API_POST_LOGIN_USER } from '../constant/api';

export const API_LOGIN = async ({ email, password }) => {
  const response = await Axios.post(ENDPOINT_API_POST_LOGIN_USER, {
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

export const API_REGISTER = async ({ name, username, email, password, role, gender, type_of_disability, birthdate }) => {
  const response = await Axios.post(ENDPOINT_API_GET_REGISTER_USER, {
    name,
    username,
    email,
    password,
    role,
    gender,
    type_of_disability,
    birthdate,
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
  const response = await Axios.get(`${ENDPOINT_API_GET_CHECK_LOGIN_USER}/${token}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};


