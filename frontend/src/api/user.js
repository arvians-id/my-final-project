import { ENDPOINT_API_GET_LIST_USER, ENDPOINT_API_GET_USER_BY_ID, ENDPOINT_API_UPDATE_USER_DETAIL_PROFIE } from "../constant/api";
import { axiosWithToken } from "./axiosWithToken";

export const API_GET_USER_DETAIL_BY_ID = async (userId) => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_USER_BY_ID}/${userId}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_UPDATE_USER_PROFILE_DETAIL = async (data, userId) => {
  const response = await axiosWithToken().put(`${ENDPOINT_API_UPDATE_USER_DETAIL_PROFIE}/${userId}`, data)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_GET_LIST_USER = async () => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_LIST_USER}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};
