"use server";

import { ApiConstants } from "@/constants/api";
import { publicAxios } from "@/lib/api";

export async function register(payload: ISignupPayload) {
  console.log({ baseURL: publicAxios.defaults.baseURL, loc: "regiseter" });
  const response = await publicAxios.post(ApiConstants.REGISTER, payload);
  console.log(response);
  return response;
}
export async function login(payload: ISigninPayload) {
  console.log({ baseURL: publicAxios.defaults.baseURL, loc: "login" });
  try {
    const response = await publicAxios.post(
      "http://localhost:8000/api/v1/authentication/login",
      payload
    );
    console.log(response);
    return response;
  } catch (error) {
    console.log(error);
    throw error;
  }
}
export async function forgotPassword() {}
export async function resetPassword() {}
