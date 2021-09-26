export interface EmailForm {
  Email: string;
  Password: string;
}

export interface PhoneForm {
  Phone: string;
  Password: string;
}

export interface SMSForm {
  Phone: string;
  Code: string;
}

export interface OauthForm {
  QRCodeSrc: string;
}

export interface LoginResponse {
  serviceResponse: {
    authenticationSuccess?: {
      redirectService: string;
    };
    authenticationFailure?: {
      description: string;
    };
  };
}
