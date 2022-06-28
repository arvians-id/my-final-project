import { ENDPOINT_API_GET_SUBMISSIONS_BY_USER, ENDPOINT_API_USER_SUBMIT_SUBMISSIONS } from "../constant/api";
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

export const API_USER_SUBMIT_SUBMISSION = async (courseCode, submissionId, file) => {
  let formData = new FormData();
  const config = {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  }
  const url = `${ENDPOINT_API_USER_SUBMIT_SUBMISSIONS(courseCode, submissionId)}`
  formData.append('file', file);
  const response = await axiosWithToken().post(`${ENDPOINT_API_USER_SUBMIT_SUBMISSIONS(courseCode, submissionId)}`, formData, config)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};