import { AxiosInstance } from "axios";
import axiosInstance from "./setting";

export default class Repository {
  axios: AxiosInstance = axiosInstance;
}
