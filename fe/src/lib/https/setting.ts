import axios from "axios";
import storage from "@/lib/storage/storageHelper";

const baseDomain: string = process.env.VUE_APP_API_URL as string;

function authHeader(formData: any) {
  const localStorage = storage.getLocalStorage();
  const userToken = localStorage.get("userToken");
  const userAuthorize = localStorage.get("userTokenCSRF");
  const sessionId = localStorage.get("sessionId");
  if (userToken) {
    if (formData) {
      return {
        "Content-Type": "multipart/form-data",
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Headers": `${process.env.VUE_APP_URL as string}`,
        Authorization: `Bearer ${userToken}`,
        Session: sessionId,
      };
    }
    return {
      "Content-Type": "application/json",
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Headers": `${process.env.VUE_APP_URL as string}`,
      Authorization: `Bearer ${userToken}`,
      Session: sessionId,
    };
  }
  if (userAuthorize) {
    localStorage.remove("userTokenCSRF");
    return {
      "Content-Type": "application/json",
      "X-CSRF-Token": userAuthorize,
    };
  }
  return {
    "Content-Type": "application/json",
    Session: sessionId,
  };
}

const axiosInstance = axios.create({
  baseURL: baseDomain,
  responseType: "json",
  withCredentials: true,
});

/* eslint-disable no-param-reassign */
// Add a request interceptor
axiosInstance.interceptors.request.use(
  (config: any) => {
    config.headers =
      config.data instanceof FormData ? authHeader(true) : authHeader(false);
    return config;
  },
  (error: any) => Promise.reject(error)
);
/* eslint-enable no-param-reassign */

axiosInstance.interceptors.response.use(
  (response: any) => response,
  (err: any) => {
    // Any status codes that falls outside the range of 2xx cause this function to trigger
    // Do something with response error
    if (err.response && err.response.status === 401) {
      localStorage.clear();
      window.location.href = "/login";
    }

    return Promise.reject(err.response.data);
  }
);

export default axiosInstance;
