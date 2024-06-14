"use server";
export type State = {
  status: "success" | "error";
  message: string;
  data?: any;
} | null;
import { ApiConstants } from "@/constants/api";
import { publicAxios } from "@/lib/api";

export async function register(payload: ISignupPayload) {
  const response = await publicAxios.post(ApiConstants.REGISTER, payload);
  console.log(response);
  return response;
}
export async function login(
  prevState: State | null,
  formData: FormData
): Promise<State> {
  const payload = {
    email: formData.get("email"),
    password: formData.get("password"),
  };
  try {
    const response = await publicAxios.post(
      "http://localhost:8000/api/v1/authentication/login",
      payload
    );
    return {
      status: "success",
      message: "Login successful.",
      data: response.data.data,
    };
  } catch (error: any) {
    return {
      status: "error",
      message: error?.response?.data?.data?.error || "Something went wrong.",
    };
  }
}
export async function forgotPassword() {}
export async function resetPassword() {}
