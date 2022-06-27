import { ENDPOINT_API_GET_COURSE_BY_CODE, ENDPOINT_API_GET_COURSE_BY_USER_LOGIN } from "../constant/api";
import { axiosWithToken } from "./axiosWithToken";

export const API_GET_COURSE_BY_USER_LOGIN = async () => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_COURSE_BY_USER_LOGIN}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};


export const API_GET_COURSE_BY_CODE = async (code) => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_COURSE_BY_CODE}/${code}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};



