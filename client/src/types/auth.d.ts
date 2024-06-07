interface ISigninPayload {
  email: string;
  password: string;
}

interface ISignupPayload extends ISigninPayload {
  fullName: string;
}
