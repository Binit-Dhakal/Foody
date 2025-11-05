import apiClient from "./api-client";

import { SignUpResturantData } from "./validators";

export async function signUpResturant(registerData: SignUpResturantData, license: File) {
  const formData = new FormData();
  formData.append("fullName", registerData.fullName);
  formData.append("email", registerData.email);
  formData.append("password", registerData.password);
  formData.append("confirmPassword", registerData.confirmPassword);
  formData.append("resturantName", registerData.resturantName);
  formData.append("resturantLicense", license);

  try {
    await apiClient.post("/accounts/registerVendor", formData, {
      headers: { "Content-Type": "multipart/form-data" },
    });
  } catch (err: any) {
    throw new Error(err?.response?.data?.message || "Failed to register")
  }
}

export async function getVendorSummary() {
  try {
    const res = await apiClient.get("/vendor/profile/summary")
    return res
  } catch {
    throw new Error("Failed to fetch vendor info")
  }
}
