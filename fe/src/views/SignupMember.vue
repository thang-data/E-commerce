<template>
  <div class="form-container sign-up-container">
    <form>
      <h1 class="mb-30">Create Account</h1>
      <span>or use your email for registration</span>
      <input v-model="email" type="email" placeholder="Email" />
      <input v-model="lastName" type="text" placeholder="Last Name" />
      <input v-model="firstName" type="text" placeholder="First Name" />
      <button @click="submit">Sign Up</button>
     <a class="signin" href="/signup">Have you already registered as a member?</a>
    </form>
  </div>
  <ContainerView :signUp="true" />
</template>
<script lang="ts" setup>
import AuthRepository from "@/lib/https/authRepository";
import RepositoryFactory from "@/lib/https/repositoryFactory";
import { ref } from "vue";
import { useRouter } from "vue-router";
import ContainerView from "../components/containerView.vue";

const email = ref("")
const lastName = ref("")
const firstName = ref("")
const {  signUpMember } = RepositoryFactory.getRepository<AuthRepository>(AuthRepository).withRouter(useRouter())
const router = useRouter()
async function submit() {

  const formData = new FormData()
  formData.append("email", email.value)
  formData.append("lastName", lastName.value)
  formData.append("firstName", firstName.value)
  try {
    await signUpMember(formData)
	  router.push('/signup')
  } catch (errors: any) {
	return
  }
}

</script>
<style lang="scss" scoped>
@import url("https://fonts.googleapis.com/css?family=Montserrat:400,800");

.form-container {
  position: absolute;
  top: 0;
  height: 100%;
  transition: all 0.6s ease-in-out;
}
.sign-up-container {
  left: 0;
  width: 50%;
  //opacity: 0;
  z-index: 1;
}

.container.right-panel-active .sign-up-container {
  transform: translateX(100%);
  opacity: 1;
  z-index: 5;
  animation: show 0.6s;
}
form {
  background-color: #ffffff;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-direction: column;
  padding: 0 50px;
  height: 100%;
  text-align: center;
}

h1 {
  font-weight: bold;
  margin: 0 0 30px;
}
.social-container {
  margin: 20px 0;
}

.social-container a {
  border: 1px solid #dddddd;
  border-radius: 50%;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  margin: 0 5px;
  height: 40px;
  width: 40px;
}

span {
  font-size: 12px;
}

input {
  background-color: #eee;
  border: none;
  padding: 12px 15px;
  margin: 8px 0;
  width: 100%;
}

button {
  border-radius: 20px;
  border: 1px solid #ff4b2b;
  background-color: #ff4b2b;
  color: #ffffff;
  font-size: 12px;
  font-weight: bold;
  padding: 12px 45px;
  letter-spacing: 1px;
  text-transform: uppercase;
  transition: transform 80ms ease-in;
}

button:active {
  transform: scale(0.95);
}

button:focus {
  outline: none;
}

button.ghost {
  background-color: transparent;
  border-color: #ffffff;
}
.signin{
    margin-top: 10px;
    font-size: 12px;
}
</style>
