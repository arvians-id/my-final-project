import { ENDPOINT_API_CREATE_MODULE_ARTICLES, ENDPOINT_API_GET_LIST_MODULE_ARTICLES_BY_COURSE_CODE, ENDPOINT_API_DELETE_MODULE_ARTICLES, ENDPOINT_API_UPDATE_MODULE_ARTICLES, ENDPOINT_API_GET_DETAIL_ARTICLE_COURSE_BY_ARTICLE_ID_AND_COURSE_CODE } from "../constant/api";
import { axiosWithToken } from "./axiosWithToken";

export const API_CREATE_MODULE_ARTICLE = async (courseCode, data) => {
  const response = await axiosWithToken().post(`${ENDPOINT_API_CREATE_MODULE_ARTICLES(courseCode)}`, data)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_GET_ALL_ARTICLE_BY_COURSE_CODE = async (courseCode) => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_LIST_MODULE_ARTICLES_BY_COURSE_CODE(courseCode)}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_UPDATE_MODULE_ARTICLES = async (courseCode, articleId, data) => {
  const response = await axiosWithToken().patch(`${ENDPOINT_API_UPDATE_MODULE_ARTICLES(courseCode, articleId)}`, data)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_DELETE_MODULE_ARTICLES = async (courseCode, articleId) => {
  const response = await axiosWithToken().delete(`${ENDPOINT_API_DELETE_MODULE_ARTICLES(courseCode, articleId)}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};

export const API_GET_ARTICLE_DETAIL = async (courseCode, articleId) => {
  const response = await axiosWithToken().get(`${ENDPOINT_API_GET_DETAIL_ARTICLE_COURSE_BY_ARTICLE_ID_AND_COURSE_CODE(courseCode, articleId)}`)
    .then((response) => {
      return response;
    })
    .catch((error) => {
      return error;
    });

  return response;
};
