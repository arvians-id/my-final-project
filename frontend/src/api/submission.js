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
  console.log(courseCode)
  console.log(submissionId)
  console.log(file)
  const url = `${ENDPOINT_API_USER_SUBMIT_SUBMISSIONS(courseCode, submissionId)}`
  console.log(url)
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


// const URL = `${baseAPIURL}file/upload`;
// const JSONAssignedPersons = JSON.stringify(selectedMembers);
// const JSONPrivilegeUser = JSON.stringify(privilegeUser);
// const JSONParentFile =
//   data?.parent === undefined ? '' : JSON.stringify(data?.parent.id);
// let formData = new FormData();
// const config = {
//   headers: {
//     'Content-Type': 'multipart/form-data',
//   },
//   onUploadProgress: function (progressEvent: any) {
//     handleOnUploadProgress(progressEvent);
//   },
// };

// formData.append('file', data?.file);
// formData.append('card_id', String(cardId));
// formData.append('privilege_user', JSONPrivilegeUser);
// formData.append('parent_id', JSONParentFile);
// formData.append('assigned_persons', JSONAssignedPersons);

// const response = await axiosWithToken()
//   .post(URL, formData, config)
//   .then((response) => {
//     return response;
//   })
//   .catch((error) => {
//     return error.response;
//   });

