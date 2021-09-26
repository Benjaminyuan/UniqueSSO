import { LoginType } from "@/constant/loginType";
import {
  EmailForm,
  LoginResponse,
  OauthForm,
  PhoneForm,
  SMSForm,
} from "@/model/form";
import axios, { AxiosInstance } from "axios";

class Endpoint {
  static emailLogin = (service: string) =>
    `/cas/login?type=${LoginType.Email}&service=${service}`;

  static phoneLogin = (service: string) =>
    `/cas/login?type=${LoginType.Phone}&service=${service}`;

  static smsLogin = (service: string) =>
    `/cas/login?type=${LoginType.SMS}&service=${service}`;

  static larkLogin = (service: string) =>
    `/cas/login?type=${LoginType.LarkOauth}&service=${service}`;
}

export class RestClient {
  instance: AxiosInstance;

  constructor(public baseURL: string) {
    this.instance = axios.create({
      baseURL: baseURL,
      headers: {
        "Content-Type": "application/json",
      },
    });
  }

  private async wrapLoginRequest(
    url: string,
    jsonData: Record<string, unknown>
  ): Promise<string> {
    // const resp = await this.instance
    //   .post(url, { json: data })
    try {
      const { data } = await this.instance.post<LoginResponse>(url, jsonData);
      if (data.serviceResponse.authenticationSuccess !== undefined) {
        return data.serviceResponse.authenticationSuccess.redirectService;
      } else {
        throw Error("response");
      }
    } catch ({ response }) {
      //FIXME
      console.log(response);
      // response.code;
      throw Error("axios error");
    }
    // if (resp.serviceResponse.authenticationSuccess) {
    //   return resp.serviceResponse.authenticationSuccess.redirectService;
    // }
    // if (resp.serviceResponse.authenticationFailure) {
    //   throw new Error(resp.serviceResponse.authenticationFailure.description);
    // }
    // throw new Error("send http request failed");
  }

  async loginByEmail(form: EmailForm, service: string): Promise<string> {
    return this.wrapLoginRequest(Endpoint.emailLogin(service), {
      email: form.Email,
      password: form.Password,
    });
  }

  async loginByPhone(form: PhoneForm, service: string): Promise<string> {
    return this.wrapLoginRequest(Endpoint.phoneLogin(service), {
      phone: form.Phone,
      password: form.Password,
    });
  }

  async loginBySMS(form: SMSForm, service: string): Promise<string> {
    return this.wrapLoginRequest(Endpoint.smsLogin(service), {
      phone: form.Phone,
      code: form.Code,
    });
  }

  async LoginByOauth(form: OauthForm, service: string): Promise<string> {
    return this.wrapLoginRequest(Endpoint.larkLogin(service), {
      qrcode_src: form.QRCodeSrc,
    });
  }
}
