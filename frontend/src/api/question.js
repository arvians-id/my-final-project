import { ENDPOINT_API_GET_QUESTIONS_BY_USER, ENDPOINT_API_GET_ALL_QUESTIONS } from "../constant/api";
import { axiosWithToken } from "./axiosWithToken";

export const API_GET_QUESTION_BY_USER_ID = async (userId) => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_QUESTIONS_BY_USER}/${userId}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_GET_ALL_QUESTION = async () => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_ALL_QUESTIONS}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};
