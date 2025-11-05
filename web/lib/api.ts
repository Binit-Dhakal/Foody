import apiClient from "./api-client";
import { SignInData, SignUpUserData } from "./validators";

export async function fetchSession() {
  const res = await apiClient.get("/accounts/session")

  if (!res.status) throw new Error("Failed to fetch session");
  return res;
}

export async function signIn(loginData: SignInData) {
  try {
    await apiClient.post("/accounts/login", {
      email: loginData.email,
      password: loginData.password
    })
  } catch (err: any) {
    throw new Error(err?.response?.data?.message || "Failed to sign in")
  }
}


export async function signUpUser(registerData: SignUpUserData) {
  try {
    await apiClient.post("/accounts/registerUser", {
      name: registerData.fullName,
      email: registerData.email,
      password: registerData.password,
      confirmPassword: registerData.confirmPassword,
    })
  } catch (err: any) {
    throw new Error(err?.response?.data?.message || "Failed to register")
  }
}


export async function logOut() {
  try {
    await apiClient.get("/accounts/logout")
  } catch (err: any) {
    throw new Error(err?.response?.data?.message || "Failed to log out")
  }
}
