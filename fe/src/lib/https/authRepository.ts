import Repository from "@/lib/https/repository";

const urlAuth = "/admin";
const urlLoginByEmail = `/login-by-email-password`;
const urlSignUpMember = "/signup-by-email-password";
const urlSignUp = "/signup-information";

const urlVerifyEmail = "/auth/verify-email";
const urlLoginGuest = "/auth/login-guest";

export default class AuthRepository extends Repository {
  signUpMember = async (formData: FormData) => {
    const data = await this.axios.post(urlAuth.concat(urlSignUpMember), formData)
    if (data) {
      console.log(data)
      return true;
    }
    return false;
  };
  signUp = async (formData: FormData) => {
    const data = await this.axios.post(urlAuth.concat(urlSignUp), formData)
    if (data) {
      return true;
    }
    return false;
  };

  login = async (formData: FormData) => {
    await this.axios.post(urlAuth.concat(urlLoginByEmail), formData)
  }

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
