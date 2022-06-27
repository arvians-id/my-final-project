import { ENDPOINT_API_GET_SUBMISSIONS_BY_USER } from "../constant/api";
import { axiosWithToken } from "./axiosWithToken";

export const API_GET_SUBMISSION_BY_USER_LOGIN = async (query) => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_SUBMISSIONS_BY_USER}${query}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};
