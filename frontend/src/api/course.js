import {
  ENDPOINT_API_CREATE_COURSE, ENDPOINT_API_GET_COURSE_BY_CODE, ENDPOINT_API_GET_COURSE_BY_USER_LOGIN, ENDPOINT_API_GET_ALL_COURSE, ENDPOINT_API_UPDATE_COURSE, ENDPOINT_API_DELETE_COURSE, ENDPOINT_API_LIST_USER_IN_COURSE,
  ENDPOINT_API_REMOVE_USER_IN_COURSE,
  ENDPOINT_API_ADD_USER_IN_COURSE,
} from "../constant/api";
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

export const API_CREATE_COURSE = async (data) => {
  const response = await axiosWithToken().post(`${ENDPOINT_API_CREATE_COURSE}`, data)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_GET_ALL_COURSE = async () => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_ALL_COURSE}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_UPDATE_COURSE = async (courseCode, data) => {
  const response = await axiosWithToken().patch(`${ENDPOINT_API_UPDATE_COURSE(courseCode)}`, data)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_DELETE_COURSE = async (courseCode) => {
  const response = await axiosWithToken().delete(`${ENDPOINT_API_DELETE_COURSE(courseCode)}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_GET_LIST_USER_IN_COURSE = async (courseCode) => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_LIST_USER_IN_COURSE(courseCode)}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_REMOVE_USER_IN_COURSE = async (userId, courseCode) => {
  const response = await axiosWithToken().delete(`${ENDPOINT_API_REMOVE_USER_IN_COURSE(userId, courseCode)}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

// {
//   "user_id" : "integer", // foreign key1
//   "course_id" : "integer" // foreign key2
// }
export const API_ADD_USER_IN_COURSE = async (data) => {
  const response = await axiosWithToken().post(`${ENDPOINT_API_ADD_USER_IN_COURSE}`, data)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

