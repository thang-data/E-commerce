import Repository from "@/lib/https/repository";
import LoginData from "@/models/LoginData";

const urlAuth = "/admin";
const urlLoginByEmail = `${urlAuth}/login-by-email-password`;
const urlSignUp = "/auth/signup-by-email-password";
const urlVerifyEmail = "/auth/verify-email";
const urlLoginGuest = "/auth/login-guest";

export default class AuthRepository extends Repository {
  signUp = async (formValues: FormData) => {
    const data = await this.axios.post(urlSignUp, formValues);
    if (data) {
      return true;
    }
    return false;
  };

  loginByEmail = async (formValues: FormData): Promise<false | LoginData> => {
    const data = await this.axios.post<LoginData>(urlLoginByEmail, formValues);
    if (data) {
      return data.data;
    }
    return false;
  };

  verifyEmail = async (code: string) => {
    const data = await this.axios.get(`${urlVerifyEmail}?code=${code}`);
    if (data) {
      return data;
    }
    return false;
  };

  loginGuest = async (formValues: FormData) => {
    const data = await this.axios.post(urlLoginGuest, formValues);
    if (data) {
      return data.data;
    }
    return false;
  };
}
